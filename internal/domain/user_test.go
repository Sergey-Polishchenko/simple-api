package domain

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id   string
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantedName string
	}{
		{
			name: "New user test with trailing spaces",
			args: args{
				id:   "",
				name: "  John  ",
			},
			wantedName: "John",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				if got := NewUser(tt.args.id, tt.args.name); !reflect.DeepEqual(
					got.name,
					tt.wantedName,
				) {
					t.Errorf("NewUser() = %v, want %v", got, tt.wantedName)
				}
			},
		)
	}
}
