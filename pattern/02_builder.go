package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Builder позволяет поэтапно создавать сложные объекты
	Плюсы: удобно использовать при конструировании сложного объекта, если у него множество необязательных полей,
		 перед созданием объекта можно проверить, что он верно создан
	Минусы:
*/

// Структура, которую нужно создать
type Document struct {
	Content string
}

// Интерфейс билдера, который предоставляет методы для создание отдельных частей
type Builder1 interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Реализация билдера
type ConcreteBuilder struct {
	doc *Document
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		new(Document),
	}
}
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.doc.Content += "<header>" + str + "</header>"
}
func (b *ConcreteBuilder) MakeBody(str string) {
	b.doc.Content += "<article>" + str + "</article>"
}
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.doc.Content += "<footer>" + str + "</footer>"
}
func (b *ConcreteBuilder) build() *Document {
	b.validation()
	return b.doc
}
func (b *ConcreteBuilder) validation() {
	fmt.Println("Проверяем что-то объект создан нормально")
}

func UseBuilder() {
	b := NewConcreteBuilder()

	b.MakeHeader("Header")
	b.MakeBody("Body")
	b.MakeFooter("Footer")

	doc := b.build()
	fmt.Println(doc.Content)
}
