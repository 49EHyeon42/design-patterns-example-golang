package singleton

/*
	전역변수 보호
	1. 예제1, 스레드 사용이 안전하지 않음
	2. 예제2, 스레드 사용이 안전함
	3. 예제3, 가장 좋은 방법
	link : https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f
*/
/*
1.
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {
	if instance == nil {
		instance = make(singleton) // <-- not thread safe
	}
	return instance
}
*/

/*
2.
var lock = &sync.Mutex{}

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = make(singleton) // <-- thread safe
	}
	return instance
}
*/

/*
var once sync.Once

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {
	once.Do(func() { // <-- atomic, does not allow repeating
		instance = make(singleton) // <-- thread safe
	})
	return instance
}
*/

/*
	1. Using  init() functions
	2. Using sysnc.Once
	link : https://medium.com/@ishagirdhar/singleton-pattern-in-golang-9f60d7fdab23
*/

/*
1. Using init() functions
type A struct {
	str string
}
var singleton *A
func init() {
	//initialize static instance on load
	singleton = &A{str :"abc"}
}
//GetInstanceA - get singleton instance pre-initialized
func GetInstanceA() *A {
	return singleton
}
*/

/*
2. Using sysnc.Once
type A struct {
    str string
}
var singleton *A
func GetA() *A {
    sync.once.Do(func() {
        singleton = &A{str: "abc"}
    })
    return singleton
}
func (st *A) GetStr() string {
    return st.str
}
func (st *A) Setstr(s string) {
    st.str = s
}
*/

/*
	1. 방법1, 스레드로부터 안전하지 않음
	2. 방법2, sync.Once로
	link : https://stackoverflow.com/questions/1823286/singleton-in-go
*/

/*
1.
package singleton

type single struct {
        O interface{};
}

var instantiated *single = nil

func New() *single {
        if instantiated == nil {
                instantiated = new(single);
        }
        return instantiated;
}
*/

/*
2.
package singleton

import "sync"

type single struct {
        O interface{};
}

var instantiated *single
var once sync.Once

func New() *single {
        once.Do(func() {
                instantiated = &single{}
        })
        return instantiated
}
*/
