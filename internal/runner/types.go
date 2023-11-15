package runner

const version = "v0.0.1"

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
	Name      string
	Tag       string
	Version   string
	Url       string
	Files     []string
	FilesLink []string
}

type Runner struct {
	options  *Options
	toolList ToolList
}
