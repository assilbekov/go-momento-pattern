package pkg

type Guardian struct {
	Items []*Snapshot
}

func (g *Guardian) Add(s *Snapshot) {
	g.Items = append(g.Items, s)
}

func (g *Guardian) Get(i int) *Snapshot {
	return g.Items[i]
}
