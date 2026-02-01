package factory

// ProductA 具体产品 A
type ProductA struct{}

func (ProductA) Name() string {
	return "ProductA"
}

// ProductB 具体产品 B
type ProductB struct{}

func (ProductB) Name() string {
	return "ProductB"
}
