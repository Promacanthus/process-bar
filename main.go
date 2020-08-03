package main

import (
	"fmt"
	"time"
)

type Bar struct {
	percent int64  // 百分比
	cur     int64  // 当前进度位置
	total   int64  // 总进度
	rate    string // 进度条
	graph   string // 显示符号
}

func main() {
	fmt.Println(`Command line progress bar example 
Copyright (C) 2020  CLPB 1.0.0    July 31 2020
	`)
	var bar Bar
	bar.NewOption(0, 11)
	for i := 0; i <= 11; i++ {
		time.Sleep(1000 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}

// NewOption create a new process bar with start position and all step
func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i++ {
		bar.rate += bar.graph // 初始化进度条位置
	}
}

func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

// NewOptionWithGraph create a new process bar with customized progress display pattern
func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

// Play start to run the process bar
func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-10s]  %3d%%  %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

// Finish indicates that the progress bar finished and Output carriage return
func (bar *Bar) Finish() {
	fmt.Println()
}
