package AbstracrFactory

import "fmt"

/**
抽象工厂模式
**/
type Lunch interface {
	Cook()
}

type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("It is Rise...")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("It is Tomato....")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpeLunchFactory struct {
}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpeLunchFactory{}
}

func (s *SimpeLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpeLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}