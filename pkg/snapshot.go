package pkg

type Snapshot struct {
	state string
}

func (s *Snapshot) GetState() string {
	return s.state
}
