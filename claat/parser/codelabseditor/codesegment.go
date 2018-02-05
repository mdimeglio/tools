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
  "github.com/googlecodelabs/tools/claat/types"
)

type codeSegment struct {
  abstractSegment
  CodeSnippet           string  `json:"codeSnippet"`
  Caption               string  `json:"caption"`
  Link                  string  `json:"link"`
  Language              string  `json:"language"`
}

func (segment codeSegment) parse() *types.ListNode {
  content := segment.abstractSegment.parse()

  // Add caption to code if caption was defined
  // If the link was defined, it will exist in the output only if the caption was defined
  if (segment.Caption != "") {
    var captionContent types.Node
    captionContent = types.NewTextNode(segment.Caption)

    if (segment.Link != "") {
      captionContent = types.NewURLNode(segment.Link, captionContent)
    }

    captionContent = types.NewHeaderNode(3, captionContent)

    content.Append(captionContent)
  }

  content.Append(types.NewCodeNode(segment.CodeSnippet, false))
  return content
}
