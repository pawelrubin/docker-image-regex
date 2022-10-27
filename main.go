package main

import (
	"fmt"
	"github.com/distribution/distribution/reference"
	"strings"
)

func main() {
	var imageUriRegexpString = reference.ReferenceRegexp.String()

	// make the regex compatible with the Python's re module
	imageUriRegexpString = strings.Replace(imageUriRegexpString, "/", "\\/", -1)
	imageUriRegexpString = strings.Replace(imageUriRegexpString, "[[:xdigit:]]", "[0-9a-fA-F]", -1)

	fmt.Printf("%s", imageUriRegexpString)
}
