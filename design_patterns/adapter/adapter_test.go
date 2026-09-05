package adapter

import "testing"

func TestTurkeyAdapter(t *testing.T) {
	tests := []struct {
		name      string
		turkey    Turkey
		wantQuack string
		wantFly   string
	}{
		{
			name:      "wild turkey adapts to duck behavior",
			turkey:    &WildTurkey{},
			wantQuack: "グワッグワッ",
			wantFly:   "短い距離しか飛べません x 5回",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duck := NewTurkeyAdapter(tt.turkey)

			if got := duck.Quack(); got != tt.wantQuack {
				t.Errorf("Quack() = %q, want %q", got, tt.wantQuack)
			}
			if got := duck.Fly(); got != tt.wantFly {
				t.Errorf("Fly() = %q, want %q", got, tt.wantFly)
			}
		})
	}
}

func TestTurkeyAdapterImplementsDuck(t *testing.T) {
	var _ Duck = NewTurkeyAdapter(&WildTurkey{})
}
