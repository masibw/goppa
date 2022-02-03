package extract

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameOfTest(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
	}{
		{
			name: "Can extract test name",
			args: args{
				line: "--- PASS: TestAdd (0.00s)",
			},
			wantName: "TestAdd",
		}, {
			name: "Can extract test name even if there are spaces in front of them.",
			args: args{
				line: "    --- PASS: TestAdd/Can_add_up_two_numbers. (0.00s)",
			},
			wantName: "TestAdd/Can_add_up_two_numbers.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantName, NameOfTest(tt.args.line))
		})
	}
}

func TestElapsed(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
	}{
		{
			name: "Can extract elapsed time.",
			args: args{
				line: "--- PASS: TestAdd (0.00s)",
			},
			wantValue: "0.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue := Elapsed(tt.args.line)
			assert.Equal(t, tt.wantValue, gotValue)
		})
	}
}
