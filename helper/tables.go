package helper

type (
	TableTickets struct {
		TicketId int64  `json:"ticket_id"`
		uid      uint64 `json:"uid"`
		Expired  int64  `json:"expired"`
		Status   int64  `json:"status"`
		Event    string `json:"event"`
	}
)
