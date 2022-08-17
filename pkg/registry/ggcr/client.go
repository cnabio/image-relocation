/*
 * Copyright (c) 2019-Present Pivotal Software, Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ggcr

import (
	"fmt"
	"net/http"
	"os"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/layout"

	"github.com/cnabio/image-relocation/pkg/image"
	"github.com/cnabio/image-relocation/pkg/registry"
)

const outputDirPermissions = 0755

// RegistryClient provides methods for building abstract images.
// This interface is not intended for external consumption.
type RegistryClient interface {
	// ReadRemoteImage builds an abstract image from a repository.
	ReadRemoteImage(n image.Name) (registry.Image, error)

	// NewImageFromManifest builds an abstract image from an image manifest.
	NewImageFromManifest(img v1.Image) registry.Image

	// NewImageFromIndex builds an abstract image from an image index.
	NewImageFromIndex(img v1.ImageIndex) registry.Image
}

type manifestWriter func(v1.Image, image.Name) error
type indexWriter func(v1.ImageIndex, image.Name) error

type client struct {
	readRemoteImage  func(image.Name) (registry.Image, error)
	writeRemoteImage manifestWriter
	writeRemoteIndex indexWriter
}

var (
	// Ensure client conforms to the relevant interfaces.
	_ RegistryClient  = &client{}
	_ registry.Client = &client{}
)

// Option represents a functional option for NewRegistryClient.
type Option func(*client)

// WithTransport overrides the default transport used for remote operations, default is http.DefaultTransport.
func WithTransport(transport http.RoundTripper) Option {
	return func(c *client) {
		// There isn't a good way to have a separate function named WithInsecure, because transport isn't available
		// at that point, insecure and transport need to be set at the same time due to how client and its configuration
		// is structured. So I am inferring that plain http is allowed when the transport specified allows insecure TLS
		// connections.
		insecure := false
		if t, ok := transport.(*http.Transport); ok {
			if t.TLSClientConfig != nil {
				insecure = t.TLSClientConfig.InsecureSkipVerify
			}
		}

		writeRemoteImageFunc := writeRemoteImage(transport, insecure)
		writeRemoteIndexFunc := writeRemoteIndex(transport, insecure)

		c.readRemoteImage = readRemoteImage(writeRemoteImageFunc, writeRemoteIndexFunc, transport, insecure)
		c.writeRemoteImage = writeRemoteImageFunc
		c.writeRemoteIndex = writeRemoteIndexFunc
	}
}

// NewRegistryClient returns a new Client.
func NewRegistryClient(options ...Option) *client {
	client := &client{}

	// default transport
	WithTransport(http.DefaultTransport)(client)

	// apply functional options
	for _, opt := range options {
		opt(client)
	}

	return client
}

func (r *client) Digest(n image.Name) (image.Digest, error) {
	img, err := r.ReadRemoteImage(n)
	if err != nil {
		return image.EmptyDigest, err
	}

	hash, err := img.Digest()
	if err != nil {
		return image.EmptyDigest, err
	}

	return image.NewDigest(hash.String())
}

func (r *client) Copy(source image.Name, target image.Name) (image.Digest, int64, error) {
	img, err := r.ReadRemoteImage(source)
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to read image %v: %v", source, err)
	}

	sourceDigest, err := img.Digest()
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to read digest of image %v: %v", source, err)
	}

	targetDigest, s, err := img.Write(target)
	if err != nil {
		return image.EmptyDigest, 0, fmt.Errorf("failed to write image %v to %v: %v", source, target, err)
	}
	if sourceDigest != targetDigest {
		return image.EmptyDigest, 0, fmt.Errorf("failed to preserve digest of image %v: source digest %v, target digest %v", source, sourceDigest, targetDigest)
	}
	return targetDigest, s, err
}

func (r *client) NewLayout(path string) (registry.Layout, error) {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		if err := os.MkdirAll(path, outputDirPermissions); err != nil {
			return nil, err
		}
	}

	lp, err := layout.Write(path, empty.Index)
	if err != nil {
		return nil, err
	}

	return NewImageLayout(r, lp), nil
}

func (r *client) ReadLayout(path string) (registry.Layout, error) {
	lp, err := layout.FromPath(path)
	if err != nil {
		return nil, err
	}
	return NewImageLayout(r, lp), nil
}

func (r *client) ReadRemoteImage(n image.Name) (registry.Image, error) {
	return r.readRemoteImage(n)
}

func (r *client) NewImageFromManifest(img v1.Image) registry.Image {
	return newImageFromManifest(img, r.writeRemoteImage)
}

func (r *client) NewImageFromIndex(idx v1.ImageIndex) registry.Image {
	return newImageFromIndex(idx, r.writeRemoteIndex)
}
