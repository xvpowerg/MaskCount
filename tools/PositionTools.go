package tools

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"tw.com.maskweb/utils"
)

func QueryPharmacyLatLngSaveToJSON() {
	//收集Position
	collectPosition()
	//產生Position
	queryLatlngByPharmacy()
	//輸出成JSON
}

func collectPosition() {

}

func queryLatlngByPharmacy() {
	path := utils.GetPharmacyCsvPath()
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return
	}
	r := csv.NewReader(f)
	_, _ = r.Read()
	for {
		pharmacy, err2 := r.Read()
		if err2 == io.EOF {
			break
		}
		fmt.Println(pharmacy[0] + ":" + pharmacy[1] + ":" + pharmacy[2] + ":" + pharmacy[3])
	}

}
