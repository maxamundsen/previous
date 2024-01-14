package auth

type Identity struct {
	IsAuthenticated bool
	RememberMe      bool

	// map[string]string over map[string]interface{} since claims can be stored in a database,
	// which only permits one datatype per column
	Claims map[string]string

	UserId string
}
