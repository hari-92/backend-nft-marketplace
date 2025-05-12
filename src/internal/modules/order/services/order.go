package order_services

type OrderService interface {
}

func NewOrderService() OrderService {
	return &orderService{}
}

type orderService struct {
}
