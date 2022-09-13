package application

type IPingService interface {
	Ping() string
}

type PingService struct {
}

func (service PingService) Ping() string {
	return "pong"
}

func NewPingService() IPingService {
	return &PingService{}
}
