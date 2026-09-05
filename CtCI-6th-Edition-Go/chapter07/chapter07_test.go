package chapter07

import (
	"strings"
	"testing"
)

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}

func TestDeck(t *testing.T) {
	deck := NewDeck()
	if deck.Remaining() != 52 {
		t.Fatal(deck.Remaining())
	}
	hand := deck.Deal(5)
	if len(hand) != 5 || deck.Remaining() != 47 {
		t.Fatal(len(hand), deck.Remaining())
	}
	control := NewDeck()
	seed := int32(1)
	shuffled := NewDeckWithNext(func(n int) int {
		seed = int32(int64(1664525)*int64(seed)+1013904223) | 0
		return int(float64(uint32(seed)) / 4294967296 * float64(n))
	})
	shuffled.Shuffle()
	before := joinCards(control.Deal(5))
	after := joinCards(shuffled.Deal(5))
	if before == after {
		t.Fatal("shuffle should change order")
	}
	mustPanic(t, func() { NewCard("Jokers", 2) })
	mustPanic(t, func() { NewCard("Hearts", 1) })
	mustPanic(t, func() { NewDeck().Deal(53) })
}

func joinCards(cards []Card) string {
	parts := make([]string, len(cards))
	for i, c := range cards {
		parts[i] = c.String()
	}
	return strings.Join(parts, "|")
}

func TestCallCenter(t *testing.T) {
	center := NewCallCenter([]*Employee{NewEmployee("Alice", Respondent), NewEmployee("Bob", Manager)})
	handler := center.Dispatch(NewCall("Jane", Respondent))
	if handler == nil || handler.Name != "Alice" {
		t.Fatal(handler)
	}
	center.Release(handler)

	respondent := NewEmployee("Alice", Respondent)
	respondent.Busy = true
	manager := NewEmployee("Bob", Manager)
	center2 := NewCallCenter([]*Employee{respondent, manager})
	if center2.Dispatch(NewCall("Jane", Respondent)) != manager {
		t.Fatal("escalate")
	}

	center3 := NewCallCenter([]*Employee{NewEmployee("Alice", Respondent)})
	center3.Dispatch(NewCall("A", Respondent))
	if center3.Dispatch(NewCall("B", Respondent)) != nil {
		t.Fatal("no handler")
	}
}

func TestJukeboxParkingReader(t *testing.T) {
	cd := NewCd("Greatest", []*Song{NewSong("A", "Band", 180), NewSong("B", "Band", 200)})
	box := NewJukebox()
	box.AddCd(cd)
	box.SelectCd(0)
	if box.Play().Title != "A" {
		t.Fatal("play")
	}
	if box.NextSong() == nil || box.Play().Title != "B" {
		t.Fatal("next")
	}
	mustPanic(t, func() { NewJukebox().Play() })

	bike := NewBike("Suzuki Access", 1, "MH14AB1234")
	if bike.IsParked() {
		t.Fatal("not parked")
	}
	zone := NewParkZone()
	token := zone.Park(bike.Vehicle)
	if token == nil || !bike.IsParked() {
		t.Fatal("park")
	}
	if zone.Depark(*token) != bike.Vehicle || bike.IsParked() {
		t.Fatal("depark")
	}

	zone2 := NewParkZone()
	car := NewCar("Honda Jazz", 5, "MU268A")
	tok := zone2.Park(car.Vehicle)
	mustPanic(t, func() { zone2.Depark(*tok + 1) })
	if zone2.Depark(*tok) != car.Vehicle {
		t.Fatal("car")
	}

	zone3 := NewParkZone(4)
	if zone3.Park(NewBus("Volvo", 5, "AN657").Vehicle) != nil {
		t.Fatal("no space")
	}
	tok3 := zone3.Park(NewCar("Civic", 2, "X1").Vehicle)
	if tok3 == nil || zone3.SpaceAvailable != 2 {
		t.Fatal(zone3.SpaceAvailable)
	}

	zone4 := NewParkZone()
	b2 := NewBike("Activa", 1, "GI653")
	zone4.Park(b2.Vehicle)
	mustPanic(t, func() { zone4.Park(b2.Vehicle) })

	zone5 := NewParkZone()
	zone5.Park(NewScooter("Honda Activa", 1, "GI653").Vehicle)
	parked := zone5.ListParkedVehicles()
	if len(parked) != 1 || parked[0].Model != "Honda Activa" || parked[0].Size != 1 || parked[0].Number != "GI653" {
		t.Fatal(parked)
	}
	mustPanic(t, func() { NewBike("", 1, "X") })
	mustPanic(t, func() { NewCar("Civic", 0, "X") })

	book := NewBook("Demo", []*Page{NewPage(1, "a"), NewPage(2, "b"), NewPage(3, "c")})
	reader := NewBookReader()
	reader.Open(book)
	if reader.CurrentPage().Content != "a" {
		t.Fatal("page a")
	}
	if reader.NextPage() == nil || reader.CurrentPage().Content != "b" {
		t.Fatal("page b")
	}
	if reader.PrevPage() == nil || reader.Goto(3).Content != "c" {
		t.Fatal("goto")
	}
	mustPanic(t, func() { NewBookReader().CurrentPage() })
}

