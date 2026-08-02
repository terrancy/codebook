package structural

// 外观模式（Facade）：为子系统中的一组接口提供一个一致的界面，
// 外观定义了一个高层接口，让子系统更容易使用。
//
// 典型场景：复杂 SDK / 微服务聚合调用，给调用方一个“一键”入口。

// 子系统 A：库存
type Inventory struct{}

func (Inventory) Check(sku string) bool { return sku != "" }

// 子系统 B：支付
type Payment struct{}

func (Payment) Pay(amount int) bool { return amount > 0 }

// 子系统 C：物流
type Shipping struct{}

func (Shipping) Dispatch(sku string) string { return "dispatched:" + sku }

// OrderFacade 外观：把下单流程对调用方隐藏。
type OrderFacade struct {
	inv Inventory
	pay Payment
	ship Shipping
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{inv: Inventory{}, pay: Payment{}, ship: Shipping{}}
}

// PlaceOrder 一个入口完成：查库存 -> 扣款 -> 发货。
func (f *OrderFacade) PlaceOrder(sku string, amount int) string {
	if !f.inv.Check(sku) {
		return "out of stock"
	}
	if !f.pay.Pay(amount) {
		return "payment failed"
	}
	return f.ship.Dispatch(sku)
}
