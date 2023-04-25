package pattern

import "fmt"

/*
	Паттерн состояние позволяет объекту  изменять свое поведение в зависимости от состояния.
*/

// Общий интерфейс для различных состояний
type MobileAlertStater interface {
	Alert() string
}

// использует метод Alert() в зависимости от своего состояния
type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// Изменить состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// Конкретное состояние
type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// Другое конкретное состояние
type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

func TestState() {
	mobile := new(MobileAlert)
	mobile.SetState(&MobileAlertVibration{})
	fmt.Println(mobile.Alert())
	mobile.SetState(&MobileAlertSong{})
	fmt.Println(mobile.Alert())
}
