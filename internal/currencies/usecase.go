//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package currencies

import (
	"context"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/models"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// UseCase Currencies use case
type UseCase interface {
	GetCurrencies(ctx context.Context, pq *utils.PaginationQuery) (*models.CurrenciesList, error)
}
