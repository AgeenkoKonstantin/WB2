package pattern

import (
	"fmt"
)
//Команда – это  паттерн, который превращает запросы в объекты,
// позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь,
// логировать их, а также поддерживать отмену операций

/*
 ПЛЮСЫ: Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их  выполняют
		Позволяет реализовать простую отмену и повтор операций
		Позволяет реализовать отложенный запуск операций

 Минусы:
 	Усложняет код программы из-за введения множества дополнительных классов
*/

// Интерфейс команды
type Command interface {
	Execute()
}

// Реализация команды
type ToggleOnCommand struct {
	receiver *Receiver
}

// Эта функция инициируется инвокером и запускает определенный метод у ресивера команды
func (c *ToggleOnCommand) Execute() {
	c.receiver.ToggleOn()
}

// Реализация команды
type ToggleOffCommand struct {
	receiver *Receiver
}

// Эта функция инициируется инвокером и запускает определенный метод у ресивера команды
func (c *ToggleOffCommand) Execute() {
	c.receiver.ToggleOff()
}

// Структура ресивера команды, имеет методы обработки каждой команды
type Receiver struct {
}

// Обработка первой команды
func (r *Receiver) ToggleOn() {
	fmt.Println("Toggle On")
}

// Обработка второй команды
func (r *Receiver) ToggleOff() {
	fmt.Println("Toggle Off")
}

// Структура инициатора команд, назначает ресиверов и запускает исполнение команд
type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// запускает исполнение команд
func (i *Invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func useCommand() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})

	invoker.Execute()
}
