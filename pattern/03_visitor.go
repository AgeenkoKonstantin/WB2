package pattern

import "fmt"

/*
	Посетитель – это  паттерн проектирования,
	 который позволяет обойти разнородные объекты, добавлять в программу новые операции,
	  не изменяя классы этих  объектов
	ПЛЮСЫ:
		Посетитель может накапливать состояние при обходе структуры элементов
		Упрощает добавление операций, работающих со сложными структурами объектов
	МИНУСЫ:
		Затруднено добавлении нового класса, который имплементирует Visitable
*/
//Интерфейс посетителя предоставляет методы для работы с каждый из посещаемых типов
type Visitor interface {
	VisitA(e *ElemA)
	VisitB(e *ElemB)
}

//Реализация посетителя
type ConcreteVisitor struct {
}

func (v *ConcreteVisitor) VisitA(e *ElemA) {
	fmt.Println("visit elemA ")
	fmt.Println("При посещении объекта можно выполнить какую-то функцию visitor-а для конкретного типа посещаемого объекта")
}
func (v *ConcreteVisitor) VisitB(e *ElemB) {
	fmt.Println("visit elemB")
	fmt.Println("При посещении объекта можно выполнить какую-то функцию visitor-а для конкретного типа посещаемого объекта")
}

//Интерфейс посещаемого типа
type Visitable interface {
	Accept(v Visitor)
}

//Конкретная структура, которую можно посетить
type ElemA struct {
}

// Принимаем посетителя
func (e *ElemA) Accept(v Visitor) {
	v.VisitA(e)
}

//Конкретная структура, которую можно посетить
type ElemB struct {
}

// Принимаем посетителя
func (e *ElemB) Accept(v Visitor) {
	v.VisitB(e)
}

//Коллекция  visitable-ов
type ListVisitable struct {
	list []Visitable
}

//Посещаем все из коллекции
func (l *ListVisitable) Accept(v Visitor) {
	for _, e := range l.list {
		e.Accept(v)
	}
}

func useVisitor() {
	collection := new(ListVisitable)
	collection.list = append(collection.list, &ElemA{}, &ElemB{}, &ElemA{})

	collection.Accept(new(ConcreteVisitor))
}
