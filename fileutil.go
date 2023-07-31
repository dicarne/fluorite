package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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

func DownloadFile(url string, filename string) {
	resp, err := http.Get(url)
	IfFatal(err, "download file error")
	defer resp.Body.Close()

	out, err := os.Create(filename)
	IfFatal(err, "create tmp file failed")
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	IfFatal(err, "save file error")
}

func UnzipFile(path string, dir string) {
	reader, err := zip.OpenReader("theme.zip")
	IfFatal(err, "unzip theme error")
	defer reader.Close()
	for _, file := range reader.File {
		if err := _unzipFile(file, dir); err != nil {
			IfFatal(err, "unzip error")
		}
	}
}

func _unzipFile(file *zip.File, dir string) error {
	// Prevent path traversal vulnerability.
	// Such as if the file name is "../../../path/to/file.txt" which will be cleaned to "path/to/file.txt".
	name := strings.TrimPrefix(filepath.Join(string(filepath.Separator), file.Name), string(filepath.Separator))
	filePath := path.Join(dir, name)

	// Create the directory of file.
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Open the file.
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	// Create the file.
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// Save the decompressed file content.
	_, err = io.Copy(w, r)
	return err
}
