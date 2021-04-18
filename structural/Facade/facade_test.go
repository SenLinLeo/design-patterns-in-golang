package Facade

import "testing"

/**
外观模式(门面模式)
**/
func TestCarFacade_CreateCompleteCar(t *testing.T) {
	f := CarFacade{}
	f.CreateCompleteCar()
}
