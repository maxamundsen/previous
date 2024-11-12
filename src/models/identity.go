package models

import (
	"time"
)

type Identity struct {
	User          User
	Authenticated bool
	RememberMe    bool
	Expiration    time.Time
}
