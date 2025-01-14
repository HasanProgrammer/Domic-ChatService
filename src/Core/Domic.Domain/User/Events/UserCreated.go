package DomainUserEvent

import "time"

type UserCreated struct {
	Id          string    `json:"Id"`
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Email       string    `json:"Email"`
	CreatedBy   string    `json:"CreatedBy"`
	CreatedRole string    `json:"CreatedRole"`
	CreatedAt   time.Time `json:"CreatedAt"`
}
