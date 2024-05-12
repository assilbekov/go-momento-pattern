package pkg

type Snapshot struct {
	State string
}

func (s *Snapshot) GetState() string {
	return s.State
}
