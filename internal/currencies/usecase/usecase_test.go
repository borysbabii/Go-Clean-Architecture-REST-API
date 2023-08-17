package usecase

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/currencies/mock"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/internal/models"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/logger"
	"github.com/borysbabii/Go-Clean-Architecture-REST-API/pkg/utils"
)

func TestCurrencyUC_GetCurrencies(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockCurrencyRepo := mock.NewMockRepository(ctrl)
	currencyUC := NewCurrenciesUseCase(nil, mockCurrencyRepo, apiLogger)

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "currencyUC.GetCurrencies")
	defer span.Finish()
	query := &utils.PaginationQuery{
		Size:    10,
		Page:    1,
		OrderBy: "",
	}

	currenciesList := &models.CurrenciesList{}

	mockCurrencyRepo.EXPECT().GetCurrencies(ctxWithTrace, query).Return(currenciesList, nil)

	currency, err := currencyUC.GetCurrencies(ctx, query)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, currency)
}
