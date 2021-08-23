package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Data int64
	}

	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observer)
		Deregister(Observer)
		Notifier(Event)
	}
)

type (
	eventObserver struct {
		id int
	}

	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

func main() {
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}

/*
어떤 객체의 상태가 변할 때 그와 연관된 객체들에게 알림을 보내는 패턴

1. Subject 가 Observer를 등록
2. Subject 의 변화가 있을 시 update 내용을 Observer 전달
3. if Subject 가 등록을 삭제 하면 update 내용을 받을 수 없음

Subject = eventNotifier
update 내용 = eventObserver 의 id
등록 삭제 = Deregister(l Observer)
*/
