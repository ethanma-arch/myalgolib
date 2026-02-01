package strategy

// AddStrategy 加法策略
type AddStrategy struct{}

func (AddStrategy) Do(a, b int) int {
	return a + b
}

// SubStrategy 减法策略
type SubStrategy struct{}

func (SubStrategy) Do(a, b int) int {
	return a - b
}

// MulStrategy 乘法策略
type MulStrategy struct{}

func (MulStrategy) Do(a, b int) int {
	return a * b
}
