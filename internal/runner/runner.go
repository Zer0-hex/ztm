package runner

import (
	"github.com/Zer0-hex/ztm/pkg/path"
	"github.com/zer0-hex/ztm/pkg/path"
)

type Runner struct {
	options *Options
}

func NewRunner(options *Options) (*Runner, error) {
	return &Runner{
		options: options,
	}, nil
}

func (r *Runner) Run() error {
	// 检查配置文件路径
	path.CheckPath(r.options.ConfigPath)
	path.CheckPath(r.options.ResPath)

	return nil
}

func (r *Runner) Flag() {
	Flag(r.options)
}
