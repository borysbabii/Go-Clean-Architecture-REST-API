package repository

import (
	"context"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestCurrencyRepo_GetCurrencies(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	currencyRepo := NewCurrenciesRepository(sqlxDB)

	t.Run("GetCurrencies", func(t *testing.T) {
		uid := uuid.New()

		totalCountRows := sqlmock.NewRows([]string{"count"}).AddRow(0)

		rows := sqlmock.NewRows([]string{"currency_id", "name", "code", "symbol"}).AddRow(
			uid, "American Dollar", "USD", "$")

		mock.ExpectQuery(getTotal).WillReturnRows(totalCountRows)
		mock.ExpectQuery(getCurrencies).WillReturnRows(rows)

		currencies, err := currencyRepo.GetCurrencies(context.Background(), &utils.PaginationQuery{
			Size:    10,
			Page:    1,
			OrderBy: "",
		})
		require.NoError(t, err)
		require.NotNil(t, currencies)
	})

}
