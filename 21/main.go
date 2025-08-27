package main

import "fmt"

type Infrastructure struct{}

func (i Infrastructure) Execute() {
	fmt.Printf("%T is doing its work\n", i)
}

type InfrastructureAdapter struct {
	infra *Infrastructure
}

func NewInfrastructureAdapter(infra *Infrastructure) *InfrastructureAdapter {
	return &InfrastructureAdapter{infra: infra}
}

func (a *InfrastructureAdapter) Perform() {
	a.infra.Execute()
}

type AdapterPort interface {
	Perform()
}

type Service struct {
	infraAdapter AdapterPort
}

func NewService(adapter AdapterPort) *Service {
	return &Service{infraAdapter: adapter}
}

func (s *Service) DoBusinessLogic() {
	s.infraAdapter.Perform()
}

func main() {
	infra := new(Infrastructure)

	adapter := NewInfrastructureAdapter(infra)

	service := NewService(adapter)

	service.DoBusinessLogic()
}
