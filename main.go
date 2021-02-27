package main

import (
	"flag"
	"fmt"
	"strings"

	"modernc.org/mathutil"
)

const root = "/v1/asset/(.*)/resize"

type Matrix []string

func (m Matrix) Len() int { return len(m) }

func (m Matrix) Less(i, j int) bool { return m[i] < m[j] }

func (m Matrix) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func main() {
	pathPtr := flag.String("path", "", "Matrix variables path segment eg ;m=exact;h=([0-9]+);w=([0-9]+)")
	rootPathPtr := flag.String("rootPath", "/v1/asset/(.*)/resize", "Root path eg /v1/asset/(.*)/resize")
	flag.Parse()

	path := *rootPathPtr + *pathPtr

	paramsStr := strings.TrimLeft(path, root)

	params := Matrix(strings.Split(paramsStr, ";")[1:])

	mathutil.PermutationFirst(params)

	total := 1

	print(params)

	for mathutil.PermutationNext(params) {
		print(params)
		total++
	}

	fmt.Printf("\n\nFor a total of %d permutations\n", total)
}

func print(matrix Matrix) {
	fmt.Println(root + ";" + strings.Join(matrix, ";"))
}
