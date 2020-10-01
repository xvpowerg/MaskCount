package tools

import (
	"encoding/csv"
	"io"
	"os"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/utils"
)

func QueryPharmacyLatLngSaveToJSON() {

	//建立chan
	positionChan := make(chan *obj.Position)
	//收集Position
	go collectPosition(positionChan)
	//產生Position
	go queryLatlngByPharmacy(positionChan)
	//輸出成JSON
}

func collectPosition(outPositionChan <-chan *obj.Position) {
	//接收Position
}

func queryLatlngByPharmacy(inPositionChan chan<- *obj.Position) {
	//傳送Position
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
		go queryLatLng(position, inPositionChan)
	}

}
