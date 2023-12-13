package main

import "testing"

func TestIncrementPoints(t *testing.T) {
	type args struct {
		matches int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{
				matches: 0,
			},
			want: 0,
		},
		{
			name: "one",
			args: args{
				matches: 1,
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				matches: 2,
			},
			want: 2,
		},
		{
			name: "three",
			args: args{
				matches: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPointsByMatches(tt.args.matches); got != tt.want {
				t.Errorf("CountPointsByMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	points := CountScratchcards("input_example.txt")
	expected := 13
	if points != expected {
		t.Fatalf("expected matches to be %v, but got %v", expected, points)
	}
}

func TestExampleCopies(t *testing.T) {
	copies := CountScratchcardsCopies("input_example.txt")
	expected := 30
	if copies != expected {
		t.Fatalf("expected copies to be %v, but got %v", expected, copies)
	}
}
