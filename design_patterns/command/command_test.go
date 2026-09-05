package command

import "testing"

func TestSimpleRemoteControl(t *testing.T) {
	light := &Light{}

	tests := []struct {
		name     string
		command  Command
		expected string
	}{
		{
			name:     "light on",
			command:  NewLightOnCommand(light),
			expected: "Light is On",
		},
		{
			name:     "light off",
			command:  NewLightOffCommand(light),
			expected: "Light is Off",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			remote := &SimpleRemoteControl{}
			remote.SetCommand(tt.command)

			if got := remote.ButtonWasPressed(); got != tt.expected {
				t.Errorf("ButtonWasPressed() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSimpleRemoteControlWithoutCommand(t *testing.T) {
	remote := &SimpleRemoteControl{}

	if got := remote.ButtonWasPressed(); got != "No command assigned" {
		t.Errorf("ButtonWasPressed() = %q, want %q", got, "No command assigned")
	}
}
