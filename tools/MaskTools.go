package tools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CsvToMaskCountMap() {
	csvURL := "http://data.nhi.gov.tw/Datasets/Download.ashx?rid=A21030000I-D50001-001&l=https://data.nhi.gov.tw/resource/mask/maskdata.csv"
	//下載
	res, err := http.Get(csvURL)
	if err != nil || res == nil {
		log.Println(err)
		return
	}
	//注意要關閉 不然會占用TCPIP數量
	defer res.Body.Close()
	//讀取
	r := csv.NewReader(bufio.NewReader(res.Body))
	//第一筆資料是Title 所以跳過
	_, _ = r.Read()
	for {
		value, err := r.Read()
		//無資料時err == io.EOF後離開迴圈
		if err == io.EOF {
			break
		}
		fmt.Println(value[0], value[1], value[2], value[3], value[4], value[5], value[6])

	}
}
