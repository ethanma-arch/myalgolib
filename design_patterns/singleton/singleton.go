// 单例模式：全局唯一实例，懒加载，线程安全（sync.Once）
package singleton

import "sync"

type Singleton struct {
	value string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance 返回单例，首次调用时初始化，之后始终返回同一实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{value: "initialized"}
	})
	return instance
}

// Value 仅作示例，可替换成实际业务字段
func (s *Singleton) Value() string {
	return s.value
}