func TestJigsawChatOthelloCircular(t *testing.T) {
	if !EdgesCompatible(EdgeTab, EdgeBlank) || EdgesCompatible(EdgeTab, EdgeTab) {
		t.Fatal("edges")
	}
	left := NewPiece(1, EdgeFlat, EdgeTab, EdgeFlat, EdgeFlat)
	right := NewPiece(2, EdgeFlat, EdgeFlat, EdgeFlat, EdgeBlank)
	if !CanPlace(left, 0, 0, right, 0, 1) {
		t.Fatal("place")
	}

	server := NewChatServer()
	room := server.CreateRoom("general")
	alice := NewUser("alice")
	bob := NewUser("bob")
	room.Join(alice)
	room.Join(bob)
	room.Send("alice", "hello")
	if len(bob.Inbox) != 1 || bob.Inbox[0] != "[general] alice: hello" || len(alice.Inbox) != 0 || len(room.History()) != 1 {
		t.Fatal(bob.Inbox, alice.Inbox)
	}
	mustPanic(t, func() { NewChatServer().GetRoom("missing") })

	game := NewOthelloGame()
	if !game.ValidMove(2, 3, OthelloBlack) || game.ApplyMove(2, 3, OthelloBlack) != 1 || game.Board[2][3] != OthelloBlack {
		t.Fatal("othello")
	}
	mustPanic(t, func() { NewOthelloGame().ApplyMove(3, 3, OthelloBlack) })

	buf := NewCircularArray[int](3)
	for _, x := range []int{1, 2, 3} {
		buf.Append(x)
	}
	buf.Append(4)
	values := buf.ToList()
	if len(values) != 3 || values[0] != 2 || values[1] != 3 || values[2] != 4 {
		t.Fatal(values)
	}
	mustPanic(t, func() {
		b := NewCircularArray[string](2)
		b.Append("a")
		b.Get(1)
	})
}

func TestMinesweeperFileHash(t *testing.T) {
	game := NewMinesweeperGame(2, 2, 0, nil)
	if game.Reveal(0, 0) != 4 || game.Cell(0, 0) != 0 || game.Cell(1, 1) != 0 {
		t.Fatal("flood")
	}
	mustPanic(t, func() {
		g := NewMinesweeperGame(2, 2, 0, nil)
		g.Reveal(0, 0)
		g.Reveal(0, 0)
	})

	root := NewFileSystemNode("root")
	Mkdir(root, "/home/user")
	WriteFile(root, "/home/user/readme.txt", "hi")
	file := root.Resolve("/home/user/readme.txt")
	if file.IsDir || file.Content != "hi" {
		t.Fatal(file)
	}
	mustPanic(t, func() { NewFileSystemNode("root").Resolve("/nope") })

	m := NewHashMap[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	if m.Get("a") != 1 {
		t.Fatal("get a")
	}
	m.Put("a", 3)
	if m.Get("a") != 3 || !m.Delete("b") || m.Get("b") != 0 || m.Count() != 1 {
		t.Fatal("update delete")
	}
	m2 := NewHashMap[int, int](2)
	for i := range 10 {
		m2.Put(i, i*2)
	}
	for i := range 10 {
		if m2.Get(i) != i*2 {
			t.Fatal(i)
		}
	}
}
