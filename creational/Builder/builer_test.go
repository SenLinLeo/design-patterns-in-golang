package Builder

import (
	"fmt"
	"testing"
)

func TestConcretebuilder_GetResult(t *testing.T) {
	builder := NewConcretebuilder()
	director := NewDirector(&builder)
	director.Construct()
	result := builder.GetResult()
	fmt.Println(result.Built)
}
