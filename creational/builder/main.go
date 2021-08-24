package main

import "fmt"

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

func main() {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		fmt.Println("Err")
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()
	if bike.Wheels != 2 {
		fmt.Println("Err")
	}
}

/*
생성과 표현의 분리
상태값을 동결시키는 패턴

복잡한 객체를 생성할 때,
서로 다른 표현방식을 가지는 객체를 동일한 방식으로 생성하고 싶은 경우


장점
- 가독성이 좋음
- 생성에 필요한 파라미터가 추가될 때 마다, 생성자 오버로딩 안해도됨

단점
- 코드가 길어지기 때문에 매개변수가 4개 이상은 될 때 사용
- 빌더 클래스 구현

비교
- 생성자
- 자바빈즈 패턴

*/
