package main

import "testing"

func TestReadByLine(t *testing.T) {
	type args struct {
		fileName    string
		lineHandler func(previousLine, currentLine, nextLine string, currentLineNumber int)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple",
			args: args{
				fileName:    "test_input_lines.txt",
				lineHandler: TstLineHandler,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadByLine(tt.args.fileName, tt.args.lineHandler)
		})
	}
}

func TstLineHandler(previousLine, currentLine, nextLine string, currentLineNumber int) {
	if currentLine == "" {
		panic("current line can not be empty")
	}
}

func TestSample(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "sample",
			input: "test_sample.txt",
			want:  4361,
		},
		{
			name:  "1",
			input: "test_input_1.txt",
			want:  1,
		},
		{
			name:  "2",
			input: "test_input_2.txt",
			want:  3,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spn := SumPartNumbers(tc.input)

			if spn != tc.want {
				t.Fatalf("expected sum of part numbers to be %d, but got %d", tc.want, spn)
			}
		})
	}
}
