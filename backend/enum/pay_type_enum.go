package enum

type PayType int

const (
	Bank   PayType = 0
	WeChat PayType = 1
	AliPay PayType = 2
)

func (p PayType) String() string {
	switch p {
	case Bank:
		return "Bank"
	case WeChat:
		return "WeChat"
	case AliPay:
		return "AliPay"
	default:
		return "UNKNOW"
	}
}
