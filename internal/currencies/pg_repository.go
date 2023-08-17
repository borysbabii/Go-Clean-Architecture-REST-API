//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package currencies

import (
	"context"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/models"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

// Repository Currencies repository
type Repository interface {
	GetCurrencies(ctx context.Context, pq *utils.PaginationQuery) (*models.CurrenciesList, error)
}
