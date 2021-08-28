package main

type Cake struct {
	Flavour string
}

type CakeFactory struct {
	Cakes map[string]Cake
}

func (f CakeFactory) NewCake(flavour string) Cake {
	if c, ok := f.Cakes[flavour]; ok {
		println("get an existing", c.Flavour, "cake from map")
	} else {
		f.Cakes[flavour] = Cake{flavour}
		println("put a new", flavour, "cake into map")
	}
	return f.Cakes[flavour]
}

func main() {
	factory := CakeFactory{make(map[string]Cake)}

	factory.NewCake("strawberry")
	factory.NewCake("chocolates")
	factory.NewCake("nynicg")

	factory.NewCake("strawberry")
	factory.NewCake("nynicg")
	factory.NewCake("chocolates")
	/*
		output:
		put a new strawberry cake into map
		put a new chocolates cake into map
		put a new nynicg cake into map
		get an existing strawberry cake from map
		get an existing nynicg cake from map
		get an existing chocolates cake from map
	*/
}

/*
플라이웨이트 패턴
인스턴스 한 개만 가지고 여러 개의 가상 인스턴스를 제공하고 싶을 때 사용하는 패턴
인스턴스를 가능한 공유시켜 새로운 클래스 생성을 통한 메모리 낭비를 줄임

장점
- 메모리 절약
- 인스턴스 집중

단점
- 특정 인스턴스만 다른 인스턴스 처럼 동작하도록 하는 것이 불가능
- 객체 값이 변경되면 가상 객체들도 영향을 받음
*/
