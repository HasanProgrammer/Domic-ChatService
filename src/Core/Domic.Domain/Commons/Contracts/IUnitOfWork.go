package DomainCommonContract

import "Domic.Domain/Commons/DTOs"

type IUnitOfWork interface {
	TransactionStart() interface{}
	TransactionCommit() DomainCommonDTO.Result[bool]
	TransactionRollback() DomainCommonDTO.Result[bool]
}
