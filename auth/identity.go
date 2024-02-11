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

func EnsureHasClaims(identity *Identity, req map[string]string) bool {
	for key, value := range req {
		if identity.Claims[key] != value {
			return false
		}
	}
	return true
}
