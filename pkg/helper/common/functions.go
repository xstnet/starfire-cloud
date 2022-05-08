package common

import (
	"encoding/json"
	"fmt"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"math"
	"reflect"
)

// ProcessPage 处理列表分页
func ProcessPage(page, pageSize, defaultPageSize int) (limit int, offset int) {
	page = int(math.Max(float64(page), 1))
	pageSize = int(math.Max(float64(pageSize), float64(defaultPageSize)))

	// 限制最大值
	if pageSize > 10000 {
		pageSize = 10000
	}

	offset = (page - 1) * pageSize
	limit = pageSize

	return
}

func DumpVal(val any) {
	fmt.Println("[dump-val], type:", reflect.TypeOf(val))
	fmt.Println("[dump-val], value:", val)
}

func Struct2Map(target any) (data d.StringMap) {
	jsonData, _ := json.Marshal(target)
	json.Unmarshal(jsonData, &data)
	return
}
