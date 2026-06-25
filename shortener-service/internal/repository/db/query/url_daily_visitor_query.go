package query

const (

	// =====================================================
	// URL DAILY VISITOR
	// =====================================================

	URLDailyVisitorCreate = `
		INSERT INTO url_daily_visitors (
			url_id,
			analytics_date,
			visitor_hash,
			created_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
		ON CONFLICT DO NOTHING
	`
)
