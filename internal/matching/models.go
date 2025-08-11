// internal/matching/models.go

package matching

// Side 定义订单方向
type Side int

const (
	Buy  Side = 1
	Sell Side = -1
)

// Order 定义一个订单
type Order struct {
	ID        int64
	UserID    int64
	Side      Side
	Price     float64 // 对于金融计算，使用 decimal/big.Float 类型更精确
	Amount    float64
	Timestamp int64
}

// OrderBook 定义订单簿
type OrderBook struct {
	Bids []*Order // 买单，价格从高到低
	Asks []*Order // 卖单，价格从低到高
}

// Trade 定义一笔成交记录
type Trade struct {
	TakerOrderID int64 // 吃单者
	MakerOrderID int64 // 挂单者
	Price        float64
	Amount       float64
	Timestamp    int64
}
