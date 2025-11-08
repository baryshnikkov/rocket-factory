package model

import "time"

type OrderDto struct {
	UUID            string
	UserUUID        string
	PartsUUIDs      []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *PaymentMethod
	Status          OrderStatus
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type OrderCreationInfo struct {
	OrderUUID  string
	TotalPrice float64
}

type OrderUpdateInfo struct {
	TotalPrice      *float64
	TransactionUUID *string
	PaymentMethod   *PaymentMethod
	Status          *OrderStatus
}

type PaymentMethod string

const (
	Unknown       PaymentMethod = "UNSPECIFIED"    // Неизвестный способ
	Card          PaymentMethod = "CARD"           //	Банковская карта
	SBP           PaymentMethod = "SBP"            // Система быстрых платежей
	CreditCard    PaymentMethod = "CREDIT_CARD"    // Кредитная карта
	InvestorMoney PaymentMethod = "INVESTOR_MONEY" // Деньги инвестора (внутренний метод)
)

type OrderStatus string

const (
	OrderStatusPendingPayment OrderStatus = "PENDING_PAYMENT"
	OrderStatusPaid           OrderStatus = "PAID"
	OrderStatusCancelled      OrderStatus = "CANCELLED"
)
