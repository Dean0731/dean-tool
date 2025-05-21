package utils

import (
	"os"
	"path/filepath"
)

// CheckIfFileExistsIgnoringExt 检查指定目录下是否存在给定名称的文件，忽略文件扩展名。
func CheckIfFileExistsIgnoringExt(dir, filenameWithoutExt string) (bool, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if !file.IsDir() && stripExt(file.Name()) == filenameWithoutExt {
			return true, nil
		}
	}
	return false, nil
}

// stripExt 去除文件名中的扩展名。
func stripExt(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
