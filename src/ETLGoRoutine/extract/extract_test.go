package extract

import (
	"testing"
)

func TestFileReader_read(t *testing.T) {
	testCases := []struct {
		name string
		f    FileReader
		want string
		err  bool
	}{
		{
			name: "positive",
			f: FileReader{
				Fname: "D:/go/src/ETL/input.txt",
			},
			want: "sfafdfsdgfg dfefsdfdsg\n" +
				"afedfsrfgdfs",
		},
		{
			name: "negative",
			f: FileReader{
				Fname: "input.txt",
			},
			err: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actualResult, errVal := testCase.f.read()
			if (errVal != nil) != testCase.err {
				t.Errorf("got unexpected error = %v, want %v", actualResult, testCase.want)
				return
			}
			if actualResult != testCase.want {
				t.Errorf("FileReader.read() = %v, want %v", actualResult, testCase.want)
			}

		})
	}
}
