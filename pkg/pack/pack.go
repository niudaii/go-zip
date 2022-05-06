package pack

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// srcFile could be a single file or a directory
func Zip(srcFile, destZip, excludeSuffix string) error {
	excludes := strings.Split(excludeSuffix, ",")
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath
		if info.IsDir() {
			header.Name += "/"
		} else {
			if info.Name() == destZip {
				return nil
			}
			ext := path.Ext(info.Name())
			for _, exclude := range excludes {
				if exclude == ext {
					return nil
				}
			}
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(filepath)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
