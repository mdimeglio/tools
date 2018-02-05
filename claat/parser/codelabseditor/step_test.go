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
  "testing"
  "reflect"
  "time"
  "github.com/googlecodelabs/tools/claat/types"
)

type stepTest struct {
  stepIn step
  stepExpected *types.Step
}

func StepTests() []stepTest {
  return []stepTest{
    {
      stepIn: step{
        Title: "A title",
        Duration: 20,
        Segments: getContainers(allSegmentTests()),
      },
      stepExpected: &types.Step{
        Title: "A title",
        Tags: []string{

        },
        Duration: time.Duration(20) * time.Minute,
        Content: types.NewListNode(getNodesExpected(allSegmentTests())...),
      },
    },
  }
}

// TestParseSegment runs tests for all segment types.
// testSets should contain a testSet for every type of segment.
// Each test set should be defined in a file named <segmenttype>_test.go.
func TestParseStep(t *testing.T) {
  for _, test := range StepTests() {
    stepOut, err := test.stepIn.parse()
    if err != nil {
      t.Error(err)
      continue
    }

    // TODO: Swap to deep equaling the whole thing when figured out issue with
    // deep equaling ListNodes described in segment_test.go
    if stepOut.Title != test.stepExpected.Title {
      t.Errorf("Title: expected %s, got %s", test.stepExpected.Title, stepOut.Title);
    }

    if !reflect.DeepEqual(stepOut.Tags, test.stepExpected.Tags) {
      t.Errorf("Tags: expected %s, got %s", test.stepExpected.Tags, stepOut.Tags);
    }

    if !reflect.DeepEqual(stepOut.Duration, test.stepExpected.Duration) {
      t.Errorf("Duration: expected %s, got %s", test.stepExpected.Duration, stepOut.Duration);
    }
  }
}