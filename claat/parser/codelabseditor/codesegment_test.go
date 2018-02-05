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

// getCodeSegmentTests returns an array of tests for codesegments
// See segment_test.go for definition of segmentTest and actual tests.
func codeSegmentTests() []segmentTest {
  return []segmentTest{
    {
      container: segmentContainer{
        Type: "CodeSegment",
        JSON: json.RawMessage(`{
            "selectedEnvironments": [],
            "codeSnippet": "code 1",
            "caption": "file 1",
            "link": "link1",
            "language": "C"
          }`,
        ),
      },
      segmentIn: codeSegment{
        abstractSegment: abstractSegment{
          SelectedEnvironments: []string{},
        },
        CodeSnippet: "code 1",
        Caption: "file 1",
        Link: "link1",
        Language: "C",
      },
      nodesExpected: types.NewListNode(
        types.NewHeaderNode(
          3,
          types.NewURLNode(
            "link1",
            types.NewTextNode("file 1"),
          ),
        ),
        types.NewCodeNode("code 1", false),
      ),
    },
    {
      container: segmentContainer{
        Type: "CodeSegment",
        JSON: json.RawMessage(`{
            "selectedEnvironments": [
              "Kiosk"
            ],
            "codeSnippet": "code 1",
            "caption": "file 1",
            "link": "link1",
            "language": "C"
          }`,
        ),
      },
      segmentIn: codeSegment{
        abstractSegment: abstractSegment{
          SelectedEnvironments: []string{
            "Kiosk",
          },
        },
        CodeSnippet: "code 1",
        Caption: "file 1",
        Link: "link1",
        Language: "C",
      },
      nodesExpected: withEnvironments(
        types.NewListNode(
          types.NewHeaderNode(
            3,
            types.NewURLNode(
              "link1",
              types.NewTextNode("file 1"),
            ),
          ),
          types.NewCodeNode("code 1", false),
        ),
        []string{"kiosk"},
      ),
    },
    {
      container: segmentContainer{
        Type: "CodeSegment",
        JSON: json.RawMessage(`{
            "selectedEnvironments": [],
            "codeSnippet": "code 2"
          }`,
        ),
      },
      segmentIn: codeSegment{
        abstractSegment: abstractSegment{
          SelectedEnvironments: []string{},
        },
        CodeSnippet: "code 2",
        Caption: "",
        Link: "",
        Language: "",
      },
      nodesExpected: types.NewListNode(
        types.NewCodeNode("code 2", false),
      ),
    },
    {
      container: segmentContainer{
        Type: "CodeSegment",
        JSON: json.RawMessage(`{
            "selectedEnvironments": [],
            "codeSnippet": "code 3",
            "caption": "Hello"
          }`,
        ),
      },
      segmentIn: codeSegment{
        abstractSegment: abstractSegment{
          SelectedEnvironments: []string{},
        },
        CodeSnippet: "code 3",
        Caption: "Hello",
        Link: "",
        Language: "",
      },
      nodesExpected: types.NewListNode(
        types.NewHeaderNode(
          3,
          types.NewTextNode("Hello"),
        ),
        types.NewCodeNode("code 3", false),
      ),
    },
  }
}