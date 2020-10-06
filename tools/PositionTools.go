package tools

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

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
	const (
		WAIT_COUNT    uint32 = 50
		WAIT_DURATION uint8  = 2
	)
	var count uint32 = 1
	for {

		pharmacy, err2 := r.Read()
		if err2 == io.EOF {
			break
		}

		if count%WAIT_COUNT == 0 {
			fmt.Println("Stop Wait:", count)
			//暫停2秒
			time.Sleep(time.Duration(WAIT_DURATION) * time.Second)
		}
		count++
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

func GetPositionList() []*obj.Position {
	f, _ := ioutil.ReadFile(utils.GetPositionJsonPath())
	//長度為0 容量為7000 想像7000cc的杯子 沒有裝東西
	//因為Split 容量不夠會自動擴充 需要計算會浪費效能
	//又因csv筆數為6871筆 所以給Split預設容量為7000
	positionList := make([]*obj.Position, 0, 7000)
	json.Unmarshal(f, &positionList)
	return positionList
}
