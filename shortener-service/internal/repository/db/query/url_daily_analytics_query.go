package query

const (

	// =====================================================
	// URL DAILY ANALYTICS
	// TOTAL CLICKS
	// =====================================================

	URLDailyAnalyticsAddClick = `
		INSERT INTO url_daily_analytics (
			url_id,
			analytics_date,
			total_clicks,
			unique_visitors,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			1,
			0,
			$3,
			$3
		)
		ON CONFLICT (
			url_id,
			analytics_date
		)
		DO UPDATE SET
			total_clicks =
				url_daily_analytics.total_clicks + 1,

			updated_at =
				EXCLUDED.updated_at
	`

	// =====================================================
	// URL DAILY ANALYTICS
	// UNIQUE VISITORS
	// =====================================================

	URLDailyAnalyticsAddUniqueVisitor = `
		INSERT INTO url_daily_analytics (
			url_id,
			analytics_date,
			total_clicks,
			unique_visitors,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			0,
			1,
			$3,
			$3
		)
		ON CONFLICT (
			url_id,
			analytics_date
		)
		DO UPDATE SET
			unique_visitors =
				url_daily_analytics.unique_visitors + 1,

			updated_at =
				EXCLUDED.updated_at
	`
)
