package chapter06

import (
	"math"
	"math/rand"
	"reflect"
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

func TestHeavyPill(t *testing.T) {
	if MinWeighingsBalance(8) != 2 || MinWeighingsBalance(1) != 0 {
		t.Fatal("min weighings")
	}
	if FindHeavyPillIndex([]float64{1.1, 1, 1, 1, 1, 1, 1, 1}) != 0 {
		t.Fatal("index 0")
	}
	if FindHeavyPillIndex([]float64{1, 1, 1, 1, 1, 1, 1, 1.1}) != 7 {
		t.Fatal("index 7")
	}
	if FindHeavyPillIndex([]float64{1, 1, 1, 1, 1.1, 1, 1, 1}) != 4 {
		t.Fatal("index 4")
	}
	mustPanic(t, func() { FindHeavyPillIndex([]float64{1, 1, 1}) })
	mustPanic(t, func() { MinWeighingsBalance(0) })
}

func TestBasketball(t *testing.T) {
	if ChooseGame(0.5) != "tie" || ChooseGame(0.9) != "game2" || ChooseGame(0.1) != "game1" {
		t.Fatal("choose game")
	}
	if ChooseGame(0) != "tie" || ChooseGame(1) != "tie" {
		t.Fatal("tie extremes")
	}
	if ProbAheadAfterShots(1, 5, 3) != 1 || ProbAheadAfterShots(0, 5, 1) != 0 {
		t.Fatal("ahead")
	}
	mustPanic(t, func() { ProbAheadAfterShots(1.5, 3, 1) })
	if ProbWinRaceToN(1, 3) != 1 || ProbWinRaceToN(0, 3) != 0 {
		t.Fatal("race")
	}
	mustPanic(t, func() { ProbWinRaceToN(0.5, 0) })
}

func TestDominosAntsJugs(t *testing.T) {
	if DominoTilings2ByN(0) != 1 || DominoTilings2ByN(1) != 1 || DominoTilings2ByN(2) != 2 {
		t.Fatal("2xn small")
	}
	if DominoTilings2ByN(3) != 3 || DominoTilings2ByN(4) != 5 || DominoTilings2ByN(8) != 34 {
		t.Fatal("2xn")
	}
	if DominoTilings3ByN(8) != 24 || DominoTilings3ByN(0) != 1 {
		t.Fatal("3xn")
	}
	mustPanic(t, func() { DominoTilings2ByN(-1) })
	for n := 3; n < 12; n++ {
		if math.Abs(CollisionProbability(n)-CollisionProbabilityClosedForm(n)) >= 1e-12 {
			t.Fatalf("ants n=%d", n)
		}
	}
	if math.Abs(CollisionProbability(3)-0.75) >= 1e-12 {
		t.Fatal("triangle")
	}
	mustPanic(t, func() { CollisionProbability(1) })
	if !CanMeasure(3, 5, 4) || CanMeasure(2, 4, 3) || !CanMeasure(3, 5, 0) {
		t.Fatal("can measure")
	}
	if MeasureFourLiters() == nil || *MeasureFourLiters() <= 0 {
		t.Fatal("four liters")
	}
	if MinPoursBfs(2, 4, 3) != nil || *MinPoursBfs(3, 5, 0) != 0 {
		t.Fatal("min pours")
	}
	amounts := MeasureableAmounts(3, 5)
	if _, ok := amounts[1]; !ok {
		t.Fatal("1")
	}
	if _, ok := amounts[4]; !ok {
		t.Fatal("4")
	}
	if _, ok := amounts[5]; !ok {
		t.Fatal("5")
	}
	if _, ok := amounts[0]; ok {
		t.Fatal("0")
	}
	mustPanic(t, func() { MeasureableAmounts(0, 5) })
}

func TestBlueEyedApocalypseEggLockersPoison(t *testing.T) {
	if DaysUntilBlueEyedLeave(0, true) != 0 || DaysUntilBlueEyedLeave(1, true) != 1 {
		t.Fatal("days")
	}
	if DaysUntilBlueEyedLeave(5, true) != 5 || DaysUntilBlueEyedLeave(100, true) != 100 {
		t.Fatal("days more")
	}
	if DaysUntilBlueEyedLeave(10, false) != 0 || SimulateLeavingDay([]string{"B", "R", "B", "R"}) != 2 {
		t.Fatal("simulate leave")
	}
	ratio := SimulateApocalypse(400, rand.New(rand.NewSource(0)))
	if math.Abs(0.5-ratio) >= 0.15 {
		t.Fatalf("ratio %v", ratio)
	}
	mustPanic(t, func() { SimulateApocalypse(0, nil) })
	mustPanic(t, func() { SimulateApocalypse(-1, nil) })
	mustPanic(t, func() { SimulateApocalypse(1_000_001, nil) })
	if ReadPositiveInt("prompt", 3, 100_000, func(string) string { return "100" }) != 100 {
		t.Fatal("read")
	}
	mustPanic(t, func() { ReadPositiveInt("prompt", 1, 100_000, func(string) string { return "0" }) })
	if MinDrops(0, 2) != 0 || MinDrops(1, 1) != 1 || MinDrops(2, 2) != 2 {
		t.Fatal("egg small")
	}
	if MinDrops(10, 2) != 4 || MinDrops(100, 2) != 14 || MinDropsBinarySearch(100, 2) != 14 {
		t.Fatal("egg")
	}
	if MinDrops(50, 1) != 50 {
		t.Fatal("one egg")
	}
	mustPanic(t, func() { MinDrops(10, 0) })
	want100 := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	if !reflect.DeepEqual(OpenLockersAfterPasses(100), want100) || !reflect.DeepEqual(OpenLockersPerfectSquares(100), want100) {
		t.Fatal("lockers 100")
	}
	if !reflect.DeepEqual(OpenLockersAfterPasses(10), []int{1, 4, 9}) || len(OpenLockersAfterPasses(0)) != 0 {
		t.Fatal("lockers")
	}
	for _, poisoned := range []int{0, 1, 42, 999} {
		world := NewWorld(10, 1000, poisoned)
		if FindPoison(world) != poisoned || world.Day() != DaysForResult {
			t.Fatalf("poison %d", poisoned)
		}
	}
	mustPanic(t, func() { NewWorld(3, 1000, 0) })
	mustPanic(t, func() { NewWorld(10, 10, 99) })
}
