// Code generated by counterfeiter. DO NOT EDIT.
package registryclientfakes

import (
	sync "sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	image "github.com/cnabio/image-relocation/pkg/image"
	registry "github.com/cnabio/image-relocation/pkg/registry"
	ggcr "github.com/cnabio/image-relocation/pkg/registry/ggcr"
)

type FakeRegistryClient struct {
	NewImageFromIndexStub        func(v1.ImageIndex) registry.Image
	newImageFromIndexMutex       sync.RWMutex
	newImageFromIndexArgsForCall []struct {
		arg1 v1.ImageIndex
	}
	newImageFromIndexReturns struct {
		result1 registry.Image
	}
	newImageFromIndexReturnsOnCall map[int]struct {
		result1 registry.Image
	}
	NewImageFromManifestStub        func(v1.Image) registry.Image
	newImageFromManifestMutex       sync.RWMutex
	newImageFromManifestArgsForCall []struct {
		arg1 v1.Image
	}
	newImageFromManifestReturns struct {
		result1 registry.Image
	}
	newImageFromManifestReturnsOnCall map[int]struct {
		result1 registry.Image
	}
	ReadRemoteImageStub        func(image.Name) (registry.Image, error)
	readRemoteImageMutex       sync.RWMutex
	readRemoteImageArgsForCall []struct {
		arg1 image.Name
	}
	readRemoteImageReturns struct {
		result1 registry.Image
		result2 error
	}
	readRemoteImageReturnsOnCall map[int]struct {
		result1 registry.Image
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistryClient) NewImageFromIndex(arg1 v1.ImageIndex) registry.Image {
	fake.newImageFromIndexMutex.Lock()
	ret, specificReturn := fake.newImageFromIndexReturnsOnCall[len(fake.newImageFromIndexArgsForCall)]
	fake.newImageFromIndexArgsForCall = append(fake.newImageFromIndexArgsForCall, struct {
		arg1 v1.ImageIndex
	}{arg1})
	fake.recordInvocation("NewImageFromIndex", []interface{}{arg1})
	fake.newImageFromIndexMutex.Unlock()
	if fake.NewImageFromIndexStub != nil {
		return fake.NewImageFromIndexStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newImageFromIndexReturns
	return fakeReturns.result1
}

func (fake *FakeRegistryClient) NewImageFromIndexCallCount() int {
	fake.newImageFromIndexMutex.RLock()
	defer fake.newImageFromIndexMutex.RUnlock()
	return len(fake.newImageFromIndexArgsForCall)
}

func (fake *FakeRegistryClient) NewImageFromIndexCalls(stub func(v1.ImageIndex) registry.Image) {
	fake.newImageFromIndexMutex.Lock()
	defer fake.newImageFromIndexMutex.Unlock()
	fake.NewImageFromIndexStub = stub
}

func (fake *FakeRegistryClient) NewImageFromIndexArgsForCall(i int) v1.ImageIndex {
	fake.newImageFromIndexMutex.RLock()
	defer fake.newImageFromIndexMutex.RUnlock()
	argsForCall := fake.newImageFromIndexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistryClient) NewImageFromIndexReturns(result1 registry.Image) {
	fake.newImageFromIndexMutex.Lock()
	defer fake.newImageFromIndexMutex.Unlock()
	fake.NewImageFromIndexStub = nil
	fake.newImageFromIndexReturns = struct {
		result1 registry.Image
	}{result1}
}

func (fake *FakeRegistryClient) NewImageFromIndexReturnsOnCall(i int, result1 registry.Image) {
	fake.newImageFromIndexMutex.Lock()
	defer fake.newImageFromIndexMutex.Unlock()
	fake.NewImageFromIndexStub = nil
	if fake.newImageFromIndexReturnsOnCall == nil {
		fake.newImageFromIndexReturnsOnCall = make(map[int]struct {
			result1 registry.Image
		})
	}
	fake.newImageFromIndexReturnsOnCall[i] = struct {
		result1 registry.Image
	}{result1}
}

func (fake *FakeRegistryClient) NewImageFromManifest(arg1 v1.Image) registry.Image {
	fake.newImageFromManifestMutex.Lock()
	ret, specificReturn := fake.newImageFromManifestReturnsOnCall[len(fake.newImageFromManifestArgsForCall)]
	fake.newImageFromManifestArgsForCall = append(fake.newImageFromManifestArgsForCall, struct {
		arg1 v1.Image
	}{arg1})
	fake.recordInvocation("NewImageFromManifest", []interface{}{arg1})
	fake.newImageFromManifestMutex.Unlock()
	if fake.NewImageFromManifestStub != nil {
		return fake.NewImageFromManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newImageFromManifestReturns
	return fakeReturns.result1
}

func (fake *FakeRegistryClient) NewImageFromManifestCallCount() int {
	fake.newImageFromManifestMutex.RLock()
	defer fake.newImageFromManifestMutex.RUnlock()
	return len(fake.newImageFromManifestArgsForCall)
}

func (fake *FakeRegistryClient) NewImageFromManifestCalls(stub func(v1.Image) registry.Image) {
	fake.newImageFromManifestMutex.Lock()
	defer fake.newImageFromManifestMutex.Unlock()
	fake.NewImageFromManifestStub = stub
}

func (fake *FakeRegistryClient) NewImageFromManifestArgsForCall(i int) v1.Image {
	fake.newImageFromManifestMutex.RLock()
	defer fake.newImageFromManifestMutex.RUnlock()
	argsForCall := fake.newImageFromManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistryClient) NewImageFromManifestReturns(result1 registry.Image) {
	fake.newImageFromManifestMutex.Lock()
	defer fake.newImageFromManifestMutex.Unlock()
	fake.NewImageFromManifestStub = nil
	fake.newImageFromManifestReturns = struct {
		result1 registry.Image
	}{result1}
}

func (fake *FakeRegistryClient) NewImageFromManifestReturnsOnCall(i int, result1 registry.Image) {
	fake.newImageFromManifestMutex.Lock()
	defer fake.newImageFromManifestMutex.Unlock()
	fake.NewImageFromManifestStub = nil
	if fake.newImageFromManifestReturnsOnCall == nil {
		fake.newImageFromManifestReturnsOnCall = make(map[int]struct {
			result1 registry.Image
		})
	}
	fake.newImageFromManifestReturnsOnCall[i] = struct {
		result1 registry.Image
	}{result1}
}

func (fake *FakeRegistryClient) ReadRemoteImage(arg1 image.Name) (registry.Image, error) {
	fake.readRemoteImageMutex.Lock()
	ret, specificReturn := fake.readRemoteImageReturnsOnCall[len(fake.readRemoteImageArgsForCall)]
	fake.readRemoteImageArgsForCall = append(fake.readRemoteImageArgsForCall, struct {
		arg1 image.Name
	}{arg1})
	fake.recordInvocation("ReadRemoteImage", []interface{}{arg1})
	fake.readRemoteImageMutex.Unlock()
	if fake.ReadRemoteImageStub != nil {
		return fake.ReadRemoteImageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readRemoteImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistryClient) ReadRemoteImageCallCount() int {
	fake.readRemoteImageMutex.RLock()
	defer fake.readRemoteImageMutex.RUnlock()
	return len(fake.readRemoteImageArgsForCall)
}

func (fake *FakeRegistryClient) ReadRemoteImageCalls(stub func(image.Name) (registry.Image, error)) {
	fake.readRemoteImageMutex.Lock()
	defer fake.readRemoteImageMutex.Unlock()
	fake.ReadRemoteImageStub = stub
}

func (fake *FakeRegistryClient) ReadRemoteImageArgsForCall(i int) image.Name {
	fake.readRemoteImageMutex.RLock()
	defer fake.readRemoteImageMutex.RUnlock()
	argsForCall := fake.readRemoteImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistryClient) ReadRemoteImageReturns(result1 registry.Image, result2 error) {
	fake.readRemoteImageMutex.Lock()
	defer fake.readRemoteImageMutex.Unlock()
	fake.ReadRemoteImageStub = nil
	fake.readRemoteImageReturns = struct {
		result1 registry.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryClient) ReadRemoteImageReturnsOnCall(i int, result1 registry.Image, result2 error) {
	fake.readRemoteImageMutex.Lock()
	defer fake.readRemoteImageMutex.Unlock()
	fake.ReadRemoteImageStub = nil
	if fake.readRemoteImageReturnsOnCall == nil {
		fake.readRemoteImageReturnsOnCall = make(map[int]struct {
			result1 registry.Image
			result2 error
		})
	}
	fake.readRemoteImageReturnsOnCall[i] = struct {
		result1 registry.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newImageFromIndexMutex.RLock()
	defer fake.newImageFromIndexMutex.RUnlock()
	fake.newImageFromManifestMutex.RLock()
	defer fake.newImageFromManifestMutex.RUnlock()
	fake.readRemoteImageMutex.RLock()
	defer fake.readRemoteImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistryClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ggcr.RegistryClient = new(FakeRegistryClient)
