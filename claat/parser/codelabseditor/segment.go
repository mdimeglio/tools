// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codelabseditor

import (
  "github.com/googlecodelabs/tools/claat/types"
  "strings"
)

type segment interface {
  parse() *types.ListNode
  selectedEnvironments() []string
}

// abstractSegment encapsulates functionality common to all segments.
// Concrete segment instances should embed abstractSegment
// All concrete segment's parse() function should use the ListNode
// returned by calling parse() on its embedded abstract segment.
type abstractSegment struct {
  SelectedEnvironments  []string  `json:"selectedEnvironments"`
}

func (segment abstractSegment) parse() *types.ListNode {
  content := types.NewListNode()
  content.MutateEnv(segment.selectedEnvironments())
  return content
}

func (segment abstractSegment) selectedEnvironments() []string {
  for index, environment  := range segment.SelectedEnvironments {
    segment.SelectedEnvironments[index] = strings.ToLower(environment)
  }
  return segment.SelectedEnvironments
}