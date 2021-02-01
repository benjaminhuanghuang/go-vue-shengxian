package model

type Product struct {
	OrderID      string  `json: "orderId" gorm:"column:order_id"`
	UserID       string  `json: "userId" gorm:"column:user_id"`
	NichName     string  `json: "nick_name" gorm:"column:nick_name"`
	Mobile       string  `json: "mobile" gorm:"column:mobile"`
	TotalPrice   string  `json: "totalPrice" gorm:"column:total_price"`
	PayStatus    string  `json: "payStatus" gorm:"column:pay_status"`
	PayType      float32 `json: "payType" gorm:"column:pay_type"`
	PayTime      string  `json: "payTime" gorm:"column:pay_time"`
	OrderStatus  string  `json: "orderStatus" gorm:"column:order_status"`
	ExtraInfo    string  `json: "extraInfo" gorm:"column:extraInfo"`
	UsereAddress string  `json: "userAddress" gorm:"column:user_address"`
	isDeleted    string  `json: "isDeleted" gorm:"column:is_deleted"`
}
