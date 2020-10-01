package tools

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/utils"
)

func QueryPharmacyLatLngSaveToJSON() {

	//建立chan
	positionChan := make(chan *obj.Position)
	positionArryChan := make(chan []*obj.Position) //用來接收collectPosition 的陣列
	//收集Position
	go collectPosition(positionChan, positionArryChan)
	//產生Position
	go queryLatlngByPharmacy(positionChan)

	//輸出成JSON
	positionArry := <-positionArryChan
	fmt.Println("write to Marshal File...")
	data, _ := json.Marshal(positionArry)
	ioutil.WriteFile(utils.GetPositionJsonPath(), data, 0666)
}

//接收Position
func collectPosition(outPositionChan <-chan *obj.Position,
	positionChanArray chan<- []*obj.Position) {
	var positionArry []*obj.Position
	for positionObj := range outPositionChan {
		if positionObj != nil {
			positionArry = append(positionArry, positionObj)
		}
	}
	positionChanArray <- positionArry
}

//傳送Position
func queryLatlngByPharmacy(inPositionChan chan<- *obj.Position) {
	//加入logger
	logger, f := utils.GetLogger("QueryLatlng:", "queryLatlng")
	defer f.Close()

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

		position := &obj.Position{
			ID:    pharmacy[0],
			Name:  pharmacy[1],
			Phone: pharmacy[2],
			Addr:  pharmacy[3],
		}
		go queryLatLng(position, inPositionChan, logger)
	}
	//因該要所有queryLatLng都做完才close
	close(inPositionChan)

}
