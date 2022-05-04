package date

import "time"

func FormatTimestamp(timestamp int64) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}
