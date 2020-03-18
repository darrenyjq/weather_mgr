package helper

type (
	TableTickets struct {
		TicketId int64 `json:"ticket_id"`
		Uid int64 `json:"uid"`
		Expired	int64 `json:"expired"`
		Status  int64 `json:"status"`
		Event  string `json:"event"`
	}
)
