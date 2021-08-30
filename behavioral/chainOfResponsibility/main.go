package main

type GameType int

const (
	TypeFPS GameType = 1
	TypeRPG          = TypeFPS << 1
)

type Game interface {
	Type() GameType
	Start(player string)
}

// chain of responsibility
type GameSelector struct {
	GameList []Game // 사슬 부분
}

func (g *GameSelector) AddGame(games ...Game) {
	g.GameList = append(g.GameList, games...)
}

func (g GameSelector) Start(t GameType, player string) {
	for _, v := range g.GameList {
		if v.Type() == t {
			v.Start(player)
			return
		}
	}
}

type FPSGame struct {
	t GameType
}

func (f FPSGame) Start(player string) {
	println(player, "join in fps game")
}

func (f FPSGame) Type() GameType {
	return f.t
}

type RPGGame struct {
	t GameType
}

func (RPGGame) Start(player string) {
	println(player, "join in rpg game")
}

func (r RPGGame) Type() GameType {
	return r.t
}

func main() {
	fps := FPSGame{TypeFPS}
	rpg := RPGGame{TypeRPG}

	sl := GameSelector{}
	sl.AddGame(fps, rpg)

	player := "icg"
	sl.Start(TypeRPG, player)
	println()
	sl.Start(TypeFPS, player)
	// output:
	/*
		icg join in rpg game
		icg join in fps game
	*/
}

/*
책임 연쇄 패턴
명령 객체와 일련의 처리 객체를 포함하는 디자인 패턴

어떤 요청이 그 요청을 담당하는 객체에 들어오면 각각의 요청에 대해서
특정한 객체가 담당하는 것이 일반적이지만 객체를 연결리스트와
같은 사슬 방식으로 연결한 후에 요청을 수행하지 못하는 객체라면
다음 객체에 넘기며 책임을 넘기는 형태의 패턴을 말한다.

장점
- 결합도를 낮추며, 요청의 발신자와 수신자를 분리시킬 수 있습니다.
- 클라이언트는 처리객체의 집합 내부의 구조를 알 필요가 없습니다.
- 집합 내의 처리 순서를 변경하거나 처리객체를 추가 또는 삭제할 수 있어 유연성이 향상됩니다.
- 새로운 요청에 대한 처리객체 생성이 매우 편리해집니다.

단점
- 충분한 디버깅을 거치지 않았을 경우 집합 내부에서 사이클이 발생할 수 있습니다.
- 디버깅 및 테스트가 쉽지 않습니다.

활용
- 요청의 발신자와 수신자를 분리하는 경우
- 요청을 처리할 수 있는 객체가 여러개일 때 그 중 하나에 요청을 보내려는 경우
- 코드에서 처리객체(handler)를 명시적으로 지정하고 싶지 않은 경우
즉, 책임 연쇄 패턴은 요청을 처리할 수 있는 객체가 여러 개이고
처리객체가 특정적이지 않을 경우 권장되는 패턴입니다.



*/
