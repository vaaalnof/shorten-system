package analytics

import "strings"

func DetectSource(
	referer string,
) string {

	referer = strings.ToLower(
		strings.TrimSpace(
			referer,
		),
	)

	switch {

	case referer == "":
		return "direct"

	case strings.Contains(
		referer,
		"instagram",
	):
		return "instagram"

	case strings.Contains(
		referer,
		"facebook",
	):
		return "facebook"

	case strings.Contains(
		referer,
		"t.co",
	):
		return "twitter"

	case strings.Contains(
		referer,
		"twitter",
	):
		return "twitter"

	case strings.Contains(
		referer,
		"linkedin",
	):
		return "linkedin"

	case strings.Contains(
		referer,
		"youtube",
	):
		return "youtube"

	case strings.Contains(
		referer,
		"tiktok",
	):
		return "tiktok"

	case strings.Contains(
		referer,
		"google",
	):
		return "google"

	case strings.Contains(
		referer,
		"bing",
	):
		return "bing"

	case strings.Contains(
		referer,
		"yahoo",
	):
		return "yahoo"

	default:
		return "other"
	}
}
