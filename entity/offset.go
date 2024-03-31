package entity

type Status int

const (
	NOT_PICKED  Status = iota
	NACK_STATUS Status = 1
	ACK_STATUS  Status = 2
)

type Offset struct {
	Offset int    `json:"offset"`
	Status Status `json:"status"`
}
