package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"tw.com.maskweb/utils"
)

func QueryLatLng() {
	addr := "臺北市北投區石牌路二段111號"
	jsonURL := utils.GetGeocodingApiUrl(addr)
	//使用Get方式呼叫API
	res, err := http.Get(jsonURL)
	if err != nil {
		log.Println(err)
		return
	}
	//完成後關閉Body
	defer res.Body.Close()
	if res == nil || res.Body == nil {
		log.Println(err)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	//JSON最外層物件
	var result map[string]interface{}
	//開始反序列化Json
	json.Unmarshal(body, &result)
	//將result map取出的內容強制轉型成為一個陣列
	results := result["results"].([]interface{})
	//取得results 第一筆地理資訊
	geometryMap := results[0].(map[string]interface{})
	//取得geometry 資訊
	locationMap := geometryMap["geometry"].(map[string]interface{})
	//取得location 資訊
	location := locationMap["location"].(map[string]interface{})
	//取得lat 與 lng 並強制轉型為float64
	fmt.Println(location["lat"].(float64), location["lng"].(float64))
}
