package InfrastructureConcrete

import "net/http"

type IdentityUser struct {
	claims map[string]interface{}
}

func (identity *IdentityUser) GetUserIdentity() string {
	return identity.claims["UserId"].(string)
}

func (identity *IdentityUser) GetUserRoles() []string {
	return identity.claims["Roles"].([]string)
}

func NewIdentityUser(r *http.Request) *IdentityUser {

	token := r.Header.Get("Authorization")

	jwtToken := NewJsonWebToken()

	identity := &IdentityUser{
		claims: jwtToken.GetClaims(token).Result,
	}

	return identity
}
