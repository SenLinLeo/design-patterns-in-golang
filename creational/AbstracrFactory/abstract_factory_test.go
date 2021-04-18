package AbstracrFactory

import "testing"

func TestNewSimpleLunxhFactory(t *testing.T) {
	factory := NewSimpleLunchFactory()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()

}
