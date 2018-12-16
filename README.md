# More Beam [![Build Status](https://img.shields.io/travis/bramp/morebeam.svg)](https://travis-ci.org/bramp/morebeam) [![Coverage](https://img.shields.io/coveralls/bramp/morebeam.svg)](https://coveralls.io/github/bramp/morebeam) [![Report card](https://goreportcard.com/badge/github.com/bramp/morebeam)](https://goreportcard.com/report/github.com/bramp/morebeam) [![GoDoc](https://godoc.org/bramp.net/morebeam?status.svg)](https://godoc.org/bramp.net/morebeam) [![Libraries.io](https://img.shields.io/librariesio/github/bramp/morebeam.svg)](https://libraries.io/github/bramp/morebeam)

Go package which provides additional functions useful when building [Apache Beam pipelines in Go](https://beam.apache.org/documentation/sdks/go/).

[Documentation available here](https://godoc.org/bramp.net/morebeam)

*This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.*

## How to use

Install:

```bash
go get bramp.net/morebeam/...
```

Example:

```go
import (
    "bramp.net/morebeam"
    "bramp.net/morebeam/csvio"
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

    // Return a new PCollection<KV<string, Painting>> where the key is the artist.
    paintingsByArtist := morebeam.AddKey(s, func(painting Painting) string {
        return painting.Artist
    }, paintings)

    debug.Print(s, paintingsByArtist)

    beamx.Run(context.Background(), p)
}
```

## Licence (Apache 2)

```
Copyright 2018 Google LLC
//
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at
//
   http://www.apache.org/licenses/LICENSE-2.0
//
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```