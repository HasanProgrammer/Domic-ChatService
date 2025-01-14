package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	transaction *gorm.DB
}

func (u *UnitOfWork) TransactionStart() interface{} {
	return u.transaction
}

func (u *UnitOfWork) TransactionCommit() DomainCommonDTO.Result[bool] {

	if u.transaction != nil {

		queryResult := u.transaction.Commit()

		if queryResult.Error != nil {
			return DomainCommonDTO.Result[bool]{
				Error:  queryResult.Error,
				Result: false,
			}
		}

		return DomainCommonDTO.Result[bool]{
			Error:  nil,
			Result: true,
		}

	}

	return DomainCommonDTO.Result[bool]{
		Error:  nil,
		Result: false,
	}

}

func (u *UnitOfWork) TransactionRollback() DomainCommonDTO.Result[bool] {

	if u.transaction != nil {

		queryResult := u.transaction.Rollback()

		if queryResult.Error != nil {
			return DomainCommonDTO.Result[bool]{
				Error:  queryResult.Error,
				Result: false,
			}
		}

		return DomainCommonDTO.Result[bool]{
			Error:  nil,
			Result: true,
		}

	}

	return DomainCommonDTO.Result[bool]{
		Error:  nil,
		Result: false,
	}

}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {

	unitOfWork := &UnitOfWork{}

	unitOfWork.transaction = db.Begin()

	return unitOfWork

}
