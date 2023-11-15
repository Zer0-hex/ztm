package path

import (
	"fmt"
	"os"
)

func CheckPath(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 获取权限信息
	permissions := fileInfo.Mode().Perm()

	// 打印权限信息
	fmt.Printf("路径 %s 的权限：%04o\n", path, permissions)
}
