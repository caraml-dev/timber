package values

import (
	"reflect"
	"testing"
)

func TestMerveEnvs(t *testing.T) {
	type args struct {
		left  []Env
		right []Env
	}
	tests := []struct {
		name string
		args args
		want []Env
	}{
		{
			name: "no duplicate",
			args: args{
				left: []Env{
					{
						Name:  "KEY1",
						Value: "VALUE1",
					},
				},
				right: []Env{
					{
						Name:  "KEY2",
						Value: "VALUE2",
					},
				},
			},
			want: []Env{
				{
					Name:  "KEY1",
					Value: "VALUE1",
				},
				{
					Name:  "KEY2",
					Value: "VALUE2",
				},
			},
		},
		{
			name: "with duplicate",
			args: args{
				left: []Env{
					{
						Name:  "KEY1",
						Value: "VALUE1",
					},
				},
				right: []Env{
					{
						Name:  "KEY1",
						Value: "UPDATED_VALUE",
					},
				},
			},
			want: []Env{
				{
					Name:  "KEY1",
					Value: "UPDATED_VALUE",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MerveEnvs(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MerveEnvs() = %v, want %v", got, tt.want)
			}
		})
	}
}
