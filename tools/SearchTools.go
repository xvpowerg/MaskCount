package tools

import (
	"bufio"
	"encoding/json"
	"os"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/utils"
)

func LoadingMaskList(positionList []*obj.Position, latlonP1 *obj.LatLng) []*obj.Mask {
	result := make([]*obj.Mask, 0, 7000)
	maskCountMap := CsvToMaskCountMap()
	for _, position := range positionList {
		maskCount, ok := maskCountMap[position.ID]
		if ok {
			//表示此藥局存在於我們的藥局資訊內
			//剩餘的map表示 不存在於我們position.json藥局資料內的藥局
			delete(maskCountMap, position.ID)
		} else {
			//表示口罩售完
			maskCount = &obj.MaskCount{
				Adult: 0,
				Ahild: 0,
			}
		}

		distance := 0.0
		if latlonP1 != nil {
			latlonP2 := &obj.LatLng{
				Lat: position.Lat,
				Lng: position.Lng,
			}
			distance = Distance(latlonP1, latlonP2)
		}

		mask := &obj.Mask{
			Position:  position,
			MaskCount: maskCount,
			Distance:  distance,
		}
		result = append(result, mask)

	}

	if len(maskCountMap) != 0 && !utils.NotFoundPharmacyJsonExist() {
		saveNotFoundMaskCountMap(maskCountMap)
	}
	return result
}

func saveNotFoundMaskCountMap(maskCountMap map[string]*obj.MaskCount) {
	f, _ := os.OpenFile(utils.GetNotFoundPharmacySaveJsonPath(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	w := bufio.NewWriter(f)
	defer f.Close()
	b, _ := json.Marshal(maskCountMap)
	w.Write(b)
	w.Flush()
}
