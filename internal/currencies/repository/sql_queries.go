package repository

const (
	// getTotal query
	getTotal = `SELECT COUNT(currency_id) FROM currencies WHERE deleted_at IS NULL`

	// getCurrencies query
	getCurrencies = `SELECT currency_id, name, code, symbol, created_at, updated_at
					FROM currencies
					WHERE deleted_at IS NULL
					ORDER BY created_at, updated_at
					OFFSET $2 LIMIT $3`
)
