package composite

type Node interface {
	Getname() string
}

type File struct {
	name string
}

func (f File) GetName() string {
	return f.name
}

type Directory struct {
	name    string
	childen []interface{}
}

func (d Directory) GetName() string {
	return d.name
}

func (d *Directory) Add(file interface{}) {
	d.childen = append(d.childen, file)
}

func main() {
	dir := new(Directory)
	dir.Add(new(File))
	dir.Add(new(Directory))
	secondDir := new(Directory)
	secondDir.Add(dir)
}

/*
reference link : https://jdm.kr/blog/228, https://jdm.kr/blog/228

컴포지트 패턴을 통한 간단한 파일 시스템 구현

장점
- 객체들이 모두 같은 타입으로 취급되기 때문에 새로운 클래스 추가가 용의
- 단일객체, 집합객체 구분하지 않고 코드 작성이 가능

단점
- 설계를 일반화 시켜 객체간의 구분, 제약 힘듬

활용
- 객체 들 간에 계급 및 계층구조가 존재가 있고, 이를 표현해야할 때
- 클라이언트가 단일 객체와 집합 객체를 구분하지 않고 동일한 형태로 사용하고자 할 때
*/
