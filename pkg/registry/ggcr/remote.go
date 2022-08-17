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
	"errors"
	"net/http"

	"github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/cnabio/image-relocation/pkg/registry"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/cnabio/image-relocation/pkg/image"
)

var (
	resolveFunc        = authn.DefaultKeychain.Resolve
	repoGetFunc        = remote.Get
	repoWriteFunc      = remote.Write
	repoIndexWriteFunc = remote.WriteIndex
)

func readRemoteImage(mfstWriter manifestWriter, idxWriter indexWriter, transport http.RoundTripper, insecure bool) func(image.Name) (registry.Image, error) {
	return func(n image.Name) (registry.Image, error) {
		auth, err := resolve(n)
		if err != nil {
			return nil, err
		}

		if n.Tag() == "" && n.Digest() == image.EmptyDigest {
			// use default tag
			n, err = n.WithTag("latest")
			if err != nil {
				return nil, err
			}
		}

		parseOpts := []name.Option{name.StrictValidation}
		if insecure {
			parseOpts = append(parseOpts, name.Insecure)
		}

		ref, err := name.ParseReference(n.String(), parseOpts...)
		if err != nil {
			return nil, err
		}

		desc, err := repoGetFunc(ref, remote.WithAuth(auth), remote.WithTransport(transport))
		if err != nil {
			return nil, err
		}

		switch desc.MediaType {
		case types.OCIImageIndex, types.DockerManifestList:
			idx, err := desc.ImageIndex()
			if err != nil {
				return nil, err
			}
			return newImageFromIndex(idx, idxWriter), nil
		default:
			// assume all other media types are images since some images don't set the media type
		}
		img, err := desc.Image()
		if err != nil {
			return nil, err
		}

		return newImageFromManifest(img, mfstWriter), nil
	}
}

func writeRemoteImage(transport http.RoundTripper, insecure bool) func(v1.Image, image.Name) error {
	return func(i v1.Image, n image.Name) error {
		auth, err := resolve(n)
		if err != nil {
			return err
		}

		ref, err := getWriteReference(n, insecure)
		if err != nil {
			return err
		}

		return repoWriteFunc(ref, i, remote.WithAuth(auth), remote.WithTransport(transport))
	}
}

func writeRemoteIndex(transport http.RoundTripper, insecure bool) func(v1.ImageIndex, image.Name) error {
	return func(i v1.ImageIndex, n image.Name) error {
		auth, err := resolve(n)
		if err != nil {
			return err
		}

		ref, err := getWriteReference(n, insecure)
		if err != nil {
			return err
		}

		return repoIndexWriteFunc(ref, i, remote.WithAuth(auth), remote.WithTransport(transport))
	}
}

func resolve(n image.Name) (authn.Authenticator, error) {
	if n == image.EmptyName {
		return nil, errors.New("empty image name invalid")
	}
	repo, err := name.NewRepository(n.WithoutTagOrDigest().String(), name.WeakValidation)
	if err != nil {
		return nil, err
	}

	return resolveFunc(repo.Registry)
}

func getWriteReference(n image.Name, insecure bool) (name.Reference, error) {
	// if target image reference is both tagged and digested, ignore the digest so the tag is preserved
	// (the digest will be preserved by go-containerregistry)
	if n.Tag() != "" {
		n = n.WithoutDigest()
	}

	parseOpts := []name.Option{name.WeakValidation}
	if insecure {
		parseOpts = append(parseOpts, name.Insecure)
	}
	return name.ParseReference(n.String(), parseOpts...)
}
