package main

import "log"

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func Double(n int) int {
	return n * 2
}

func main() {
	f := LogDecorate(Double)
	f(5)
}

// Starting execution with the integer 5
// Execution is completed with the result 10

/*
데코레이터 패턴
- 객체의 결합을 통해 기능을 동적으로 유연하게 확장

장점
- 객체에 동적으로 기능 추가가 간단하게 가능

단점
- 작은 데코레이터 클래스들이 계속 추가되어 클래스가 많아짐
- 객체의 정체를 알기 힘들고 복잡함

활용
- 객체가 상황에 따라 다양한 기능이 추가되거나 삭제 되어야 할 때
*/
