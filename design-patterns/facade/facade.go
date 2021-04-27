package facade

import "fmt"

func NewApi() Api {
	return &apiImpl{
		a: NewAModuleApi(),
		b: NewBModuleApi(),
	}
}

type Api interface {
	Test() string
}

type apiImpl struct {
	a AModuleApi
	b BModuleApi
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAModuleApi() AModuleApi {
	return &aModuleImpl{}
}

func ModuleApi() AModuleApi {
	return &aModuleImpl{}
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "a module running"
}

type AModuleApi interface {
	TestA() string
}

type BModuleApi interface {
	TestB() string
}

func NewBModuleApi() BModuleApi {
	return &bModuleImpl{}
}

type bModuleImpl struct {
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
