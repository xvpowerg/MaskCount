package main

import "tw.com.maskweb/utils"

func main() {
	logger, f := utils.GetLogger("webError", "web")
	defer f.Close()
	logger.Println("TestMsg")
}
