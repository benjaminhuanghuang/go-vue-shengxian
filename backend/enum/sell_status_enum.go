package enum

type SellStatus int

const (
	Selling  SellStatus = 0
	StopSell SellStatus = 1
)

func (p SellStatus) String() string {
	switch p {
	case Selling:
		return "Selling"
	case StopSell:
		return "StopSell"
	default:
		return "UNKNOW"
	}
}
