package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	RootPath string
	WinPath  string
	Res      []Resource
}

type Resource struct {
	Tag    string
	Action string
	Sub    []SubResource
}

type SubResource struct {
	Name       string
	URL        string
	TagVersion string
	Files      []string
}

func loadconfig() {
	// 读取配置文件
	var config Config
	var toBe map[string]string

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	fmt.Println("Unix/Linux Path:", config.RootPath)
	fmt.Println("Windows Path:", config.WinPath)
	for _, res := range config.Res {
		fmt.Println("+--------------------------------------+")
		fmt.Println("Res Tag:", res.Tag)
		fmt.Println("Res Action:", res.Action)
		fmt.Println("----------------------------------------")
		for _, subres := range res.Sub {
			fmt.Println("\tName:", subres.Name)
			fmt.Println("\tUrl:", subres.URL)

			if res.Action == "download" {
				v, err := getVersion(subres.URL)
				if err != nil {
					fmt.Println("[-] Request error", err, subres.URL)
				}
				subres.URL = v
				temp := strings.Split(v, "/")
				// 获取切割后的切片的最后一部分
				tagVersion := temp[len(temp)-1]
				subres.TagVersion = tagVersion

			}
			savePath := config.RootPath + "/" + res.Tag + "/" + subres.Name
			fmt.Println(savePath)
			fmt.Println("\tVersion:", subres.TagVersion)
			fmt.Printf("\tFiles: ")
			for i, file := range subres.Files {
				fmt.Printf("%s, ", file)
				if i%4 == 3 {
					fmt.Printf("\n\t\t")
				}
			}
			fmt.Println()
			fmt.Println("----------------------------------------")
		}
	}

	filepath := "./downloads/" // 替换为要保存的文件夹路径

	// 创建等待组和进度通道
	var wg sync.WaitGroup
	progressChan := make(chan int)

	// 启动进度条更新协程
	go func() {
		progressBars := make(map[int]*ProgressBar)
		for progressBarID := range progressChan {
			if progressBar, ok := progressBars[progressBarID]; ok {
				progressBar.printProgress()
			} else {
				progressBars[progressBarID] = NewProgressBar(100, nil)
			}
		}
	}()

	// 启动多个协程下载文件
	for _, url := range toBe {
		wg.Add(1)
		go downloadFile(url, filepath, &wg, progressChan)
	}

	// 等待所有协程完成
	wg.Wait()

	// 关闭进度通道
	close(progressChan)

	fmt.Println("所有文件下载完成！")

}

func getVersion(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("无法下载文件 %s: %v\n", url, err)
		return "", err
	}
	return resp.Request.URL.String(), nil
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

// 进度条结构体
type ProgressBar struct {
	Total         int        // 总大小
	Current       int        // 当前进度
	ProgressChan  chan<- int // 进度通道
	ProgressBarID int        // 进度条ID
}

// Write implements io.Writer.
func (*ProgressBar) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}

// 创建新的进度条
func NewProgressBar(total int, progressChan chan<- int) *ProgressBar {
	return &ProgressBar{
		Total:         total,
		Current:       0,
		ProgressChan:  progressChan,
		ProgressBarID: len(progressChan),
	}
}

// 更新进度条进度
func (p *ProgressBar) UpdateProgress(progress int) {
	p.Current = progress
	p.ProgressChan <- p.ProgressBarID
}

// 打印进度条
func (p *ProgressBar) printProgress() {
	progress := p.Current * 100 / p.Total
	fmt.Printf("进度条 %d: %3d%%\n", p.ProgressBarID, progress)
}
