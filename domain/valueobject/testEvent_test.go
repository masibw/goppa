package valueobject

import "testing"

func TestTestEvent_IsSlowerThan(t1 *testing.T) {
	type fields struct {
		Name    string
		Elapsed float64
	}
	type args struct {
		prev   float64
		border float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Return true elapsed is larger than prev * border.",
			fields: fields{
				Name:    "",
				Elapsed: 2,
			},
			args: args{
				prev:   1.2,
				border: 1.5,
			},
			want: true,
		},
		{
			name: "Return false elapsed is smaller than prev * border.",
			fields: fields{
				Name:    "",
				Elapsed: 1.5,
			},
			args: args{
				prev:   1.2,
				border: 1.5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TestEvent{
				Name:    tt.fields.Name,
				Elapsed: tt.fields.Elapsed,
			}
			if got := t.IsSlowerThan(tt.args.prev, tt.args.border); got != tt.want {
				t1.Errorf("IsSlowerThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
