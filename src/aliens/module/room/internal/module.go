package internal

type Module struct {
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {

}

func (m *Module) OnDestroy() {

}

func (s *Module) Run(closeSig chan bool) {

}
