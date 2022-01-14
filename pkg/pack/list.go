package pack

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

func ListFileType(dir string) {
	results := map[string]int{}
	filepath.Walk(dir, func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() {
			ext := path.Ext(info.Name())
			if ext != "" {
				results[ext] += 1
			}
		}
		return nil
	})
	for k, v := range results {
		fmt.Println(k, v)
	}
}
