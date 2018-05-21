package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	fmt.Printf("%s %d %d %p \n", path, len(path), cap(path), &path)
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Printf("dir1 => %s %d %d %p \n", string(dir1), len(dir1), cap(dir1), &dir1) //prints: dir1 => AAAA
	fmt.Printf("dir2 => %s %d %d %p \n", string(dir2), len(dir2), cap(dir2), &dir2) //prints: dir2 => BBBBBBBBB
	fmt.Printf("/ %p \n", &path[sepIndex])

	dir1 = append(dir1, "suffix"...)

	fmt.Printf("path 1 => %s %p \n", string(path), &path)

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB (not ok)

	fmt.Println("new path =>", string(path))
}
