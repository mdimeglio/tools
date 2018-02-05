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
  "reflect"
  "strings"
  "testing"
  "encoding/json"
)

func TestDecodeSnapshot(t *testing.T) {
  codelabFileJson := `
  {
    "appId": "abcde",
    "serverRevision": 1,
    "data": {
      "id": "root",
      "type": "Map",
      "value": {
        "_snapshots": {
          "id": "VkO1rvHbz5YH",
          "type": "Snapshots",
          "value": {
            "_snapshot": {
              "json": {
                "codelab": {
                  "steps": [
                    {
                      "segments": [
                        {
                          "type": "CodeSegment",
                          "segment": {
                            "selectedEnvironments": [
                              "Kiosk"
                            ],
                            "codeSnippet": "code 1",
                            "caption": "file 1",
                            "link": "link1"
                          }
                        },
                        {
                          "type": "CodeSegment",
                          "segment": {
                            "selectedEnvironments": [
                              "Web",
                              "Kiosk"
                            ],
                            "codeSnippet": "code2",
                            "caption": "file2",
                            "link": "link2"
                          }
                        },
                        {
                          "type": "CodeSegment",
                          "segment": {
                            "selectedEnvironments": [
                              "Web"
                            ],
                            "codeSnippet": "code3",
                            "caption": "file3",
                            "link": "link3"
                          }
                        }
                      ],
                      "title": "Step 1",
                      "duration": 6
                    },
                    {
                      "segments": [
                        {
                          "type": "CodeSegment",
                          "segment": {
                            "selectedEnvironments": [],
                            "codeSnippet": "code 4",
                            "caption": "File 4",
                            "link": "link 4"
                          }
                        }
                      ],
                      "title": "Step 2",
                      "duration": 1
                    },
                    {
                      "segments": [

                      ],
                      "title": "Step 3",
                      "duration": 0
                    }
                  ],
                  "metadata": {
                    "summary": "A sample codelab",
                    "relativeURL": "sample2",
                    "feedbackURL": "feedback.com",
                    "analyticsAccountID": "lkjhbv",
                    "category": "Ads",
                    "status": "Published",
                    "environments": [
                      "Web",
                      "Kiosk"
                    ]
                  }
                },
                "title": "Untitled",
                "author": "Joe Bloggs",
                "time": 1518046576888
              }
            }
          }
        }
      }
    }
  }`

  snapshotExpected := &snapshot{
    Codelab: codelab{
      Steps: []step{
        step{
          Title: "Step 1",
          Duration: 6,
          Segments: []segmentContainer{
            segmentContainer{
              Type: "CodeSegment",
              JSON: json.RawMessage(`{
                            "selectedEnvironments": [
                              "Kiosk"
                            ],
                            "codeSnippet": "code 1",
                            "caption": "file 1",
                            "link": "link1"
                          }`,
              ),
            },
            segmentContainer{
              Type: "CodeSegment",
              JSON: json.RawMessage(`{
                            "selectedEnvironments": [
                              "Web",
                              "Kiosk"
                            ],
                            "codeSnippet": "code2",
                            "caption": "file2",
                            "link": "link2"
                          }`,
              ),
            },
            segmentContainer{
              Type: "CodeSegment",
              JSON: json.RawMessage(`{
                            "selectedEnvironments": [
                              "Web"
                            ],
                            "codeSnippet": "code3",
                            "caption": "file3",
                            "link": "link3"
                          }`,
              ),
            },
          },
        },
        step{
          Title: "Step 2",
          Duration: 1,
          Segments: []segmentContainer{
            segmentContainer{
              Type: "CodeSegment",
              JSON: json.RawMessage(`{
                            "selectedEnvironments": [],
                            "codeSnippet": "code 4",
                            "caption": "File 4",
                            "link": "link 4"
                          }`,
              ),
            },
          },
        },
        step{
          Title: "Step 3",
          Duration: 0,
          Segments: []segmentContainer{},
        },
      },
      Metadata: metadata{
        Summary: "A sample codelab",
        ID: "sample2",
        FeedbackURL: "feedback.com",
        AnalyticsAccountID:"lkjhbv",
        Category: "Ads",
        Status: "Published",
        Environments: []string{
          "Web",
          "Kiosk",
        },
      },
    },
    Author: "Joe Bloggs",
    Title: "Untitled",
    Time: 1518046576888,
  }


  reader := strings.NewReader(codelabFileJson)
  snapshotOut, err := decodeSnapshot(reader)
  if err != nil {
    t.Error(err)
  }

  if !reflect.DeepEqual(snapshotExpected, snapshotOut) {
    t.Errorf("decodeSnapshot(s) == %+v.\n Expected %+v\n", snapshotOut, snapshotExpected)
  }
}


func TestDecodeSnapshotNoSteps(t *testing.T) {
  codelabFileJson := `
  {
    "appId": "abcde",
    "serverRevision": 1,
    "data": {
      "id": "root",
      "type": "Map",
      "value": {
        "_snapshots": {
          "id": "VkO1rvHbz5YH",
          "type": "Snapshots",
          "value": {
            "_snapshot": {
              "json": {
                "codelab": {
                  "steps": [],
                  "metadata": {
                    "summary": "A sample codelab",
                    "relativeURL": "sample2",
                    "feedbackURL": "feedback.com",
                    "analyticsAccountID": "lkjhbv",
                    "category": "Ads",
                    "status": "Published",
                    "environments": [
                      "Web",
                      "Kiosk"
                    ]
                  }
                },
                "title": "Untitled",
                "author": "Joe Bloggs",
                "time": 1518046576888
              }
            }
          }
        }
      }
    }
  }`

  snapshotExpected := &snapshot{
    Codelab: codelab{
      Steps: []step{},
      Metadata: metadata{
        Summary: "A sample codelab",
        ID: "sample2",
        FeedbackURL: "feedback.com",
        AnalyticsAccountID:"lkjhbv",
        Category: "Ads",
        Status: "Published",
        Environments: []string{
          "Web",
          "Kiosk",
        },
      },
    },
    Author: "Joe Bloggs",
    Title: "Untitled",
    Time: 1518046576888,
  }


  reader := strings.NewReader(codelabFileJson)
  snapshotOut, err := decodeSnapshot(reader)
  if err != nil {
    t.Error(err)
  }

  if !reflect.DeepEqual(snapshotExpected, snapshotOut) {
    t.Errorf("decodeSnapshot(reader) == %+v.\n Expected %+v\n", snapshotOut, snapshotExpected)
  }
}

func TestDecodeSegment(t *testing.T) {
  for _, test := range allSegmentTests() {
    segmentOut, err := test.container.decode()
    if err != nil {
      t.Errorf("Test: %+v, Error: %+v", test, err)
      continue
    }

    if !reflect.DeepEqual(test.segmentIn, segmentOut) {
      t.Errorf("container.decode() == %+v.\n Expected %+v\n", segmentOut, test.segmentIn)
    }
  }
}

func TestDecodeSegments(t *testing.T) {
  containersIn := getContainers(allSegmentTests())
  segmentsExpected := getSegmentsIn(allSegmentTests())

  segmentsOut, err := decodeSegments(containersIn)
  if err != nil {
    t.Error(err)
  }

  if !reflect.DeepEqual(segmentsExpected, segmentsOut) {
    t.Errorf("decodeSegments(containersIn) == %+v.\n Expected %+v\n", segmentsOut, segmentsExpected)
  }
}