package DomainCommonContract

type IIdentityUser interface {
	GetUserIdentity() string
	GetUserRoles() []string
}
