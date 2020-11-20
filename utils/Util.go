package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

//宣告常數
const (
	ROOT               string = `C:\MyFile\OnlineClass\it360\Golang\MaskCountWeb`
	ASSET_PATH         string = "asset"
	PHARMACY_CSV_FILE  string = "pharmacy.csv"
	API_KEY            string = ""
	GEOCODING_API_URI  string = "https://maps.googleapis.com/maps/api/geocode/json?key=%s&address=%s"
	MASK_COUNT_CSV_URL string = "http://data.nhi.gov.tw/Datasets/Download.ashx?rid=A21030000I-D50001-001&l=https://data.nhi.gov.tw/resource/mask/maskdata.csv"
	//position json檔
	POSITION_JSON_FILE string = "position.json"
	//log路徑
	LOG_PATH string = "log"
	//新增not_found_pharmacy json的路徑
	NOT_FOUND_PHARMACY_FILE_NAME string = "not_found_pharmacy"

	//not_found_pharmacy經緯度的json檔名稱
	NOT_FOUND_PHARMACY_LATLNG_FILE string = "not_found_pharmacy_latlng.json"
)

func GetNotFoundPharmacyJsonPath() string {
	return filepath.Join(getAssetPath(), NOT_FOUND_PHARMACY_FILE_NAME+".json")
}

func GetNotFoundPharmacyLatLngJsonPath() string {
	return filepath.Join(getAssetPath(), NOT_FOUND_PHARMACY_LATLNG_FILE)
}

//取得不存在於position的藥局存檔路徑
func GetNotFoundPharmacySaveJsonPath() string {
	tmpLogName := fmt.Sprintf("%s%s.json",
		NOT_FOUND_PHARMACY_FILE_NAME,
		time.Now().Format("-2006-01-02"))
	return filepath.Join(getAssetPath(), tmpLogName)
}

//測試NotFoundPharmacyJson檔案是否存在
func NotFoundPharmacyJsonExist() bool {
	_, err := os.Stat(GetNotFoundPharmacySaveJsonPath())
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//取得GeocodingApi網址
func GetGeocodingApiUrl(address string) string {
	return fmt.Sprintf(GEOCODING_API_URI, API_KEY, address)
}

//取得AssetPath
func getAssetPath() string {
	return filepath.Join(ROOT, ASSET_PATH)
}

//取得PharmacyCsvPath
func GetPharmacyCsvPath() string {
	return filepath.Join(getAssetPath(), PHARMACY_CSV_FILE)
}

//取得Position.jspn存取路徑
func GetPositionJsonPath() string {
	return filepath.Join(getAssetPath(), POSITION_JSON_FILE)
}

//取得Log的資料夾路徑
func getLogDirPath() string {
	return filepath.Join(ROOT, LOG_PATH)
}

//取得Log的檔案名稱格式為年月日
func getLogFilePath(logName string) string {
	tmpLogName := fmt.Sprintf("%s%s.log",
		logName,
		time.Now().Format("-2006-01-02")) //表示日期格式年月日
	return filepath.Join(getLogDirPath(), tmpLogName)
}
func GetLogger(prefix string, logFileName string) (*log.Logger, *os.File) {
	f, _ := os.OpenFile(getLogFilePath(logFileName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//0666 表示這檔案可讀可寫
	//log.LstdFlags 表示 log的格式為 2009/01/23 01:23:23 message
	logger := log.New(f, prefix, log.LstdFlags)
	return logger, f
}
