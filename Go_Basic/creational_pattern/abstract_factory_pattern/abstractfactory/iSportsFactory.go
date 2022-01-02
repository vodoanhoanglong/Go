package abstractfactory

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

func GetSportsFactory(brand string) ISportsFactory {
	switch brand {
	case "adidas":
		return &Adidas{}
	case "nike":
		return &Nike{}
	}
	return nil
}
