package main

import (
	"fmt"

	"github.com/doniacld/hello/morestrings"
	"github.com/google/go-cmp/cmp"

)

func main() {
		helloWorld := "hello, world\n"
		fmt.Printf(helloWorld)
		reverseString := morestrings.ReverseRunes(helloWorld)
		fmt.Printf(reverseString)
}

