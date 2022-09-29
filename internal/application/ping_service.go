package application

type IPingService interface {
	Ping() string
}

type PingService struct {
}

func NewPingService() *PingService {
	return &PingService{}
}

func (p PingService) Ping() string {
	return "pong"
}
