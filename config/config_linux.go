//go:build linux || darwin || freebsd || openbsd || solaris

package config

import (
	"github.com/dean0731/dean-tool/exception"
	"github.com/dean0731/dean-tool/utils"
	"path/filepath"
)

func GetConfigDir(dir string) string {
	if home := utils.GetEnv(LinuxHomeEnv); home != "" {
		return filepath.Join(home, dir)
	}
	panic(exception.GetUserHomeDirError)
}
