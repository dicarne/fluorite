package main

import (
	"io"
	"os"
)

func ClearFolder(path string) {
	os.RemoveAll(path)
	os.MkdirAll(path, 0666)
}

func CopyFile(srcFile string, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}
