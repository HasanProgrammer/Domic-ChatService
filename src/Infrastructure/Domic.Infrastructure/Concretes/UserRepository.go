package InfrastructureConcrete

import (
	"Domic.Domain/Commons/DTOs"
	"Domic.Domain/User/Entities"
	"Domic.Persistence/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (userRepository *UserRepository) Add(entity *DomainUserEntity.User) DomainCommonDTO.Result[bool] {

	model := Models.UserModel{
		Id:          entity.Id,
		FullName:    entity.FullName,
		Email:       entity.Email,
		CreatedBy:   entity.CreatedBy,
		CreatedAt:   entity.CreatedAt,
		CreatedRole: entity.CreatedRole,
		Version:     entity.Version,
		IsActive:    entity.IsActive,
		IsDeleted:   entity.IsDeleted,
	}

	queryResult := userRepository.db.Create(&model)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error == nil,
	}

}

func (userRepository *UserRepository) AddRange(entities *[]DomainUserEntity.User) DomainCommonDTO.Result[bool] {

	queryResult := userRepository.db.CreateInBatches(entities, len(*entities))

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}

}

func (userRepository *UserRepository) Change(entity *DomainUserEntity.User) DomainCommonDTO.Result[bool] {

	queryResult := userRepository.db.Save(entity)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}
}

func (userRepository *UserRepository) Remove(entity *DomainUserEntity.User) DomainCommonDTO.Result[bool] {

	queryResult := userRepository.db.Delete(entity)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}
}

func (userRepository *UserRepository) FindById(id string) DomainCommonDTO.Result[*DomainUserEntity.User] {

	var user *DomainUserEntity.User

	queryResult := userRepository.db.First(user, "id = ?", id)

	return DomainCommonDTO.Result[*DomainUserEntity.User]{
		Error:  queryResult.Error,
		Result: user,
	}

}

func (userRepository *UserRepository) FindAllPaginated(paginationRequest *DomainCommonDTO.PaginationRequest) DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[DomainUserEntity.User]] {

	var users *[]DomainUserEntity.User

	take := int(paginationRequest.PageSize)
	skip := int((paginationRequest.PageIndex - 1) * paginationRequest.PageSize)

	queryResult := userRepository.db.Limit(take).Offset(skip).Find(&users)

	return DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[DomainUserEntity.User]]{
		Error: queryResult.Error,
		Result: DomainCommonDTO.PaginationResponse[DomainUserEntity.User]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     users,
		},
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
