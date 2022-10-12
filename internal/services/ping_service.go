package services

type IPingService interface {
	Ping() string
}

type PingService struct {
}

func NewPingService() *PingService {
	return &PingService{}
}

func (service PingService) Ping() string {
	return "pong"
}
