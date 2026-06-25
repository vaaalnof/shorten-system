package analytics

import "strings"

func DetectOS(
	userAgent string,
) string {

	ua := strings.ToLower(
		userAgent,
	)

	switch {

	case strings.Contains(
		ua,
		"windows",
	):
		return "windows"

	case strings.Contains(
		ua,
		"android",
	):
		return "android"

	case strings.Contains(
		ua,
		"iphone",
	):
		return "ios"

	case strings.Contains(
		ua,
		"ipad",
	):
		return "ios"

	case strings.Contains(
		ua,
		"mac os",
	):
		return "macos"

	case strings.Contains(
		ua,
		"linux",
	):
		return "linux"

	default:
		return "unknown"
	}
}
