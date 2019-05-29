package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func main() {
	h := fnv.New64()
	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to receive input: %v\n", err)
		return
	}
	fmt.Printf("%d\n", int64(h.Sum64()))
}
