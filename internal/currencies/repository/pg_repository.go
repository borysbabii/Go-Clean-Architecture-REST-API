package repository

import (
	"context"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/models"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// Currency Repository
type currenciesRepo struct {
	db *sqlx.DB
}

// NewCurrenciesRepository Currency repository constructor
func NewCurrenciesRepository(db *sqlx.DB) currencies.Repository {
	return &currenciesRepo{db: db}
}

// GetCurrencies returns list of currencies
func (r *currenciesRepo) GetCurrencies(ctx context.Context, query *utils.PaginationQuery) (*models.CurrenciesList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "currenciesRepo.GetCurrencies")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotal); err != nil {
		return nil, errors.Wrap(err, "currenciesRepo.GetCurrencies.GetContext")
	}

	if totalCount == 0 {
		return &models.CurrenciesList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Currencies: make([]*models.Currency, 0),
		}, nil
	}

	var currencyList = make([]*models.Currency, 0, query.GetSize())
	rows, err := r.db.QueryxContext(ctx, getCurrencies, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "currenciesRepo.GetCurrencies.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.Currency{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "currenciesRepo.GetCurrencies.StructScan")
		}
		currencyList = append(currencyList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "currenciesRepo.GetCurrencies.rows.Err")
	}

	return &models.CurrenciesList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Currencies: currencyList,
	}, nil
}
