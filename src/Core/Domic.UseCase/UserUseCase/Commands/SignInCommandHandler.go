package UserUseCaseCommand

import (
	"Domic.UseCase/Commons/Contracts"
	"Domic.UseCase/UserUseCase/DTOs"
	"errors"
	"regexp"
	"time"
)

type SignInCommandHandler struct {
	distributedCache UseCaseCommonContract.IInternalDistributedCache
	fullName         string
	eMail            string
}

func (commandHandler *SignInCommandHandler) Handle() bool {
	newPublicUser := UserUseCaseDTO.PublicUser{
		FullName: commandHandler.fullName,
		EMail:    commandHandler.eMail,
	}

	return commandHandler.distributedCache.Set(newPublicUser, "PublicUser", 30*time.Minute)
}

func NewSignInCommandHandler(distributedCache UseCaseCommonContract.IInternalDistributedCache, fullName string, eMail string) (*SignInCommandHandler, []error) {
	var errs []error

	if len(fullName) <= 3 || len(fullName) >= 100 {
		errs = append(errs, errors.New("نام و نام خانوادگی باید بیشتر از 3 و کمتر از 100 عبارت داشته باشد!"))
	}

	regexCompile := regexp.MustCompile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")

	if !regexCompile.MatchString(eMail) {
		errs = append(errs, errors.New("پست الکترونیکی ارسالی معتبر نمی باشد!"))
	}

	if len(errs) > 0 {
		return nil, errs
	}

	return &SignInCommandHandler{
		distributedCache: distributedCache,
		fullName:         fullName,
		eMail:            eMail,
	}, nil
}
