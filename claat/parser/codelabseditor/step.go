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
  //"fmt"
  "sort"
  "time"

  "github.com/googlecodelabs/tools/claat/types"
)


type step struct {
  Segments    []segmentContainer  `json:"segments"` // See decode.go and segment.go
  Title       string              `json:"title"`
  Duration    int                 `json:"duration"`
}

// parse returns a Step with the contents of all its segments
func (step step) parse() (*types.Step, error) {
  segments, err := decodeSegments(step.Segments)
  if err != nil {
    return nil, err
  }

  content := types.NewListNode()
  for _, segment := range segments {
    content.Append(segment.parse())
  }

  return &types.Step {
    Title: step.Title,
    Tags: environmentsForStep(segments), //Environment selected on a segment-by-segment basis
    Duration: time.Duration(step.Duration) * time.Minute,
    Content: content,
  }, nil
}

// environmentsForStep returns a normalized array of environments accumulated
// over the selected environments of each segment, unless there is a segment
// with no selected environment, in which case it returns an empty array.
// This is so that a step is included in the rendering for a given environment
// only when it has at least one segment which should display in that environment.
func environmentsForStep(segments []segment) []string {
  environments := []string{}
  for _, segment := range segments {
    selectedEnvironments := segment.selectedEnvironments()

    if len(selectedEnvironments) == 0 {
      return []string{}
    }
    environments = append(environments, selectedEnvironments...)
  }

  return normalizeEnvironments(environments)
}

// normalizeEnvironments sorts the environments and removes duplicates
// This is expected anywhere environments are specified
// including types.Step.Tags, types.Meta.Tags and types.Node.Env.
func normalizeEnvironments(environments []string) []string {
  sort.Strings(environments)

  uniqueEnvironments := []string{}
  for index, environment := range environments {
    if index == 0 || environment != environments[index - 1] {
      uniqueEnvironments = append(uniqueEnvironments, environment)
    }
  }

  return uniqueEnvironments
}
