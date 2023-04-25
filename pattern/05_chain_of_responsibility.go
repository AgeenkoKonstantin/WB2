package pattern

import "fmt"

/* 
	Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса,
	 при этом давая шанс обработать этот запрос нескольким объектам.
	  Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.
	  ПЛЮСЫ:
				
	  МИНУСЫ:

*/

//Интерфейс хэндлера - элемента в цепочки обработчиков,
// в методе sendRequest - проверка на возможность обработки данным хэндлером или передача дальше по цепочке
type Handler interface {
	SendRequest(message int) string
}

//Реализация хэндлера, имеет сслыку на следующего обработчика в цепочке и реализует интерфейс Handler
type ConcreteHandlerA struct {
	next Handler
}
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

//Реализация хэндлера, имеет сслыку на следующего обработчика в цепочке и реализует интерфейс Handler
type ConcreteHandlerB struct {
	next Handler
}
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

//Реализация хэндлера, имеет сслыку на следующего обработчика в цепочке и реализует интерфейс Handler
type ConcreteHandlerC struct {
	next Handler
}
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}


func useChain() {

	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	result := handlers.SendRequest(2)
	fmt.Println(result)
}