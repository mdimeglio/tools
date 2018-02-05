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
  "fmt"
  "io"

  "github.com/googlecodelabs/tools/claat/parser"
  "github.com/googlecodelabs/tools/claat/types"
)

type Parser struct {
}

type snapshot struct {
  Codelab codelab       `json:"codelab"`
  Author  string        `json:"author"`
  Title   string        `json:"title"`
  Time    int64         `json:"time"`
}

type codelab struct {
  Steps     []step    `json:"steps"`    // see step.go
  Metadata  metadata  `json:"metadata"` // see metadata.go
}

func init() {
  parser.Register("codelabseditor", &Parser{})
}

func (p *Parser) Parse(r io.Reader) (*types.Codelab, error) {
  snapshot, err := decodeSnapshot(r)
  if err != nil {
    return nil, err
  }
  return (*snapshot).parse()
}

func (p *Parser) ParseFragment(r io.Reader) ([]types.Node, error) {
  return nil, fmt.Errorf("Not implemented")
}


func (snapshot snapshot) parse() (*types.Codelab, error) {
  codelab := snapshot.Codelab

  parsedMetadata, err := codelab.Metadata.parse(codelab.Steps, snapshot.Title, snapshot.Author)
  if err != nil {
    return nil, err
  }

  fmt.Printf("Parsed metadata: %+v\n", parsedMetadata)

  var parsedSteps []*types.Step
  for _, step := range codelab.Steps {
    parsedStep, err := step.parse()
    if err != nil {
      return nil, err
    }

    fmt.Printf("Parsed step: %+v\n", parsedStep)
    parsedSteps = append(parsedSteps, parsedStep)
  }

  fmt.Printf("Finished parsing codelab!\n")

  return &types.Codelab{
    Meta:  parsedMetadata,
    Steps: parsedSteps,
  }, nil
}