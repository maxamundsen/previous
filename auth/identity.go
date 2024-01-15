package auth

import "time"

type Identity struct {
	UserId string
	// map[string]string instead of map[string]interface{} since claims can be stored in a database,
	// which only permits one datatype per column
	Claims          map[string]string
	IsAuthenticated bool
	RememberMe      bool
	UserAgent       string
	IpAddr          string
	LoginTime       time.Time
}
