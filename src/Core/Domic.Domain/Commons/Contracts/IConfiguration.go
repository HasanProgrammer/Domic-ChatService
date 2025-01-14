package DomainCommonContract

type IConfiguration interface {
	GetPostgreSqlConnectionString(key string) (string, error)
}
