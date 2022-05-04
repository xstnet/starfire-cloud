package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 处理列表分页
func ProcessPageByList(page, pageSize, defaultPageSize int) (limit int, offset int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = defaultPageSize
	}

	// 限制最大值
	if pageSize > 10000 {
		pageSize = 10000
	}

	offset = (page - 1) * pageSize
	limit = pageSize

	return
}

func DumpVal(val interface{}) {
	fmt.Println("[dump-val], type:", reflect.TypeOf(val))
	fmt.Println("[dump-val], value:", val)
}

func Struct2Map(target interface{}) (data map[string]interface{}) {
	jsonData, _ := json.Marshal(target)
	json.Unmarshal(jsonData, &data)
	return
}
