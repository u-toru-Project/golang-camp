package iterator

import (
	"bytes"
	"strings"
	"testing"
)

func TestMenuItemIteratorNext(t *testing.T) {
	tests := []struct {
		name         string
		items        []*MenuItem
		wantNames    []string
		wantFinalNil bool
	}{
		{
			name: "returns items in order",
			items: []*MenuItem{
				newMenuItem("first", "desc", true, 1.0),
				newMenuItem("second", "desc", false, 2.0),
			},
			wantNames:    []string{"first", "second"},
			wantFinalNil: true,
		},
		{
			name:         "returns false for empty iterator",
			items:        nil,
			wantNames:    nil,
			wantFinalNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := newMenuItemIterator(tt.items)

			for _, wantName := range tt.wantNames {
				item, ok := iterator.Next()
				if !ok {
					t.Fatalf("Next() ok = false, want true")
				}
				if item == nil {
					t.Fatalf("Next() item = nil, want menu item")
				}
				if item.Name != wantName {
					t.Fatalf("Next() item.Name = %q, want %q", item.Name, wantName)
				}
			}

			item, ok := iterator.Next()
			if ok {
				t.Fatalf("final Next() ok = true, want false")
			}
			if (item == nil) != tt.wantFinalNil {
				t.Fatalf("final Next() item nil = %t, want %t", item == nil, tt.wantFinalNil)
			}
		})
	}
}

func TestDinerMenuCreateIterator(t *testing.T) {
	menu := NewDinerMenu()
	iterator := menu.CreateIterator()

	var got []string
	for {
		item, ok := iterator.Next()
		if !ok {
			break
		}
		got = append(got, item.Name)
	}

	want := []string{"Vegetarian BLT", "BLT"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("iterator items = %v, want %v", got, want)
	}
}

func TestWaitressPrintMenu(t *testing.T) {
	var out bytes.Buffer
	waitress := NewWaitress(NewPancakeHouseMenu(), NewDinerMenu(), &out)
	waitress.PrintMenu()

	checks := []string{
		"MENU",
		"BREAKFAST",
		"K&B's Pancake Breakfast, $2.99 -- Pancakes with scrambled eggs and toast",
		"LUNCH",
		"BLT, $2.99 -- Bacon with lettuce & tomato on whole wheat",
	}

	output := out.String()
	for _, check := range checks {
		if !strings.Contains(output, check) {
			t.Fatalf("PrintMenu() output missing %q\noutput:\n%s", check, output)
		}
	}
}

func TestNewWaitressPanicsWhenWriterIsMissing(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic when writer is missing")
		}
	}()

	NewWaitress(NewPancakeHouseMenu(), NewDinerMenu(), nil)
}
