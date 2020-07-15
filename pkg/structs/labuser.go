package structs

import (
	"fmt"
	"time"
)

// LabUser is the glu
type LabUser struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	State           string    `json:"state"`
	AvatarURL       string    `json:"avatar_url"`
	WebURL          string    `json:"web_url"`
	CreatedAt       time.Time `json:"created_at"`
	Bio             string    `json:"bio"`
	Location        string    `json:"location"`
	PublicEmail     string    `json:"public_email"`
	Skype           string    `json:"skype"`
	Linkedin        string    `json:"linkedin"`
	Twitter         string    `json:"twitter"`
	WebsiteURL      string    `json:"website_url"`
	Organization    string    `json:"organization"`
	LastSignInAt    time.Time `json:"last_sign_in_at"`
	ConfirmedAt     time.Time `json:"confirmed_at"`
	LastActivityOn  string    `json:"last_activity_on"`
	Email           string    `json:"email"`
	ThemeID         int       `json:"theme_id"`
	ColorSchemeID   int       `json:"color_scheme_id"`
	ProjectsLimit   int       `json:"projects_limit"`
	CurrentSignInAt time.Time `json:"current_sign_in_at"`
	Identities      []struct {
		Provider  string `json:"provider"`
		ExternUID string `json:"extern_uid"`
	} `json:"identities"`
	CanCreateGroup   bool `json:"can_create_group"`
	CanCreateProject bool `json:"can_create_project"`
	TwoFactorEnabled bool `json:"two_factor_enabled"`
	External         bool `json:"external"`
	PrivateProfile   bool `json:"private_profile"`
}

// GetID for LabUser
func (l LabUser) GetID() string {
	return fmt.Sprintf("%v+%v", l.ID, l.Username)
}
