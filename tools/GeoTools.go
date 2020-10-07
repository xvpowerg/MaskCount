package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	geo "gopkg.in/Billups/golang-geo.v2"
	"tw.com.maskweb/obj"
	"tw.com.maskweb/utils"
)

func queryLatLng(position *obj.Position,
	positionChan chan<- *obj.Position,
	mylog *log.Logger) {
	jsonURL := utils.GetGeocodingApiUrl(position.Addr)
	//使用Get方式呼叫API
	res, err := http.Get(jsonURL)
	if err != nil {
		mylog.Println(err)
		positionChan <- nil
		return
	}
	//完成後關閉Body
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	//JSON最外層物件
	var result map[string]interface{}
	//開始反序列化Json
	json.Unmarshal(body, &result)
	//將result map取出的內容強制轉型成為一個陣列
	results := result["results"].([]interface{})

	if len(results) == 0 {
		mylog.Println("result is nil:", position.ID, position.Addr)
		positionChan <- nil
		return
	}

	//取得results 第一筆地理資訊
	geometryMap := results[0].(map[string]interface{})

	if geometryMap == nil {
		mylog.Println("geometryMap is nil:", position.ID, position.Addr)
		positionChan <- nil
		return
	}
	//取得geometry 資訊
	locationMap := geometryMap["geometry"].(map[string]interface{})
	//取得location 資訊
	location := locationMap["location"].(map[string]interface{})
	//取得lat 與 lng 並強制轉型為float64
	position.Lat = location["lat"].(float64)
	position.Lng = location["lng"].(float64)
	positionChan <- position
}
func Distance(latLng1 *obj.LatLng, latLng2 *obj.LatLng) float64 {
	p1 := geo.NewPoint(latLng1.Lat, latLng1.Lng)
	p2 := geo.NewPoint(latLng2.Lat, latLng2.Lng)
	//單位公里
	return p1.GreatCircleDistance(p2)
}
