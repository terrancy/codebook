package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	n := &Navigator{}
	n.SetStrategy(WalkStrategy{})
	if got := n.Go("A", "B"); got != "walk: A -> B" {
		t.Fatalf("walk got %q", got)
	}
	n.SetStrategy(DriveStrategy{})
	if got := n.Go("A", "B"); got != "drive: A -> B" {
		t.Fatalf("drive got %q", got)
	}
}

func TestObserver(t *testing.T) {
	pub := &NewsPublisher{}
	r1 := &Reader{Name: "alice"}
	r2 := &Reader{Name: "bob"}
	pub.Attach(r1)
	pub.Attach(r2)
	pub.Notify("breaking news")

	if r1.Got != "breaking news" || r2.Got != "breaking news" {
		t.Fatalf("observer not notified: %+v %+v", r1, r2)
	}
	pub.Detach(r1)
	pub.Notify("second")
	if r1.Got != "breaking news" {
		t.Fatalf("detached observer should not receive: %+v", r1)
	}
	if r2.Got != "second" {
		t.Fatalf("observer2 should receive second: %+v", r2)
	}
}

func TestTemplateMethod(t *testing.T) {
	b := &Beverage{}
	b.SetSteps(Tea{})
	if got := b.Make(); got != "boil water -> add tea leaf -> steep 3 min" {
		t.Fatalf("tea template got %q", got)
	}
	b.SetSteps(CoffeeRecipe{})
	if got := b.Make(); got != "boil water -> add coffee powder -> drip brew" {
		t.Fatalf("coffee template got %q", got)
	}
}

func TestChainOfResponsibility(t *testing.T) {
	auth := &AuthHandler{}
	log := &LogHandler{}
	auth.SetNext(log)

	if got := auth.Handle("unauthorized"); got != "AuthHandler: rejected (no auth)" {
		t.Fatalf("rejected case got %q", got)
	}
	if got := auth.Handle("ok"); got != "LogHandler: done" {
		t.Fatalf("passed case got %q", got)
	}
}

func TestCommand(t *testing.T) {
	r := &Remote{}
	r.SetCommand(TurnOnCommand{light: Light{}})
	if got := r.Press(); got != "light on" {
		t.Fatalf("execute got %q", got)
	}
	if got := r.PressUndo(); got != "light off" {
		t.Fatalf("undo got %q", got)
	}
}

func TestIterator(t *testing.T) {
	c := NewNameCollection([]string{"a", "b", "c"})
	it := c.Iterator()
	var got []string
	for it.HasNext() {
		got = append(got, it.Next())
	}
	if len(got) != 3 || got[0] != "c" || got[2] != "a" {
		t.Fatalf("reverse iterator wrong: %v", got)
	}
}

func TestMediator(t *testing.T) {
	room := NewChatRoom()
	u1 := NewUser("alice", room)
	u2 := NewUser("bob", room)
	u1.Send("hello") // 经中介者转发，bob 收到，alice 自己不收
	if len(u2.Inbox()) != 1 || u2.Inbox()[0] != "hello" {
		t.Fatalf("bob inbox wrong: %v", u2.Inbox())
	}
	if len(u1.Inbox()) != 0 {
		t.Fatalf("alice should not receive own msg: %v", u1.Inbox())
	}
}

func TestMemento(t *testing.T) {
	o := &Originator{}
	caretaker := &Caretaker{}
	o.SetState("v1")
	caretaker.Push(o.Save())
	o.SetState("v2")
	if o.State() != "v2" {
		t.Fatalf("should be v2, got %s", o.State())
	}
	o.Restore(caretaker.Pop())
	if o.State() != "v1" {
		t.Fatalf("should restore to v1, got %s", o.State())
	}
}

func TestState(t *testing.T) {
	ctx := NewContext(PendingState{})
	if got := ctx.Request(); got != "pending -> paid" {
		t.Fatalf("pending got %q", got)
	}
	if got := ctx.Request(); got != "paid -> done" {
		t.Fatalf("paid got %q", got)
	}
	if got := ctx.Request(); got != "already done" {
		t.Fatalf("done got %q", got)
	}
}

func TestVisitor(t *testing.T) {
	els := []Element{&FileElement{Name: "a"}, &DirElement{Name: "d"}}
	sv := SizeVisitor{}
	got := []string{els[0].Accept(sv), els[1].Accept(sv)}
	if got[0] != "a:1" || got[1] != "d:5" {
		t.Fatalf("size visitor wrong: %v", got)
	}
	pv := PathVisitor{}
	got2 := []string{els[0].Accept(pv), els[1].Accept(pv)}
	if got2[0] != "/files/a" || got2[1] != "/dirs/d" {
		t.Fatalf("path visitor wrong: %v", got2)
	}
}

func TestInterpreter(t *testing.T) {
	expr := NewAnd(NewTerminal(true), NewOr(NewTerminal(false), NewTerminal(true)))
	if !expr.Interpret() {
		t.Fatal("expected true: true AND (false OR true)")
	}
	expr2 := NewAnd(NewTerminal(true), NewTerminal(false))
	if expr2.Interpret() {
		t.Fatal("expected false: true AND false")
	}
}
