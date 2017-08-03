# media-type
[![Build Status](https://img.shields.io/travis/cssivision/media-type.svg?style=flat-square)](https://travis-ci.org/cssivision/media-type)
[![Coverage Status](http://img.shields.io/coveralls/cssivision/media-type.svg?style=flat-square)](https://coveralls.io/github/cssivision/media-type?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/cssivision/media-type/blob/master/LICENSE)

RFC 6838 media type parser.
parse a media type into type, subtype, and suffix, or format a normal media types string use those parts.

# Installation

```sh
go get github.com/cssivision/media-type
```

# Usage
## parse media type
```go
package main

import (
	"fmt"

	mediaType "github.com/cssivision/media-type"
)

func main() {
	mt, err := mediaType.Parse("application/json+xml")
	if err != nil {
		panic(err)
	}

	fmt.Println("type: ", mt.Type)
	fmt.Println("subtype: ", mt.SubType)
	fmt.Println("suffix: ", mt.Suffix)
}
```
## format media type
```go
package main

import (
	"fmt"

	mediaType "github.com/cssivision/media-type"
)

func main() {
	mt := &mediaType.MediaType{
		Type:    "application",
		SubType: "json",
		Suffix:  "xml",
	}
	str, err := mt.Format()
	if err != nil {
		panic(err)
	}
	fmt.Println("media type:", str)
}
```


# Licenses

All source code is licensed under the [MIT License](https://github.com/cssivision/media-type/blob/master/LICENSE).