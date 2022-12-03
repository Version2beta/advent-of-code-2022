package futils

import (
	"bufio"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Lines(path string) []string {
	f, err := os.Open(path)
	checkErr(err)
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	result := []string{}
	for fs.Scan() {
		result = append(result, fs.Text())
	}

	f.Close()
	return result
}
