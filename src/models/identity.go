package models

import (
	"time"
)

type Identity struct {
	UserId        int
	SecurityStamp string
	Authenticated bool
	RememberMe    bool
	Expiration    time.Time
}
