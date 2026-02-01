// 策略模式：定义一族算法，封装成可互换的策略，由上下文在运行时选择其一执行。
package strategy

// Strategy 策略接口：具体算法由实现类完成
type Strategy interface {
	Do(a, b int) int
}

// Context 上下文：持有当前策略，委托执行
type Context struct {
	s Strategy
}

func NewContext(s Strategy) *Context {
	return &Context{s: s}
}

func (c *Context) SetStrategy(s Strategy) {
	c.s = s
}

func (c *Context) Execute(a, b int) int {
	return c.s.Do(a, b)
}
