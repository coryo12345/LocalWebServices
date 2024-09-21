package common

type Service interface {
	Start(urlPrefix string) error
	Shutdown()
}
