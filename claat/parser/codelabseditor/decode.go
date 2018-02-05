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
  "io"
  "encoding/json"
)

// Types only used for the JSON decoding:

type document struct {
  AppId     string          `json:"appId"`
  Revision  int             `json:"revision"`
  Data      rootContainer   `json:"data"`
}

type rootContainer struct {
  Id    string    `json:"id"`
  Type  string    `json:"type"`
  Value root      `json:"value"`
}

type root struct {
  Snapshots snapshotsContainer `json:"_snapshots"`
}

type snapshotsContainer struct {
  Id    string    `json:"id"`
  Type  string    `json:"type"`
  Value snapshots `json:"value"`
}

type snapshots struct {
  Snapshot snapshotContainer `json:"_snapshot"`
}

type snapshotContainer struct {
  Value *snapshot `json:"json"`
}

// segmentContainer allows choice between different segment types.
// JSON should be decoded to a segment subtype determined by Type.
type segmentContainer struct {
  Type  string          `json:"type"`
  JSON  json.RawMessage `json:"segment"`
}

// decodeSnapshot extracts the Codelab snapshot out of
// the exported JSON of a realtime codelabs-editor document.
// The root of the codelab snapshot is not the root of the JSON file.
func decodeSnapshot(r io.Reader) (*snapshot, error) {
  decoder := json.NewDecoder(r)

  var document document
  err := decoder.Decode(&document)
  if err != nil {
    return nil, err
  }

  return document.Data.Value.Snapshots.Value.Snapshot.Value, nil
}

// decode converts a segmentContainer into a segment of the correct type.
// json.Unmarshal can only unmarshal to concrete types. Hence we must repeat the unmarshal code
// for each concrete type we wish to unmarshal.
func (segmentContainer segmentContainer) decode() (segment, error) {
  switch (segmentContainer.Type) {
  case "CodeSegment":
    var segment codeSegment
    err := json.Unmarshal(segmentContainer.JSON, &segment)
    return segment, err
  default:
    var segment unimplementedSegment
    err := json.Unmarshal(segmentContainer.JSON, &segment.abstractSegment)
    segment.message = segmentContainer.Type
    return segment, err
  }
}

// decodeSegments convers an array of segment containers to an array of
// segments of the correct type.
func decodeSegments(segmentContainers []segmentContainer) ([]segment, error) {
  var segments []segment
  for _, segmentContainer := range segmentContainers {
    segment, err := segmentContainer.decode()
    if err != nil {
      return nil, err
    }

    segments = append(segments, segment)
  }
  return segments, nil
}
