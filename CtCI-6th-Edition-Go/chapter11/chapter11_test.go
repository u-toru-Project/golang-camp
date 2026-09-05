package chapter11

import "testing"

func TestCountUniqueChars(t *testing.T) {
	if CountUniqueChars("aabBC") != 4 || CountUniqueChars("") != 0 {
		t.Fatal("CountUniqueChars mismatch")
	}
	if CountUniqueCharsBuggy("!!") == CountUniqueChars("!!") {
		t.Fatal("buggy should differ from fixed")
	}
}

func TestCounters(t *testing.T) {
	expected := 8 * 500
	if result := RunUnsafe(8, 500); result > expected {
		t.Fatalf("unsafe %d > %d", result, expected)
	}
	if RunSafe(8, 500) != 4000 {
		t.Fatal("safe counter")
	}
}

func TestChessBoard(t *testing.T) {
	board := ParseBoard("8/8/8/8/8/8/8/k7")
	row, col := KingPosition(board, "k")
	if row != 7 || col != 0 {
		t.Fatalf("king %d,%d", row, col)
	}
	attacked := ParseBoard("8/8/8/8/8/8/r7/k7")
	if !IsSquareAttacked(attacked, 7, 0, false) {
		t.Fatal("rook should attack king")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()
		ParseBoard("8/8")
	}()
}

func TestNormalizeAndShipping(t *testing.T) {
	if NormalizeName(" ada ", " lovelace ") != "Ada Lovelace" {
		t.Fatal(NormalizeName(" ada ", " lovelace "))
	}
	if NormalizeName("ALAN", "turing") != "Alan Turing" {
		t.Fatal(NormalizeName("ALAN", "turing"))
	}
	if ShippingCost(10.0, 100.0) != 25.0 {
		t.Fatal(ShippingCost(10.0, 100.0))
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()
		NormalizeName("", "turing")
	}()
}

func TestPen(t *testing.T) {
	pen := NewPen(0.002)
	if pen.Write("hello") != "he" || !pen.IsEmpty() {
		t.Fatal("ink runout")
	}
	full := NewPen(1.0)
	full.Write("abc")
	if full.WrittenChars() != 3 || !(full.RemainingInk() < 1.0) {
		t.Fatal("pen properties")
	}
}

func TestAtm(t *testing.T) {
	atm := NewAtm(100.0, "1234")
	atm.InsertCard()
	if !atm.EnterPin("1234") || atm.Withdraw(40.0) != 40.0 || atm.Balance() != 60.0 {
		t.Fatal("happy path")
	}
	locked := NewAtm(50.0, "9999")
	locked.InsertCard()
	for range 3 {
		if locked.EnterPin("0000") {
			t.Fatal("wrong pin")
		}
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()
		locked.Withdraw(10.0)
	}()
}
