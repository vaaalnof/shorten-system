package query

const (

	// =====================================================
	// ANALYTICS EVENT
	// =====================================================

	AnalyticsEventCreate = `
		INSERT INTO analytics_events (
			url_id,
			short_code,
			referer,
			source,
			ip_address,
			user_agent,
			browser,
			os,
			device,
			country,
			clicked_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
		)
	`
)
