/*
 * Copyright (c) 2020-Present Pivotal Software, Inc. All rights reserved.
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

package images

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/cnabio/image-relocation/pkg/image"
)

// Set is an immutable collection of image references without duplicates.
type Set struct {
	m map[image.Name]struct{} // wrap map in a struct to avoid exposing map operations
}

// Empty is a Set with no image references.
var Empty Set

func init() {
	Empty = empty()
}

func empty() Set {
	return Set{
		m: make(map[image.Name]struct{}),
	}
}

// New constructs a Set from some image references.
func New(ss ...string) (Set, error) {
	set := empty()
	for _, s := range ss {
		name, err := image.NewName(s)
		if err != nil {
			return Set{}, err
		}
		set.m[name] = struct{}{}
	}
	return set, nil
}

func (s Set) clone() Set {
	c := empty()
	for i := range s.m {
		c.m[i] = struct{}{}
	}
	return c
}

// Union returns the mathematical union of this Set with another Set.
func (s Set) Union(t Set) Set {
	u := s.clone()
	for i := range t.m {
		u.m[i] = struct{}{}
	}
	return u
}

// Slice returns this set as an unsorted slice of image references.
func (s Set) Slice() []image.Name {
	var result []image.Name
	for i := range s.m {
		result = append(result, i)
	}
	return result
}

// Strings returns the image references in the set as a sorted slice of strings.
func (s Set) Strings() []string {
	var result []string
	for i := range s.m {
		result = append(result, i.String())
	}
	sort.Strings(result)
	return result
}

// String returns a sorted, string representation of the set.
func (s Set) String() string {
	return fmt.Sprintf("[%s]", strings.Join(s.Strings(), ", "))
}

// MarshalJSON encodes this Set as a JSON array of image references.
func (s Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Strings())
}

// UnmarshalJSON decodes a JSON array of image references into a Set.
func (s *Set) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v == nil {
		*s = Empty
		return nil
	}

	refs, ok := v.([]interface{})
	if !ok {
		return fmt.Errorf("unmarshalled data not a slice: %v", v)
	}

	var strs []string
	for _, ref := range refs {
		x, ok := ref.(string)
		if !ok {
			return fmt.Errorf("unmarshalled slice contains a value which is not a string: %v", x)
		}
		strs = append(strs, x)
	}
	var err error
	*s, err = New(strs...)
	return err
}
