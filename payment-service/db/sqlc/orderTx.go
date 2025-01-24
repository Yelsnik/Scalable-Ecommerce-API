package db

import (
	"context"

	"payment-service/cart/cart-service"
	"payment-service/product/product-service"
	"payment-service/util"
)

type OrderTxResult struct {
	Order     Order                    `json:"order"`
	OrderItem OrderItem                `json:"order_item"`
	Cart      *cart.RemoveCartTxResult `json:"cart"`
	CartItem  *cart.CartItemResponse   `json:"cart_item"`
	Payment   Payment                  `json:"payment"`
}

type OrderTxParams struct {
	PaymentIntent   string  `json:"payment_intent"`
	UserName        string  `json:"user_name"`
	BuyerID         string  `json:"buyer_id"`
	SellerID        string  `json:"seller_id"`
	CartItemId      string  `json:"cart_item_id"`
	TotalPrice      float64 `json:"total_price"`
	DeliveryAddress string  `json:"delivery_address"`
	Country         string  `json:"country"`
	PaymentStatus   string  `json:"status"`
	OrderStatus     string  `json:"order_status"`
	GetCartItem     func(ctx context.Context, cartItemId string) (*cart.CartItemResponse, error)
	GetProductByID  func(ctx context.Context, productId string) (*product.ProductResponse, error)
	RemoveCartTx    func(ctx context.Context, cartItemId string) (*cart.RemoveCartTxResult, error)
}

func (store *SQLStore) CreateOrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// get cart item that has been paid for
		result.CartItem, err = arg.GetCartItem(ctx, arg.CartItemId)
		if err != nil {
			return err
		}
		productId := result.CartItem.Product

		// update payment status
		result.Payment, err = q.UpdatePaymentStatus(ctx, UpdatePaymentStatusParams{
			ID:     arg.PaymentIntent,
			Status: arg.PaymentStatus,
		})

		buyerId, err := util.ConvertStringToUUID(arg.BuyerID)
		if err != nil {
			return err
		}

		sellerId, err := util.ConvertStringToUUID(arg.SellerID)
		if err != nil {
			return err
		}

		// create the order
		result.Order, err = q.CreateOrders(ctx, CreateOrdersParams{
			UserName:        arg.UserName,
			BuyerID:         buyerId,
			SellerID:        sellerId,
			TotalPrice:      arg.TotalPrice,
			DeliveryAddress: arg.DeliveryAddress,
			Country:         arg.Country,
			Status:          arg.OrderStatus,
		})
		if err != nil {
			return err
		}

		// get the product associated with the cart
		product, err := arg.GetProductByID(ctx, productId)
		if err != nil {
			return err
		}

		// add the product to the order items
		result.OrderItem, err = q.CreateOrderitems(ctx, CreateOrderitemsParams{
			ItemName:     product.Product.ProductName,
			ItemSubTotal: float64(result.CartItem.SubTotal),
			Quantity:     result.CartItem.Quantity,
			ItemID:       product.Product.Id,
			OrderID:      result.Order.ID,
		})

		// remove the cart item
		result.Cart, err = arg.RemoveCartTx(ctx, arg.CartItemId)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
