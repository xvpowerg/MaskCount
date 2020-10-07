package tools

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strconv"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/utils"
)

func CsvToMaskCountMap() map[string]*obj.MaskCount {
	csvURL := utils.MASK_COUNT_CSV_URL
	//下載
	res, err := http.Get(csvURL)
	if err != nil || res == nil {
		log.Println(err)
		return nil
	}
	maskCountMap := make(map[string]*obj.MaskCount)
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

		//Atoi 字串轉整數

		//成人口罩數量
		adult, e1 := strconv.Atoi(value[4])
		//兒童口罩數量
		ahild, e2 := strconv.Atoi(value[5])
		updateTime := value[6]
		//如果字串轉整數錯誤就將口罩數量設為0
		if e1 != nil {
			adult = 0
		}
		if e2 != nil {
			ahild = 0
		}
		//建立MaskCount物件
		maskCount := &obj.MaskCount{ID: value[0],
			Name:       value[1],
			Addr:       value[2],
			Phone:      value[3],
			Adult:      adult,
			Ahild:      ahild,
			UpdateTime: updateTime}

		maskCountMap[maskCount.ID] = maskCount
	}
	return maskCountMap
}
