package main

import (
	"fmt"
	"momento/pkg"
)

func main() {
	storage := &pkg.Guardian{
		Items: make([]*pkg.Snapshot, 0),
	}
	creator := &pkg.Creator{
		State: "state1",
	}
	fmt.Println("Creator state:", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("state2")
	fmt.Println("Creator state:", creator.GetState())
	storage.Add(creator.Create())

	creator.Restore(storage.Get(0))
	fmt.Println("Creator state:", creator.GetState())
	creator.Restore(storage.Get(1))
	fmt.Println("Creator state:", creator.GetState())
}
