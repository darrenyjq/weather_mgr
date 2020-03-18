package xconfig

type (
	EventConfig struct {
		Owner	string `json:"owner"`
		Event   string `json:"event"`
		Direc   string `json:"direc"`
		DayLimit int64 `json:"day_limit"`
		OnceLimit	int64 `json:"once_limit"`
		CurrentNum int64 `json:"current_num"`
		TicketExpiredDay	int `json:"ticket_expired_day"`
	}
	WithdrawGoodsConfig struct {
		Show   bool `json:"show"`
		GoodsId	int64 `json:"goods_id"`
		Bank	string `json:"bank"`
		Amount  int64 `json:"amount"`
		TicketNum int64 `json:"ticket_num"`
		GoodsType string `json:"goods_type"`
		MaxWithdrawNum int64 `json:"max_withdraw_num"`
		CurrentWithdrawNum int64 `json:"current_withdraw_num"`
	}
	WithdrawConfig struct {
		AmountEvents	[]EventConfig `json:"amount_events"`
		TicketEvents	[]EventConfig `json:"ticket_events"`
		Goods []WithdrawGoodsConfig `json:"goods"`
		Rules 	string `json:"rules"`
	}
)
