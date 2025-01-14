package DomainCommonContract

import DomainCommonEntity "Domic.Domain/Commons/Entities"

type IConsumerEventHandler interface {
	Handle(event *DomainCommonEntity.Event) error
}
