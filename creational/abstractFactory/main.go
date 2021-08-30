package main

import "fmt"

// Computer
type ComputerFactory struct{}

func NewComputerFactory(brand string) {
	NewKeyboradFactory().CreateKeyborad(brand).NewKeyborad()
	NewMouseFactory().CreateMouse(brand).NewMouse()
}

// Keyborad
type KeyboradFactory struct{}

func NewKeyboradFactory() *KeyboradFactory {
	return &KeyboradFactory{}
}

func (kf *KeyboradFactory) CreateKeyborad(keyborad string) keyborad {
	switch keyborad {
	case "LG":
		return new(LGKeyborad)
	case "Samsung":
		return new(SamsungKeyborad)
	default:
		return nil
	}
}

type keyborad interface {
	NewKeyborad()
}

type LGKeyborad struct{}

func (k LGKeyborad) NewKeyborad() {
	fmt.Println("LG 키보드 생성")
}

type SamsungKeyborad struct{}

func (k SamsungKeyborad) NewKeyborad() {
	fmt.Println("삼성 키보드 생성")
}

//Mouse
type MouseFactory struct{}

func NewMouseFactory() *MouseFactory {
	return &MouseFactory{}
}

func (mf *MouseFactory) CreateMouse(mouse string) Mouse {
	switch mouse {
	case "LG":
		return new(LGMouse)
	case "Samsung":
		return new(SamsungMouse)
	default:
		return nil
	}
}

type Mouse interface {
	NewMouse()
}

type LGMouse struct{}

func (m LGMouse) NewMouse() {
	fmt.Println("LG 마우스 생성")
}

type SamsungMouse struct{}

func (m SamsungMouse) NewMouse() {
	fmt.Println("삼성 마우스 생성")
}

func main() {
	fmt.Println("Create LG computer")
	NewComputerFactory("LG")
	fmt.Println("Create Samsung computer")
	NewComputerFactory("Samsung")
}

/*
reference link : https://victorydntmd.tistory.com/300
https://devowen.com/326

추상 팩토리 패턴
- 서로 관련이 있는 객체들을 통째로 묶어서 팩토리 클래스로 만들고,
  이들 팩토리를 조건에 따라 생성하도록 다시 팩토리를 만들어서 객체를 생성하는 패턴
- 팩토리 메서드 패턴을 좀 더 캡슐화한 방식

장점
- 구체적인 클래스 분리
- 제품군을 쉽게 대체할 수 있음

단점
- 새로운 종류의 제품을 제공하기 어려움

구현 방법 여러가지
- 팩토리 단일체로 정의
- 제품을 생성
- 확장 가능한 팩토리들을 정의

활용
- 객체가 생성되거나 구성, 표현되는 방식과 무관하게 시스템을 독립적으로 만들고자 할 때
- 여러 제품군 중 하나를 선택해서 시스템을 설정해야 하고 한 번 구성한 제품을 다른 것으로 대체할 수 있을 때
- 관련된 제품 객체들이 함께 사용되도록 설계되었고, 이 부분에 대한 제약이 외부에도 지켜지도록 하고 싶을 때
- 제품에 대한 클래스 라이브러리를 제공하고, 그들의 구현이 아닌 인터페이스를 노출 시키고 싶을 때
*/
