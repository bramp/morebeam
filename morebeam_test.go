// Copyright 2018 Google LLC
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package morebeam_test

import (
	"context"
	"flag"
	"reflect"

	"bramp.net/morebeam"
	"bramp.net/morebeam/csvio"
	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/go/pkg/beam/x/debug"
)

// Painting represents a single record in the csv file.
type Painting struct {
	Artist  string `csv:"artist"`
	Title   string `csv:"title"`
	Year    int    `csv:"year"`
	NotUsed string `csv:"-"` // Ignored field
}

func Example() {
	flag.Parse()
	beam.Init()

	p, s := beam.NewPipelineWithRoot()

	// Read the CSV file as a PCollection<Painting>.
	paintings := csvio.Read(s, "paintings.csv", reflect.TypeOf(Painting{}))

	// Reshuffle the CSV output to improve parallelism.
	paintings = morebeam.Reshuffle(s, paintings)

	// Return a new PCollection<KV<string, Painting>> where the key is the artist.
	paintingsByArtist := morebeam.AddKey(s, func(painting Painting) string {
		return painting.Artist
	}, paintings)

	debug.Print(s, paintingsByArtist)

	beamx.Run(context.Background(), p)

	// Output:
}
