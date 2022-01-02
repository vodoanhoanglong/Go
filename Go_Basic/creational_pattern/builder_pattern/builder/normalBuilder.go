package builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden door"
}

func (b *normalBuilder) setFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
