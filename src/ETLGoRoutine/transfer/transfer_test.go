package transfer

import (
	"reflect"
	"testing"
)

func TestConvertUpper(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive",
			args: args{
				content: "sfafdfsdgfg dfefsdfdsg\n" +
					"afedfsrfgdfs",
			},
			want: "SFAFDFSDGFG DFEFSDFDSG\n" +
				"AFEDFSRFGDFS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertUpper(tt.args.content); got != tt.want {
				t.Errorf("ConvertUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	type args struct {
		content string
	}
	m := make(map[string]int)
	m["sfafdfsdgfg"] = 1
	m["dfefsdfdsg"] = 1
	m["afedfsrfgdfs"] = 1

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Positive",
			args: args{
				content: "sfafdfsdgfg dfefsdfdsg\n" +
					"afedfsrfgdfs",
			},
			want: m,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
