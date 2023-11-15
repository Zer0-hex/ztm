package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
)

func GetVersion(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return resp.Request.URL.String(), nil
}

func ChangeDir(path string) error {
	systype := runtime.GOOS
	var cmd *exec.Cmd
	c := "cd " + path
	if systype == "windows" {
		cmd = exec.Command("cmd.exe", "/c", c)
	} else {
		cmd = exec.Command("bash", "-c", c)
	}

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func downloadFile(url, filepath string, wg *sync.WaitGroup, progressChan chan int) {
	defer wg.Done()

	// 发起HTTP GET请求获取文件内容
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("无法下载文件 %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// 创建本地文件
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("无法创建文件 %s: %v\n", filepath, err)
		return
	}
	defer file.Close()

	// 获取文件大小
	fileSize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	// 创建多个写入器，以便同时将内容写入文件和进度条
	writer := io.MultiWriter(file, NewProgressBar(fileSize, progressChan))

	// 将响应内容复制到文件和进度条
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		fmt.Printf("复制文件内容时出错 %s: %v\n", url, err)
		return
	}

	fmt.Printf("下载完成 %s\n", url)
}
