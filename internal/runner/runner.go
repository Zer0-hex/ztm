package runner

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Zer0-hex/ztm/pkg"
)

func NewRunner(options *Options) (*Runner, error) {
	return &Runner{
		options: options,
	}, nil
}

func (r *Runner) Run() error {
	// 检查配置文件路径

	if err := pkg.CheckFilePath(r.options.ConfigPath); err != nil {
		log.Fatal(err)
	}
	if err := pkg.CheckDirPath(r.options.ResPath); err != nil {
		log.Fatal(err)
	}
	if err := r.Loadconfig(); err != nil {
		log.Fatal(err)
	}
	switch {
	case r.options.Update:
		if err := r.Update(); err != nil {
			return err
		}
	case r.options.Install:
		if err := r.Install(); err != nil {
			return err
		}
	case r.options.Upgrade:
		if err := r.Upgrade(); err != nil {
			return err
		}
	case r.options.Go:
		if err := r.Go(); err != nil {
			log.Fatalf("[-] Change dir fail: %v", err)
		}
	}

	return nil
}

func (r *Runner) Flag() {
	Flag(r.options)
}

func (r *Runner) Loadconfig() error {
	// 读取配置文件
	var toollist ToolList
	if _, err := toml.DecodeFile(r.options.ConfigPath, &toollist); err != nil {
		return err
	}
	return nil
}

func (r *Runner) Update() error {
	if err := r.getVersion(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) Install() error {
	if err := r.getVersion(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) Upgrade() error {
	if err := r.getVersion(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) Go() error {
	if err := pkg.ChangeDir(r.options.ResPath); err != nil {
		return err
	}
	log.Println("[+] Change dir to", r.options.ResPath)
	return nil
}

func (r *Runner) getVersion() error {
	for _, tool := range r.toolList.tools {
		if version, err := pkg.GetVersion(tool.Url); err != nil {
			return err
		} else {
			tool.Version = version
		}
	}
	return nil
}

// func () {
// 	for _, tool := range toollist.tools {
// 		fmt.Println("Res Name:", tool.Name)
// 		fmt.Println("Res Tag:", tool.Tag)
// 		fmt.Println("Res Url:", tool.Url)
// 		fmt.Println("Res Files:", tool.Files)
// 		fmt.Println("+--------------------------------------+")

// 		for _, subres := range res.Sub {
// 			fmt.Println("\tName:", subres.Name)
// 			fmt.Println("\tUrl:", subres.URL)

// 			if res.Action == "download" {
// 				v, err := getVersion(subres.URL)
// 				if err != nil {
// 					fmt.Println("[-] Request error", err, subres.URL)
// 				}
// 				subres.URL = v
// 				temp := strings.Split(v, "/")
// 				// 获取切割后的切片的最后一部分
// 				tagVersion := temp[len(temp)-1]
// 				subres.TagVersion = tagVersion

// 			}
// 			savePath := config.RootPath + "/" + res.Tag + "/" + subres.Name
// 			fmt.Println(savePath)
// 			fmt.Println("\tVersion:", subres.TagVersion)
// 			fmt.Printf("\tFiles: ")
// 			for i, file := range subres.Files {
// 				fmt.Printf("%s, ", file)
// 				if i%4 == 3 {
// 					fmt.Printf("\n\t\t")
// 				}
// 			}
// 			fmt.Println()
// 			fmt.Println("----------------------------------------")
// 		}
// 	}

// 	filepath := "./downloads/" // 替换为要保存的文件夹路径

// 	// 创建等待组和进度通道
// 	var wg sync.WaitGroup
// 	progressChan := make(chan int)

// 	// 启动进度条更新协程
// 	go func() {
// 		progressBars := make(map[int]*ProgressBar)
// 		for progressBarID := range progressChan {
// 			if progressBar, ok := progressBars[progressBarID]; ok {
// 				progressBar.printProgress()
// 			} else {
// 				progressBars[progressBarID] = NewProgressBar(100, nil)
// 			}
// 		}
// 	}()

// 	// 启动多个协程下载文件
// 	for _, url := range toBe {
// 		wg.Add(1)
// 		go downloadFile(url, filepath, &wg, progressChan)
// 	}

// 	// 等待所有协程完成
// 	wg.Wait()

// 	// 关闭进度通道
// 	close(progressChan)

// 	fmt.Println("所有文件下载完成！")

// }
