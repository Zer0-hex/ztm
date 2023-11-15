package pkg

type Options struct {
	ConfigPath string
	ResPath    string

	Update  bool
	Install bool
	Upgrade bool
	Go      bool
}

type ToolList struct {
	tools []Tool
}

type Tool struct { // webshell, seclists
	Name    string
	Tag     string
	Version string
	Url     string
	Files   []string
}

// 进度条结构体
type ProgressBar struct {
	Total         int        // 总大小
	Current       int        // 当前进度
	ProgressChan  chan<- int // 进度通道
	ProgressBarID int        // 进度条ID
}
