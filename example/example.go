package main

import "github.com/imwithye/logor"

func main() {
	logor := logor.New()
	logor.Info("log info")
	logor.Warn("log warn")
	logor.Debug("log warn")
	logor.Error("log error")
}
