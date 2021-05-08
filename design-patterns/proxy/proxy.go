package proxy

// 接口
type Subject interface {
	Do() string
}

// 类
type RealSubject struct{}

// 实现接口Subject的Do方法
func (RealSubject) Do() string {
	return "real"
}

// 代理类
type Proxy struct {
	real RealSubject
}

// 代理类实现
func (p Proxy) Do() string {
	var res string

	res += "pre:"

	res += p.real.Do()

	res += ":after"

	return res
}
