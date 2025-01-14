package DomainCommonContract

import (
	"Domic.Domain/Commons/DTOs"
)

type IRepository[TIdentity any, TEntity any] interface {
	Add(entity *TEntity) DomainCommonDTO.Result[bool]
	AddRange(entities *[]TEntity) DomainCommonDTO.Result[bool]
	Change(entity *TEntity) DomainCommonDTO.Result[bool]
	Remove(entity *TEntity) DomainCommonDTO.Result[bool]
	FindById(id TIdentity) DomainCommonDTO.Result[*TEntity]
	FindAllPaginated(paginationRequest *DomainCommonDTO.PaginationRequest) DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[TEntity]]
}
