package query

const (

	// =====================================================
	// REPORT SUMMARY
	// =====================================================

	GetReportSummary = `
		SELECT
			COALESCE(SUM(total_clicks), 0),
			COALESCE(SUM(unique_visitors), 0),
			COALESCE(SUM(
				CASE
					WHEN analytics_date = $2 THEN total_clicks
					ELSE 0
				END
			), 0),
			COALESCE(SUM(
				CASE
					WHEN analytics_date = $2 THEN unique_visitors
					ELSE 0
				END
			), 0)
		FROM url_daily_analytics
		WHERE url_id = $1
	`

	// =====================================================
	// REPORT CHART
	// =====================================================

	GetReportChart = `
		SELECT
			analytics_date,
			total_clicks,
			unique_visitors
		FROM url_daily_analytics
		WHERE url_id = $1
		ORDER BY analytics_date ASC
	`

	// =====================================================
	// REPORT REFERRERS
	// =====================================================

	GetReportReferrers = `
		SELECT
			COALESCE(
				NULLIF(
					TRIM(source),
					''
				),
				'direct'
			) AS referrer,
			COUNT(*) AS clicks
		FROM analytics_events
		WHERE url_id = $1
		GROUP BY referrer
		ORDER BY clicks DESC
	`

	// =====================================================
	// REPORT COUNTRIES
	// =====================================================

	GetReportCountries = `
		SELECT
			country,
			COUNT(*) AS clicks
		FROM (
			SELECT
				CASE
					WHEN country IS NULL THEN 'other'
					WHEN TRIM(country) = '' THEN 'other'
					WHEN LOWER(TRIM(country)) = 'unknown' THEN 'other'
					ELSE TRIM(country)
				END AS country
			FROM analytics_events
			WHERE url_id = $1
		) t
		GROUP BY country
		ORDER BY clicks DESC
	`

	// =====================================================
	// REPORT DEVICES
	// =====================================================

	GetReportDevices = `
		SELECT
			CASE
				WHEN device IS NULL THEN 'other'
				WHEN TRIM(device) = '' THEN 'other'
				WHEN LOWER(device) = 'unknown' THEN 'other'
				ELSE LOWER(device)
			END AS device,
			COUNT(*) AS clicks
		FROM analytics_events
		WHERE url_id = $1
		GROUP BY
			CASE
				WHEN device IS NULL THEN 'other'
				WHEN TRIM(device) = '' THEN 'other'
				WHEN LOWER(device) = 'unknown' THEN 'other'
				ELSE LOWER(device)
			END
		ORDER BY clicks DESC
	`

	// =====================================================
	// REPORT BROWSERS
	// =====================================================

	GetReportBrowsers = `
		SELECT
			CASE
				WHEN browser IS NULL THEN 'other'
				WHEN TRIM(browser) = '' THEN 'other'
				WHEN LOWER(browser) = 'unknown' THEN 'other'
				ELSE browser
			END AS browser,
			COUNT(*) AS clicks
		FROM analytics_events
		WHERE url_id = $1
		GROUP BY
			CASE
				WHEN browser IS NULL THEN 'other'
				WHEN TRIM(browser) = '' THEN 'other'
				WHEN LOWER(browser) = 'unknown' THEN 'other'
				ELSE browser
			END
		ORDER BY clicks DESC
	`

	// =====================================================
	// TOP LINKS
	// =====================================================

	GetTopLinks = `
		SELECT
			u.id,
			u.short_code,
			u.original_url,
			COALESCE(SUM(a.total_clicks), 0) AS clicks,
			COALESCE(SUM(a.unique_visitors), 0) AS unique_visitors,
			u.created_at
		FROM urls u
		LEFT JOIN url_daily_analytics a
			ON a.url_id = u.id
		WHERE u.user_id = $1
		GROUP BY
			u.id,
			u.short_code,
			u.original_url,
			u.created_at
		ORDER BY
			clicks DESC,
			u.created_at DESC
		LIMIT 10
	`
)
