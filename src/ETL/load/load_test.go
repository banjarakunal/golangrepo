package load

import "testing"

func TestLoad_LoadContent(t *testing.T) {
	type args struct {
		s interface{}
	}
	m := make(map[string]int)
	m["sfafdfsdgfg"] = 1
	m["dfefsdfdsg"] = 1
	m["afedfsrfgdfs"] = 1

	tests := []struct {
		name string
		l    Load
		args args
	}{
		{
			name: "Positive",
			l: Load{
				FileName: "input.txt",
			},
			args: args{
				s: "SFAFDFSDGFG DFEFSDFDSG\n" +
					"AFEDFSRFGDFS",
			},
		},
		{
			name: "Positive",
			l: Load{
				FileName: "input.txt",
			},
			args: args{
				s: m,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.LoadContent(tt.args.s)
		})
	}
}
