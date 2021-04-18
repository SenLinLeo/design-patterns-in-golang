package Facade

import "fmt"

/**
外观模式(门面模式)
**/
type CarModel struct{}

func NewCarModel() *CarModel {
	return &CarModel{}
}
func (c *CarModel) SetModel() {
	fmt.Println("CarModel - SetModel")
}

type Carengine struct {}

func (c *Carengine) SetEngine() {
	fmt.Println("CarEngine - SetEngine")
}

type CarBody struct {

}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetCarBody() {
	fmt.Println("Carbody - Set Body")
}

type CarFacade struct {
	model CarModel
	engine Carengine
	body CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model : CarModel{},
		engine:Carengine{},
		body  :CarBody{},
	}
}


func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetCarBody()
}