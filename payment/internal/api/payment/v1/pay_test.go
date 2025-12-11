package v1

import (
	"context"
	"errors"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/baryshnikkov/rocket-factory/payment/internal/model"
	paymentV1 "github.com/baryshnikkov/rocket-factory/shared/pkg/proto/payment/v1"
)

func (s *APISuite) TestPayOrderSuccess() {
	var (
		orderUUID       = gofakeit.UUID()
		userUUID        = gofakeit.UUID()
		transactionUUID = gofakeit.UUID()

		req = &paymentV1.PayOrderRequest{
			OrderUuid:     orderUUID,
			UserUuid:      userUUID,
			PaymentMethod: paymentV1.PaymentMethod_PAYMENT_METHOD_CARD,
		}

		reqModel = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: model.Card,
		}

		resModel = model.PayOrderResponse{
			TransactionUUID: transactionUUID,
		}

		expectedResponse = &paymentV1.PayOrderResponse{
			TransactionUuid: transactionUUID,
		}
	)

	s.paymentService.On("PayOrder", s.ctx, reqModel).Return(resModel, nil)

	res, err := s.api.PayOrder(s.ctx, req)

	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(expectedResponse.TransactionUuid, res.TransactionUuid)
}

func (s *APISuite) TestPayOrderPaymentInternalError() {
	var (
		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()

		req = &paymentV1.PayOrderRequest{
			OrderUuid:     orderUUID,
			UserUuid:      userUUID,
			PaymentMethod: paymentV1.PaymentMethod_PAYMENT_METHOD_SBP,
		}

		reqModel = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: model.SBP,
		}

		expectedErr = model.ErrPaymentInternalError
	)

	s.paymentService.On("PayOrder", s.ctx, reqModel).Return(model.PayOrderResponse{}, expectedErr)

	res, err := s.api.PayOrder(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
	s.Require().Contains(err.Error(), "Payment service error")
	s.Require().Nil(res)
}

func (s *APISuite) TestPayOrderContextDeadlineExceeded() {
	var (
		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()

		req = &paymentV1.PayOrderRequest{
			OrderUuid:     orderUUID,
			UserUuid:      userUUID,
			PaymentMethod: paymentV1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD,
		}

		reqModel = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: model.CreditCard,
		}

		expectedErr = context.DeadlineExceeded
	)

	s.paymentService.On("PayOrder", s.ctx, reqModel).Return(model.PayOrderResponse{}, expectedErr)

	res, err := s.api.PayOrder(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(codes.Unavailable, status.Code(err))
	s.Require().Contains(err.Error(), "Paymnent service timeout")
	s.Require().Nil(res)
}

func (s *APISuite) TestPayOrderContextCanceled() {
	var (
		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()

		req = &paymentV1.PayOrderRequest{
			OrderUuid:     orderUUID,
			UserUuid:      userUUID,
			PaymentMethod: paymentV1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY,
		}

		reqModel = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: model.InvestorMoney,
		}

		expectedErr = context.Canceled
	)

	s.paymentService.On("PayOrder", s.ctx, reqModel).Return(model.PayOrderResponse{}, expectedErr)

	res, err := s.api.PayOrder(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(codes.Unavailable, status.Code(err))
	s.Require().Contains(err.Error(), "Paymnent service timeout")
	s.Require().Nil(res)
}

func (s *APISuite) TestPayOrderUnknownError() {
	var (
		orderUUID = gofakeit.UUID()
		userUUID  = gofakeit.UUID()

		req = &paymentV1.PayOrderRequest{
			OrderUuid:     orderUUID,
			UserUuid:      userUUID,
			PaymentMethod: paymentV1.PaymentMethod_PAYMENT_METHOD_CARD,
		}

		reqModel = model.PayOrderRequest{
			OrderUUID:     orderUUID,
			UserUUID:      userUUID,
			PaymentMethod: model.Card,
		}

		expectedErr = errors.New("unknown error")
	)

	s.paymentService.On("PayOrder", s.ctx, reqModel).Return(model.PayOrderResponse{}, expectedErr)

	res, err := s.api.PayOrder(s.ctx, req)

	s.Require().Error(err)
	s.Require().Equal(expectedErr, err)
	s.Require().Nil(res)
}
