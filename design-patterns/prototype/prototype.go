package prototype

type Cloneable interface {
	Clone() Cloneable
}

type Manager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *Manager {
	return &Manager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *Manager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *Manager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
