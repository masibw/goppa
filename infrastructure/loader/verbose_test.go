package loader

import (
	"github.com/masibw/goppa/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerboseLoader_Load(t *testing.T) {
	testdataDirPath := "../../"
	type args struct {
		fileName string
	}
	tests := []struct {
		name         string
		args         args
		wantTestData []valueobject.TestEvent
		wantErr      bool
	}{
		{
			name: "Can load test data from file created with go test -v option.",
			args: args{
				fileName: "testdata/prev.txt",
			},
			wantTestData: []valueobject.TestEvent{
				{
					Name:    "TestAdd",
					Elapsed: 0,
				},
				{
					Name:    "TestAdd/Can_add_up_two_numbers.",
					Elapsed: 0,
				},
				{
					Name:    "TestAdd/Can_add_up_two_numbers(includes_negative_value).",
					Elapsed: 0,
				},
				{
					Name:    "TestMinus",
					Elapsed: 0,
				},
				{
					Name:    "TestMinus/Can_subtract_b_from_a.",
					Elapsed: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &VerboseLoader{}
			gotTestData, gotErr := l.Load(testdataDirPath + tt.args.fileName)

			assert.Equal(t, tt.wantTestData, gotTestData)
			assert.Equal(t, tt.wantErr, gotErr != nil)
		})
	}
}
