package utils

import "time"

func FormatUnixTime(unix int64, layout string) string {
	if unix <= 0 {
		return ""
	}
	return time.Unix(unix, 0).Format(layout)
}
