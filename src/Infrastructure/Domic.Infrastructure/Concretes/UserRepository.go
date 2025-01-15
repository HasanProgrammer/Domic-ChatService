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

	var userModels []*Models.UserModel

	for _, entity := range *entities {
		userModels = append(userModels, &Models.UserModel{
			Id:          entity.Id,
			FullName:    entity.FullName,
			Email:       entity.Email,
			CreatedBy:   entity.CreatedBy,
			CreatedAt:   entity.CreatedAt,
			CreatedRole: entity.CreatedRole,
			Version:     entity.Version,
			IsActive:    entity.IsActive,
			IsDeleted:   entity.IsDeleted,
		})
	}

	queryResult := userRepository.db.CreateInBatches(userModels, len(*entities))

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}

}

func (userRepository *UserRepository) Change(entity *DomainUserEntity.User) DomainCommonDTO.Result[bool] {

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

	queryResult := userRepository.db.Save(model)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}
}

func (userRepository *UserRepository) Remove(entity *DomainUserEntity.User) DomainCommonDTO.Result[bool] {

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

	queryResult := userRepository.db.Delete(model)

	return DomainCommonDTO.Result[bool]{
		Error:  queryResult.Error,
		Result: queryResult.Error != nil,
	}
}

func (userRepository *UserRepository) FindById(id string) DomainCommonDTO.Result[*DomainUserEntity.User] {

	var model *Models.UserModel

	queryResult := userRepository.db.First(model, "id = ?", id)

	return DomainCommonDTO.Result[*DomainUserEntity.User]{
		Error: queryResult.Error,
		Result: &DomainUserEntity.User{
			Id:          model.Id,
			FullName:    model.FullName,
			Email:       model.Email,
			CreatedBy:   model.CreatedBy,
			CreatedAt:   model.CreatedAt,
			CreatedRole: model.CreatedRole,
			Version:     model.Version,
			IsActive:    model.IsActive,
			IsDeleted:   model.IsDeleted,
		},
	}

}

func (userRepository *UserRepository) FindAllPaginated(paginationRequest *DomainCommonDTO.PaginationRequest) DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[DomainUserEntity.User]] {

	var models *[]Models.UserModel

	take := int(paginationRequest.PageSize)
	skip := int((paginationRequest.PageIndex - 1) * paginationRequest.PageSize)

	queryResult := userRepository.db.Limit(take).Offset(skip).Find(&models)

	var users []DomainUserEntity.User

	for _, entity := range *models {
		users = append(users, DomainUserEntity.User{
			Id:          entity.Id,
			FullName:    entity.FullName,
			Email:       entity.Email,
			CreatedBy:   entity.CreatedBy,
			CreatedAt:   entity.CreatedAt,
			CreatedRole: entity.CreatedRole,
			Version:     entity.Version,
			IsActive:    entity.IsActive,
			IsDeleted:   entity.IsDeleted,
		})
	}

	return DomainCommonDTO.Result[DomainCommonDTO.PaginationResponse[DomainUserEntity.User]]{
		Error: queryResult.Error,
		Result: DomainCommonDTO.PaginationResponse[DomainUserEntity.User]{
			PageSize:  paginationRequest.PageSize,
			PageIndex: paginationRequest.PageIndex,
			Items:     &users,
		},
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
