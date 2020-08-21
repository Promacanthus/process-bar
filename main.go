package main

import (
	"flag"
	"time"
)

// 进度条
// func main() {
// 	fmt.Println(`Command line progress bar example
// Copyright (C) 2020  CLPB 1.0.0    July 31 2020
// 	`)
// 	var bar Bar
// 	bar.NewOption(0, 11)
// 	for i := 0; i <= 11; i++ {
// 		time.Sleep(1000 * time.Millisecond)
// 		bar.Play(int64(i))
// 	}
// 	bar.Finish()
// }

// 旋转加载
func main() {
	flag.Parse()
	s := NewSpinner("working...")
	for i := 0; i < 100; i++ {
		if isTTY() {
			s.Tick()
		}
		time.Sleep(100 * time.Millisecond)
	}
}