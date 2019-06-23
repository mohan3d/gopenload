# gopenload
Golang client of the [openload.co](https://openload.co/) service. 

[![Build Status](https://travis-ci.org/mohan3d/gopenload.svg?branch=master)](https://travis-ci.org/mohan3d/gopenload)
[![Go Report Card](https://goreportcard.com/badge/github.com/mohan3d/gopenload?branch=master)](https://goreportcard.com/report/github.com/mohan3d/gopenload)

# Installation

```bash
$ go get github.com/mohan3d/gopenload
```

# Usage

implemented [API](https://openload.co/api) features.

**Retrieve account info**
```golang
package main

import (
	"fmt"

	"github.com/mohan3d/gopenload/openload"
)

func main() {
	// Create a client.
	client := openload.New("<LOGIN>", "<KEY>", nil)

	// Get account info.
	info, err := client.AccountInfo()

	if err != nil {
		panic(err)
	}
	fmt.Println(info.Email)
	fmt.Println(info.SignupAt)
}
```

**Upload file**
```golang
package main

import (
	"fmt"

	"github.com/mohan3d/gopenload/openload"
)

func main() {
	client := openload.New("<LOGIN>", "<KEY>", nil)
	uploaded, err := client.Upload("/path/dummyfile.txt", "", "", false)

	if err != nil {
		panic(err)
	}
	fmt.Println(uploaded.URL)
	fmt.Println(uploaded.ID)
	fmt.Println(uploaded.Size)
}
```

**Retrieve file info**
```golang
package main

import (
	"fmt"

	"github.com/mohan3d/gopenload/openload"
)

func main() {
	client := openload.New("<LOGIN>", "<KEY>", nil)
	info, err := client.FileInfo("uxbligkQAiN")

	if err != nil {
		panic(err)
	}
	fmt.Println(info.Name)
	fmt.Println(info.Size)
	fmt.Println(info.Status)
}
```