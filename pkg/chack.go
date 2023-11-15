package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

// 检查文件路径
func CheckFilePath(path string) error {
	// 检查文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) { // 文件不存在
			dirPath := filepath.Dir(path)
			err := CheckDirPath(dirPath) // 检查目录
			if err != nil {
				return err
			}
			file, createErr := os.Create(path) // 创建文件
			if createErr != nil {
				return createErr
			}
			file.Close()
		} else { // 存在，检查是否可读
			file, err := os.Open(path)
			if err != nil { // 不可读
				return err
			}
			defer file.Close()
		}
	}

	return nil
}

func CheckDirPath(dirPath string) error {
	stat, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 目录不存在，创建目录
			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("无法访问目录：%s", dirPath)
		}

	} else if !stat.IsDir() { // 检查是否为目录
		return fmt.Errorf("不是有效的目录路径：%s", dirPath)
	}

	return nil
}
