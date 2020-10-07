package main

import (
	"fmt"

	"tw.com.maskweb/tools"
)

func main() {
	maskmap := tools.CsvToMaskCountMap()
	fmt.Println(maskmap)
}
