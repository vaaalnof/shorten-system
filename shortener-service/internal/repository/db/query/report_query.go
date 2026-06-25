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
)
