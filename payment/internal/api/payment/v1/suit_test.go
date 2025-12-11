package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/baryshnikkov/rocket-factory/payment/internal/service/mocks"
)

type APISuite struct {
	suite.Suite
	ctx context.Context

	paymentService *mocks.PaymentService
	api            *api
}

func (s *APISuite) SetupTest() {
	s.ctx = context.Background()
	s.paymentService = mocks.NewPaymentService(s.T())
	s.api = NewAPI(s.paymentService)
}

func (s *APISuite) TearDownSuite() {}

func TestAPIInventoryIntegration(t *testing.T) {
	suite.Run(t, new(APISuite))
}
