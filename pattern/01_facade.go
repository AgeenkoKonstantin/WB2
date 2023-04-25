package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Фасад предоставляет общий упрощенный интерфейс для некоторой сложной системы

ПЛЮСЫ:  изолирует клиента от сложной системы, перенос этой сложной системы также упрощен,

	потому что клиент пользуется фасадом

МИНУСЫ: фасад может стать так называемым god object-ом, так как имеет множество осей отвественностей
*/

type requestHandler interface {
	getReq()
	sendResp()
}

type requestHandlerImpl struct {
}

func (h *requestHandlerImpl) getReq() {
	fmt.Println("get request")
}

func (h *requestHandlerImpl) sendResp() {
	fmt.Println("send response")
}

type service interface {
	somelogic()
}
type serviceImpl struct {
}

func (s *serviceImpl) somelogic() {
	fmt.Println("some business logic")
}

type logger interface {
	log(info string)
}

type loggerImpl struct {
}

func (l *loggerImpl) log(info string) {
	fmt.Println(time.Now().String() + " " + info)
}

type ServiceFacade struct {
	logger
	service
	requestHandler
}

func (s *ServiceFacade) someWork() {
	s.log("... logging ....")
	s.getReq()
	s.somelogic()
	s.sendResp()
}

func useFacade() {
	facade := &ServiceFacade{
		&loggerImpl{},
		&serviceImpl{},
		&requestHandlerImpl{},
	}

	facade.someWork()
}
