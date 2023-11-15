package pkg

import "fmt"

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
