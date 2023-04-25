package pattern

import "fmt"

type Builder interface {
	build() *House // factory method
}
type WoodBuilder struct {
}

func (b WoodBuilder) build() House {
	return new(WoodHouse)
}

type PanelBuilder struct {
}

func (b PanelBuilder) build() House {
	return new(PanelHouse)
}

type House interface {
	show()
}
type WoodHouse struct {
}

func (h WoodHouse) show() {
	fmt.Println("wood house")
}

type PanelHouse struct {
}

func (h PanelHouse) show() {
	fmt.Println("panel house")
}

func useFactoryMethod() {
	b1 := new(WoodBuilder)
	b1.build().show()
	b2 := new(PanelBuilder)
	b2.build().show()
}
