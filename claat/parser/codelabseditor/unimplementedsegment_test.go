// Copyright 2018 Google Inc. All Rights Reserved.
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
  "encoding/json"
  "github.com/googlecodelabs/tools/claat/types"
)

// getUnimplementedSegmentTests returns an array of tests for unimplementedSegments
// See segment_test.go for definition of segmentTest and actual tests.
func unimplementedSegmentTests() []segmentTest {
  return []segmentTest{
    {
      container: segmentContainer{
        Type: "A message",
        JSON: json.RawMessage(`{
            "selectedEnvironments": []
          }`,
        ),
      },
      segmentIn: unimplementedSegment{
        abstractSegment: abstractSegment{
          SelectedEnvironments: []string{},
        },
        message: "A message",
      },
      nodesExpected: types.NewListNode(
        types.NewHeaderNode(2, types.NewTextNode("Unimplemented segment")),
        types.NewTextNode("A message"),
      ),
    },
  }
}