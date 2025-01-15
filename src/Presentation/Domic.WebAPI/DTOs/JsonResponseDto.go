package WebAPIDTO

type JsonResponseDto struct {
	Code    int
	Message string
	Body    interface{}
}
