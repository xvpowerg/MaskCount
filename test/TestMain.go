package main

import (
	"fmt"

	"tw.com.maskweb/tools"
)

func main() {
	pList := tools.GetPositionList()
	list := tools.LoadingMaskList(pList, nil)
	fmt.Println(list[0])
}
