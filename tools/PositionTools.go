package tools

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/utils"
)

func QueryPharmacyLatLngSaveToJSON() {

	//建立chan
	positionChan := make(chan *obj.Position)
	positionArryChan := make(chan []*obj.Position) //用來接收collectPosition 的陣列
	//加入一個WaitGroup
	var wg sync.WaitGroup
	//收集Position
	go collectPosition(positionChan, positionArryChan, &wg)
	//產生Position
	go queryLatlngByPharmacy(positionChan, &wg)

	//輸出成JSON
	positionArry := <-positionArryChan
	fmt.Println("write to Marshal File...")
	data, _ := json.Marshal(positionArry)
	ioutil.WriteFile(utils.GetPositionJsonPath(), data, 0666)
}

//接收Position
func collectPosition(outPositionChan <-chan *obj.Position,
	positionChanArray chan<- []*obj.Position, wg *sync.WaitGroup) {
	var positionArry []*obj.Position
	for positionObj := range outPositionChan {
		if positionObj != nil {
			positionArry = append(positionArry, positionObj)
		}
		wg.Done()
	}
	positionChanArray <- positionArry
}

//傳送Position
func queryLatlngByPharmacy(inPositionChan chan<- *obj.Position, wg *sync.WaitGroup) {
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
		wg.Add(1)
		go queryLatLng(position, inPositionChan, logger)
	}

	wg.Wait()
	//因該要所有queryLatLng都做完才close
	close(inPositionChan)

}
