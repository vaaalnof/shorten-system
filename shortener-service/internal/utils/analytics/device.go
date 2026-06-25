package analytics

import "strings"

func DetectDevice(
	userAgent string,
) string {

	ua := strings.ToLower(
		userAgent,
	)

	switch {

	case strings.Contains(
		ua,
		"mobile",
	):
		return "mobile"

	case strings.Contains(
		ua,
		"android",
	):
		return "mobile"

	case strings.Contains(
		ua,
		"iphone",
	):
		return "mobile"

	case strings.Contains(
		ua,
		"ipad",
	):
		return "tablet"

	default:
		return "desktop"
	}
}
