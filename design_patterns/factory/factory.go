// 工厂模式：由工厂根据参数创建具体产品，调用方只依赖产品接口，不依赖具体类型。
package factory

// Product 产品接口
type Product interface {
	Name() string
}

// NewProduct 工厂函数：根据类型名返回对应产品，调用方无需知道具体类型
func NewProduct(typ string) Product {
	switch typ {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		return nil
	}
}
