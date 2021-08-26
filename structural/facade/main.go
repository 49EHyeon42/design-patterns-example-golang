package facade

import (
	"errors"
	"fmt"
)

type Shop interface {
	Sell() error
}

type EggShop struct{}

func (EggShop) Sell() error {
	return errors.New("no more eggs left")
}

type MilkShop struct{}

func (MilkShop) Sell() error {
	return errors.New("no more milk left")
}

type WheatFlourShop struct{}

func (WheatFlourShop) Sell() error {
	return errors.New("no more wheat flour left")
}

type DealerFacade struct {
	EgShop Shop
	MkShop Shop
	WfShop Shop
}

func (d DealerFacade) BuyAll() {
	e1 := d.EgShop.Sell()
	e2 := d.MkShop.Sell()
	e3 := d.WfShop.Sell()
	if e1 == nil && e2 == nil && e3 == nil {
		//success
	} else {
		//fail and rollback
		fmt.Printf("error:\n%v\n%v\n%v", e1, e2, e3)
	}
}

func main() {
	dealer := DealerFacade{
		EggShop{},
		MilkShop{},
		WheatFlourShop{},
	}
	dealer.BuyAll()
	/*
		output:
		error:
		no more eggs left
		no more milk left
		no more wheat flour left
	*/
}

/*
어떤 소프트웨어의 다른 커라란 코드 부분에 대하여 간략화된
인터페이스를 제공해주는 디자인 패턴

구조
퍼시드 = sell
클라이언트 = main
패키지 = shop

장점
쉬운 이해 및 사용
가독성 증가
의존성 감소, 유현성 증가
*/
