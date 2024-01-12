package auth

type Identity struct {
	IsAuthenticated bool
	RememberMe      bool
	Role            string
	Username        string
}
