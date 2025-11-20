package orders

import (
	"context"

	repo "github.com/Yujonpradhananga/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductId int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerID"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
