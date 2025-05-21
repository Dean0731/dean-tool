//go:build windows

package config

import (
	"github.com/dean0731/dean-tool/exception"
	"github.com/dean0731/dean-tool/utils"
	"os"
	"path/filepath"
)

func GetConfigDir(dir string) string {
	if home := utils.GetEnv(WindowsHomeEnv); home != "" {
		dir := filepath.Join(home, dir)
		os.MkdirAll(dir, 0755)
		return dir
	}
	panic(exception.GetUserHomeDirError)
}
