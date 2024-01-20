# CutInfo

CutInfo is a lightweight Go library designed to extract information from text based on specified start and end tags. It is a simple and efficient tool for retrieving content between two reference points in a text string.

## Features

- Simple and easy-to-use API
- Lightweight and fast
- Retrieves content between specified start and end tags
- Handles cases with multiple occurrences of end tags by selecting content up to the first occurrence

## Objective

The main goal of CutInfo is to provide a quick and easy solution for extracting information from text strings. It is particularly useful for scenarios where you need to parse HTML or similar structures and retrieve specific content between tags.

## How to Install

To use CutInfo in your Go project, you can simply run the following command:

```bash
go get -u github.com/simplyYan/cutinfo
```

## How to use
```go
package main

import (
	"fmt"
	"github.com/simplyYan/cutinfo"
)

func main() {
    ci := cutinfo.New()

    text := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
    <p class="test">Hello World</p>
</body>
</html>
`

    result := ci.Target(text, '<p class="test">', '</p>')
    fmt.Println(result) // Output: Hello World
}
```
1. Import the cutinfo package.
2. Create a new CutInfo instance using cutinfo.New().
3. Use the Target method to extract content between specified start and end tags.
