package pkg

type Creator struct {
	State string
}

func (c *Creator) Create() *Snapshot {
	return &Snapshot{State: c.State}
}

func (c *Creator) Restore(s *Snapshot) {
	c.State = s.GetState()
}

func (c *Creator) SetState(state string) {
	c.State = state
}

func (c *Creator) GetState() string {
	return c.State
}
