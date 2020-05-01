package playermodel

// Player :
type Player struct {
	ID             int    `json:"id"`
	PlayerID       string `json:"playerId"`
	UserName       string `json:"userName"`
	EmailAddress   string `json:"emailAddress"`
	DisplayName    string `json:"displayName"`
	AvatarURL      string `json:"avatarUrl"`
	AvatarBlurHash string `json:"avatarBlurHash"`
}

// PlayerAuthentication :
type PlayerAuthentication struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// PlayerRegistration :
type PlayerRegistration struct {
	UserName        string `json:"userName,omitempty" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
