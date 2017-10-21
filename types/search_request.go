package types

import "time"

//搜索请求
type SearchRequest struct {
	Latitude float64
	Longitude float64
	CountOnly bool
	Offset int64
	Limit int64
	SearchOption *SearchOptions //可留空，使用引擎默认参数
}

//可不设置的搜索参数
type SearchOptions struct{
	Refresh bool
	OrderAsc bool
	Timeout time.Duration
	Accuracy int //计算进度
	Circles int //圈数，默认1，不扩散
	Excepts map[int64]bool //排除指定ID
	Filter func(doc IndexedDocument)bool
}

const (
	_ = iota
	ACCURATE
	APPROX
)
