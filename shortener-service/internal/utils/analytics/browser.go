package analytics

import "strings"

func DetectBrowser(
	userAgent string,
) string {

	ua := strings.ToLower(
		userAgent,
	)

	switch {

	case strings.Contains(
		ua,
		"instagram",
	):
		return "Instagram"

	case strings.Contains(
		ua,
		"facebook",
	):
		return "Facebook"

	case strings.Contains(
		ua,
		"edg",
	):
		return "Edge"

	case strings.Contains(
		ua,
		"opr",
	):
		return "Opera"

	case strings.Contains(
		ua,
		"chrome",
	):
		return "Chrome"

	case strings.Contains(
		ua,
		"firefox",
	):
		return "Firefox"

	case strings.Contains(
		ua,
		"safari",
	):
		return "Safari"

	default:
		return "Unknown"
	}
}
