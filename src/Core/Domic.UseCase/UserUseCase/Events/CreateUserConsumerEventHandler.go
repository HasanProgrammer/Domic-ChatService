package DomainUserUseCaseEvents

import (
	"Domic.Domain/Commons/Entities"
	"Domic.Domain/User/Contracts"
	"Domic.Domain/User/Entities"
	"Domic.Domain/User/Events"
	"Domic.Infrastructure/Concretes"
)

type CreateUserConsumerEventHandler struct {
	unitOfWork              InfrastructureConcrete.UnitOfWork
	globalIdentityGenerator InfrastructureConcrete.GlobalIdentityGenerator
	serializer              InfrastructureConcrete.Serializer
	userCommandRepository   UserContract.IUserCommandRepository
}

func (consumer *CreateUserConsumerEventHandler) Handle(event *DomainCommonEntity.Event) error {

	var userCreatedEvent *DomainUserEvent.UserCreated

	err := consumer.serializer.Deserialize(event.GetPayload(), &userCreatedEvent)

	if err != nil {
		return err
	}

	result := consumer.userCommandRepository.Add(&DomainUserEntity.User{
		Id:          userCreatedEvent.Id,
		FullName:    userCreatedEvent.FirstName + " " + userCreatedEvent.LastName,
		Email:       userCreatedEvent.Email,
		CreatedBy:   event.GetCreatedBy(),
		CreatedAt:   event.GetCreatedAt(),
		CreatedRole: event.GetCreatedRole(),
		Version:     consumer.globalIdentityGenerator.Generate(),
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewCreateUserConsumer(serializer InfrastructureConcrete.Serializer, userCommandRepository UserContract.IUserCommandRepository,
	globalIdentityGenerator InfrastructureConcrete.GlobalIdentityGenerator, unitOfWork InfrastructureConcrete.UnitOfWork,
) *CreateUserConsumerEventHandler {
	return &CreateUserConsumerEventHandler{
		globalIdentityGenerator: globalIdentityGenerator,
		serializer:              serializer,
		userCommandRepository:   userCommandRepository,
		unitOfWork:              unitOfWork,
	}
}
