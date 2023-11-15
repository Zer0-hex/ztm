package runner

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const version = "v0.0.1"

var (
	// retrieve home directory or fail
	homeDir = func() string {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Failed to get user home directory: %s", err)
		}
		return home
	}()

	defaultConfigLocation = filepath.Join(homeDir, ".ztm/config.toml")
	defaultPath           = filepath.Join(homeDir, ".ztm/resource")
)

// Options contains the configuration options for tuning the enumeration process.
type Options struct {
	ConfigFile string
	ResPath    string

	Update  bool
	Install bool
	Upgrade bool
	Version string
}

func Banner() {
	banner := `
███████╗████████╗███╗   ███╗
╚══███╔╝╚══██╔══╝████╗ ████║
  ███╔╝    ██║   ██╔████╔██║
 ███╔╝     ██║   ██║╚██╔╝██║
███████╗   ██║   ██║ ╚═╝ ██║
╚══════╝   ╚═╝   ╚═╝     ╚═╝ 
                         Author: Zer0-hex 
						 Version: ` + version + `
`
	print(banner)
}

func Flag(options Options) {
	Banner()
	flag.StringVar(&options.ConfigFile, "config", "", "配置文件路径")
	flag.StringVar(&options.ResPath, "ResPath", "", "保存路径")
	flag.BoolVar(&options.Update, "update", false, "检查更新")
	flag.BoolVar(&options.Install, "install", false, "安装")
	flag.BoolVar(&options.Upgrade, "upgrade", false, "更新")
	flag.StringVar(&options.Version, "version", "", "版本")
	flag.Parse()
}
