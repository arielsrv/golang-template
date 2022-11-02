package services

type IPingService interface {
	Ping() string
}

type PingService struct {
}

func NewPingService() IPingService {
	return PingService{}
}

func (service PingService) Ping() string {
	return "pong"
}
