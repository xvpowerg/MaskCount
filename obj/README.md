# Object
專案內所有物件於此檔案
```go
type Position struct {
	ID    string
	Name  string
	Phone string
	Addr  string
	Lat   float64
	Lng   float64
}
type MaskCount struct {
	ID         string
	Name       string
	Phone      string
	Addr       string
	Adult      int
	Ahild      int
	UpdateTime string
}

type LatLng struct {
	Lat float64
	Lng float64
}

//使用嵌入Embed
type Mask struct {
	*Position
	*MaskCount
	Distance float64
}

```