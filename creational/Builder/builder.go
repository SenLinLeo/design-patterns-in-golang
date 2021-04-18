package Builder

type Builder interface {
	Build()
}

type Director struct {
	builder Builder
}

func NewBuilder(b Builder) Director {
	return Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.Build()
}

func NewDirector(b Builder) Director{
	return Director{builder:b}
}

type Concretebuilder struct {
	built bool
}

func NewConcretebuilder() Concretebuilder {
	return Concretebuilder{false}
}

func (b *Concretebuilder) Build() {
	b.built=true
}
type Product struct {
	Built bool
}

func (b *Concretebuilder) GetResult() Product {
	return Product{b.built}
}