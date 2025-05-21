package utils

import (
	"io"
	"os"
	"os/exec"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func ExecuteCommandToFile(command string, outputFilePath string, args ...string) error {
	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	cmd := exec.Command(command, args...)

	cmd.Stdout = outputFile
	cmd.Stderr = outputFile

	// 运行命令
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func ExecuteCommandToNull(command string, dirPath string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dirPath
	// 都会丢弃
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	// 运行命令
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
