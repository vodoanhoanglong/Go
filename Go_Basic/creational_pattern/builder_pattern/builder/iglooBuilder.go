package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow door"
}

func (b *iglooBuilder) setFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
