package date

import (
	"golang.org/x/exp/constraints"
	"time"
)

func Format[T constraints.Integer](timestamp T) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}