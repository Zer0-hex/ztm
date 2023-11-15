package runner

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	// retrieve home directory or fail
	homeDir = func() string {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get user home directory: %s", err)
		}
		return home
	}()

	defaultConfigPath = filepath.Join(homeDir, ".ztm/config.toml")
	defaultResPath    = filepath.Join(homeDir, ".ztm/resource")
)

func Banner() {
	banner := `
███████╗████████╗███╗   ███╗
╚══███╔╝╚══██╔══╝████╗ ████║
  ███╔╝    ██║   ██╔████╔██║
 ███╔╝     ██║   ██║╚██╔╝██║
███████╗   ██║   ██║ ╚═╝ ██║
╚══════╝   ╚═╝   ╚═╝     ╚═╝ 
    Zer0-hex      ` + version + `
`
	print(banner)
}

func Flag(options *Options) {
	Banner()
	flag.StringVar(&options.ConfigPath, "config", defaultConfigPath, "配置文件路径($HOME/.ztm/config.toml)")
	flag.StringVar(&options.ResPath, "respath", defaultResPath, "资源保存路径($HOME/.ztm/resources)")
	flag.BoolVar(&options.Update, "update", false, "检查更新")
	flag.BoolVar(&options.Install, "install", false, "安装")
	flag.BoolVar(&options.Upgrade, "upgrade", false, "更新")
	flag.BoolVar(&options.Go, "go", false, "切换到资源目录")
	flag.Parse()
}
