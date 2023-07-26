package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func CopyDir(srcPath, desPath string) error {
	srcPath = filepath.FromSlash(filepath.ToSlash(srcPath))
	desPath = filepath.FromSlash(filepath.ToSlash(desPath))
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			log.Fatalln("srcPath is not Dir")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			log.Fatalln("desPath is not Dir")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		log.Fatalln("srcPath is same as desPath! This is not allowed")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if path == srcPath {
			return nil
		}

		destNewPath := strings.Replace(path, srcPath, desPath, -1)
		if !f.IsDir() {
			_, err = CopyFile(path, destNewPath)
			IfFatal(err, "Copy file error")
		} else {
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			log.Fatalln("MakeDir failed:", err)
			return err
		}
	}
	return nil
}
