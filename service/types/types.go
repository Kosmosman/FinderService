package types

import "sync"

type Cache struct {
	Mutex sync.RWMutex
	Data  map[string][]byte
}

type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
}

type Payment struct {
	Transaction  string  `json:"transaction" validate:"required"`
	RequestId    string  `json:"request_id,omitempty" validate:"omitempty"`
	Currency     string  `json:"currency" validate:"required"`
	Provider     string  `json:"provider" validate:"required"`
	Amount       float64 `json:"amount" validate:"min=0"`
	PaymentDt    int64   `json:"payment_dt" validate:"min=0"`
	Bank         string  `json:"bank" validate:"required"`
	DeliveryCost int     `json:"delivery_cost" validate:"min=0"`
	GoodsTotal   int     `json:"goods_total" validate:"min=1"`
	CustomFee    int     `json:"custom_fee" validate:"min=0"`
}

type Item struct {
	ChrtID      int     `json:"chrt_id" validate:"min=0"`
	TrackNumber string  `json:"track_number" validate:"required"`
	Price       float64 `json:"price" validate:"min=0"`
	RID         string  `json:"rid" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Sale        int     `json:"sale" validate:"min=0"`
	Size        string  `json:"size" validate:"required"`
	TotalPrice  float64 `json:"total_price" validate:"min=0"`
	NmID        int     `json:"nm_id" validate:"min=0"`
	Brand       string  `json:"brand" validate:"required"`
	Status      int     `json:"status" validate:"min=0"`
}

type Order struct {
	OrderUID          string   `json:"order_uid" validate:"required"`
	TrackNumber       string   `json:"track_number" validate:"required"`
	Entry             string   `json:"entry" validate:"required"`
	Delivery          Delivery `json:"delivery" validate:"required"`
	Payment           Payment  `json:"payment" validate:"required"`
	Items             []Item   `json:"items" validate:"required,dive,required"`
	Locale            string   `json:"locale" validate:"required"`
	InternalSignature string   `json:"internal_signature,omitempty" validate:"omitempty"`
	CustomerID        string   `json:"customer_id" validate:"required"`
	DeliveryService   string   `json:"delivery_service" validate:"required"`
	ShardKey          string   `json:"shardkey" validate:"required"`
	SMID              int      `json:"sm_id" validate:"min=0"`
	DateCreated       string   `json:"date_created" validate:"required"`
	OofShard          string   `json:"oof_shard" validate:"required"`
}
