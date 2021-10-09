package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

// download bilibili srt files first
// https://gitee.com/KGDKL/BiliCC-Srt/
func main() {
	filepath := "P19_en-US.srt"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	// caution: scanner.Scan() is limited to MaxScanTokenSize = 64K
	// if line is too long, need to resize the buffer
	// 	const maxCapacity = longLineLen  // your required line length
	// buf := make([]byte, maxCapacity)
	// scanner.Buffer(buf, maxCapacity)

	var b bytes.Buffer
	var cnt int
	for scanner.Scan() {
		if (cnt-2)%4 == 0 {
			if b.Len() != 0 {
				b.WriteByte(' ')
			}
			b.Write(scanner.Bytes())
		}
		cnt++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// open output file
	fo, err := os.Create("P19_en-US.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	fo.Write(b.Bytes())
}
