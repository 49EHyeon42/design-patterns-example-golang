package factorymethod

func GetShape(shapeType string) Shape {
	if shapeType == "circle" {
		return &Circle{}
	} else if shapeType == "rectangle" {
		return &Rectangle{}
	} else if shapeType == "square" {
		return &Square{}
	} else {
		return nil
	}
}

// https://niceman.tistory.com/143
