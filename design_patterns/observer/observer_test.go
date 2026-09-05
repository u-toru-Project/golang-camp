package observer

import "testing"

func TestObserverNotification(t *testing.T) {
	item := NewItem("PlayStation 5")
	customer1 := &Customer{id: "TestUser1"}
	customer2 := &Customer{id: "TestUser2"}
	logger := &SystemLogger{}

	item.Register(customer1)
	item.Register(customer2)
	item.Register(logger)

	item.UpdateAvailability()

	if len(customer1.ReceivedMsgs) != 1 {
		t.Fatalf("customer1 should have exactly 1 message, got %d", len(customer1.ReceivedMsgs))
	}

	expectedMsg1 := "お客様 TestUser1: PlayStation 5 が入荷しましたよ!\n"
	if customer1.ReceivedMsgs[0] != expectedMsg1 {
		t.Fatalf("customer1 message = %q, want %q", customer1.ReceivedMsgs[0], expectedMsg1)
	}

	if len(customer2.ReceivedMsgs) != 1 {
		t.Fatalf("customer2 should have exactly 1 message, got %d", len(customer2.ReceivedMsgs))
	}

	if len(logger.LogData) != 1 {
		t.Fatalf("logger should have exactly 1 log entry, got %d", len(logger.LogData))
	}

	expectedLog := "[Log] システム記録: アイテム「PlayStation 5」の入荷イベントを検知しました。\n"
	if logger.LogData[0] != expectedLog {
		t.Fatalf("logger log = %q, want %q", logger.LogData[0], expectedLog)
	}
}

func TestItemDeregister(t *testing.T) {
	item := NewItem("Nintendo Switch")
	customer := &Customer{id: "Alice"}

	item.Register(customer)
	item.Deregister(customer)
	item.UpdateAvailability()

	if len(customer.ReceivedMsgs) != 0 {
		t.Fatalf("deregistered customer should have 0 messages, got %d", len(customer.ReceivedMsgs))
	}
}
