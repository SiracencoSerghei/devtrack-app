package auth

type CurrentUser struct {
	UserID string
	Email  string
	Roles  []string
}

func (u *CurrentUser) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}