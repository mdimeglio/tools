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
  "github.com/googlecodelabs/tools/claat/types"
)

type segmentTest struct {
  container segmentContainer
  segmentIn segment
  nodesExpected *types.ListNode
}


// allSegmentTests returns all tests for all segment types
// testSets should contain a testSet for every type of segment.
// Each test set should be defined in a file named <segmenttype>_test.go.
func allSegmentTests() []segmentTest {
  testSets := [][]segmentTest{
    codeSegmentTests(),
    unimplementedSegmentTests(),
  }

  var tests []segmentTest
  for _, testSet := range testSets {
    tests = append(tests, testSet...)
  }

  return tests
}

// TestParseSegment runs tests for all segment types.
func TestParseSegment(t *testing.T) {
  for _, test := range allSegmentTests() {
    nodesOut := test.segmentIn.parse()

    // TODO: Tests don't pass if compare test.nodesExpected and nodesOut directly!
    // Passes using third party version of DeepEqual https://godoc.org/github.com/juju/testing/checkers#DeepEqual
    // Likely problem that DeepEqual considers nil and empty slices to be different
    if !reflect.DeepEqual(test.nodesExpected.Nodes, nodesOut.Nodes) {
      t.Errorf("segmentIn:\n%+v,\nsegmentIn.parse():\n%+v\n Expected\n%+v\n", test.segmentIn, nodesOut.Nodes, test.nodesExpected.Nodes)
    }
  }
}

// withEnvironments is a builder that adds environments to a ListNode and returns it
// useful when declaring segment tests for segments with environments
func withEnvironments(node *types.ListNode, environments []string) *types.ListNode {
  node.MutateEnv(environments)
  return node
}

// getContainers extracts the containers from an array of segment tests
func getContainers(segmentTests []segmentTest) []segmentContainer {
  var containers []segmentContainer
  for _, segmentTest := range segmentTests {
    containers = append(containers, segmentTest.container)
  }
  return containers
}

// getSegmentsIn extracts the segments from an array of segment tests
func getSegmentsIn(segmentTests []segmentTest) []segment {
  var segments []segment
  for _, segmentTest := range segmentTests {
    segments = append(segments, segmentTest.segmentIn)
  }
  return segments
}

// getNodesExpected extracts the parsed nodes from an array of segment tests
func getNodesExpected(segmentTests []segmentTest) []types.Node {
  var nodes []types.Node
  for _, segmentTest := range segmentTests {
    nodes = append(nodes, segmentTest.nodesExpected)
  }
  return nodes
}
