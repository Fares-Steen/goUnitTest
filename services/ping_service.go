package services

import "fmt"

type pingService interface {
	HandlePing() (string, error)
}
type pingServiceImpl struct{}

const (
	pong = "pong"
)

var (
	PingService pingService = pingServiceImpl{}
)

func (service pingServiceImpl) HandlePing() (string, error) {
	fmt.Println("doing some complex stuff...")
	return pong, nil
}
