package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/tools"
)

var tpl *template.Template
var positionList []*obj.Position

func init() {
	tpl = template.Must(template.ParseFiles("../template/main.html"))
	positionList = tools.GetPositionList()
}

func home(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Println("err:", err)
	}
}

//地址查詢口罩販賣地點
func addr(res http.ResponseWriter, req *http.Request) {
	addr := req.FormValue("addr")
	//產生maskCount
	maskList := tools.LoadingMaskList(positionList, nil)
	//依地址過濾
	result := tools.SearchByAddress(addr, maskList)

	err := tpl.Execute(res, result)
	if err != nil {
		log.Println("err:", err)
	}
}

//經緯度查詢口罩販賣地點
func searchLatlng(res http.ResponseWriter, req *http.Request) {
	latF := req.FormValue("lat")
	lngF := req.FormValue("lng")

	//將字串轉換成浮點數
	lat, _ := strconv.ParseFloat(latF, 64)
	lng, _ := strconv.ParseFloat(lngF, 64)
	latLng := &obj.LatLng{
		Lat: lat,
		Lng: lng,
	}
	maskList := tools.LoadingMaskList(positionList, latLng)
	result := tools.SortByDistance(10, maskList)
	err := tpl.Execute(res, result)

	if err != nil {
		log.Println("err:", err)
	}

}

func main() {
	//註冊網址為/latlng
	http.Handle("/latlng", http.HandlerFunc(searchLatlng))
	//註冊網址為/addr
	http.Handle("/addr", http.HandlerFunc(addr))
	//註冊 根目錄網址
	http.Handle("/", http.HandlerFunc(home))

	http.ListenAndServe(":8080", nil)
}
