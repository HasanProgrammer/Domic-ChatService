package UseCaseCommonContract

import "time"

type IInternalDistributedCache interface {
	Set(object interface{}, key string, timeout time.Duration) bool
	Get(key string, target interface{}) error
	Delete(key string) bool
}
