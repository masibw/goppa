package usecase

import (
	"github.com/masibw/goppa/infrastructure/loader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareWithPrev(t *testing.T) {
	testdataDirPath := "../"
	type args struct {
		prevFileName    string
		currentFileName string
		border          float64
	}
	tests := []struct {
		name     string
		args     args
		wantDiff []string
	}{
		{
			name: "Can can return the one with the problem.",
			args: args{
				prevFileName:    "testdata/prev.txt",
				currentFileName: "testdata/slower.txt",
				border:          1.5,
			},
			wantDiff: []string{
				"'TestAdd', prev: 0s, current: 2s",
				"'TestAdd/Can_add_up_two_numbers.', prev: 0s, current: 1s",
				"'TestAdd/Can_add_up_two_numbers(includes_negative_value).', prev: 0s, current: 1s",
			},
		},
		{
			name: "If there is no problems, return nil.",
			args: args{
				prevFileName:    "testdata/prev.txt",
				currentFileName: "testdata/not-slower.txt",
				border:          1.5,
			},
			wantDiff: nil,
		},
		{
			name: "If it is less than border times, return nil.",
			args: args{
				prevFileName:    "testdata/prev.txt",
				currentFileName: "testdata/slower.txt",
				border:          1000.0,
			},
			wantDiff: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := loader.NewVerboseLoader()
			assert.Equal(t, tt.wantDiff, CompareWithPrev(testdataDirPath+tt.args.prevFileName, testdataDirPath+tt.args.currentFileName, l, tt.args.border))
		})
	}
}
