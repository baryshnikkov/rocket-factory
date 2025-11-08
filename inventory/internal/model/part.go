package model

import "time"

type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      Metadata
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

type Category string

const (
	CategoryUnspecified Category = "UNKNOWN"  // Неизвестная категория
	CategoryEngine      Category = "ENGINE"   // Двигатель
	CategoryFuel        Category = "FUEL"     //  Топливо
	CategoryPorthole    Category = "PORTHOLE" // Иллюминатор
	CategoryWing        Category = "WING"     // Крыло

)

type Dimensions struct {
	Length float64 // Длина в см
	Width  float64 // Ширина в см
	Height float64 // Высота в см
	Weight float64 // Масса в кг
}

type Manufacturer struct {
	Name    string
	Country string
	Website string
}

type Metadata struct {
	StringValue *string
	Int64Value  *int64
	DoubleValue *float64
	BoolValue   *bool
}

type PartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
