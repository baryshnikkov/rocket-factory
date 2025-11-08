package model

type PaymentMethod string

const (
	Unspecified   PaymentMethod = "UNSPECIFIED"    // Неизвестный способ
	Card          PaymentMethod = "CARD"           //	Банковская карта
	SBP           PaymentMethod = "SBP"            // Система быстрых платежей
	CreditCard    PaymentMethod = "CREDIT_CARD"    // Кредитная карта
	InvestorMoney PaymentMethod = "INVESTOR_MONEY" // Деньги инвестора (внутренний метод)
)

type PayOrderRequest struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod PaymentMethod
}

type PayOrderResponse struct {
	TransactionUUID string
}
