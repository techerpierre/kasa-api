package entities

type AuthStatus string

const (
	AuthStatus_CLIENT    = "CLIENT"
	AuthStatus_OWNER     = "OWNER"
	AuthStatus_MODERATOR = "MODERATOR"
	AuthStatus_ADMIN     = "ADMIN"
)
