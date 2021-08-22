package main

import "log"

// 1. Host
type Host interface {
	Accept(Visitor)
}

// 1-1. Player host
type PlayerHost struct {
	Name  string
	Level int
}

func (p PlayerHost) Accept(v Visitor) {
	v.VisitPlayer(p)
}

// 1-2. NPC host
type NPCHost struct {
	Name       string
	IsImmortal bool
}

func (n NPCHost) Accept(v Visitor) {
	v.VisitNPC(n)
}

// 1-3. Object host
type ObjectHost struct {
	Name  string
	Price int
}

func (o ObjectHost) Accept(v Visitor) {
	v.VisitObject(o)
}

// 2. visitor
type Visitor interface {
	VisitPlayer(PlayerHost)
	VisitNPC(NPCHost)
	VisitObject(ObjectHost)
}

type InfoVisitor struct{}

func (InfoVisitor) VisitPlayer(p PlayerHost) {
	log.Printf("Player -> Name:%s ,Level:%d\n", p.Name, p.Level)
}

func (InfoVisitor) VisitNPC(n NPCHost) {
	log.Printf("NPC -> Name:%s ,Immortal:%v\n", n.Name, n.IsImmortal)
}

func (InfoVisitor) VisitObject(o ObjectHost) {
	log.Printf("Object -> Name:%s ,Price:%d\n", o.Name, o.Price)
}

// 3. Behavior
type AggressiveVisitor struct{}

func (AggressiveVisitor) VisitPlayer(p PlayerHost) {
	log.Printf("Attack %s\n", p.Name)
}

func (AggressiveVisitor) VisitNPC(n NPCHost) {
	log.Printf("Attack NPC %s\n", n.Name)
}
func (AggressiveVisitor) VisitObject(o ObjectHost) {
	log.Printf("Unsupported target %s\n", o.Name)
}

func main() {
	infoVst := InfoVisitor{}
	agrVst := AggressiveVisitor{}

	pA := PlayerHost{"icg", 1}
	pB := PlayerHost{"sz", 2}
	npc := NPCHost{"nyn", true}
	obj := ObjectHost{"cake", 19}

	hostList := []Host{pA, npc, obj, pB}

	for _, v := range hostList {
		v.Accept(infoVst)
	}
	println()
	for _, v := range hostList {
		v.Accept(agrVst)
	}
	/*
		output:
		v.Accept(infoVst), 작업 대상 = InfoVisitor{}
		2019/05/04 10:00:49 Player -> Name:icg ,Level:1
		2019/05/04 10:00:49 NPC -> Name:nyn ,Immortal:true
		2019/05/04 10:00:49 Object -> Name:cake ,Price:19
		2019/05/04 10:00:49 Player -> Name:sz ,Level:2

		v.Accept(agrVst), 작업 항목 = AggressiveVisitor{}
		2019/05/04 10:00:49 Attack icg
		2019/05/04 10:00:49 Attack NPC nyn
		2019/05/04 10:00:49 Unsupported target cake
		2019/05/04 10:00:49 Attack sz

		데이터와 알고리즘이 분리되어 데이터의 독립성 UP
		작업 대상자의 경우, v.Accept() 로 인터페이스를 통일

		단점
		작업 대상이 추가될 때마다 작업 주체도 로직을 추가해야함
		작업 대상과 항목의 결합도가 높아짐

		데이터와 알고리즘을 분리해야 되는 경우,
		데이터 구조보다 알고리즘이 더 자주 바뀌는 경우 사용
	*/

}
