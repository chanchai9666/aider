package aider

import (
	"testing"
)

func TestRound2d(t *testing.T) {
	type args struct {
		value float64
		d     int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "round2d",
			args: args{value: 1234.455555556, d: 2},
			want: 1234.45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round2d(tt.args.value, tt.args.d); got != tt.want {
				t.Errorf("Round2d() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNonNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ทดสอบ 1",
			args: args{s: "1234567sjanckasc.12312"},
			want: "1234567.12312",
		}, {
			name: "ทดสอบ 2",
			args: args{s: "aaaaa"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNonNumeric(tt.args.s); got != tt.want {
				t.Errorf("removeNonNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat64(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test1",
			args: args{value: "1234567"},
			want: 1234567,
		}, {
			name: "Test2",
			args: args{value: "salkdmaslkda1234567lkdmalsdm"},
			want: 1234567,
		}, {
			name: "Test3",
			args: args{value: "asdaldasl;dka;lsdk;"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToFloat64(tt.args.value); got != tt.want {
				t.Errorf("StringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
