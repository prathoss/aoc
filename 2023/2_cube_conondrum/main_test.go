package main

import (
	"maps"
	"testing"
)

func TestNewGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "simple line",
			args: args{
				line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: Game{
				Id: 1,
				Plays: []map[string]int{
					{
						"blue": 3,
						"red":  4,
					},
					{
						"red":   1,
						"green": 2,
						"blue":  6,
					},
					{
						"green": 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGame(tt.args.line)
			if got.Id != tt.want.Id {
				t.Errorf("want %d id, but got %d", tt.want.Id, got.Id)
			}
			if len(tt.want.Plays) != len(got.Plays) {
				t.Fatalf("number of plays do not equal, want %d, but got %d", len(tt.want.Plays), len(got.Plays))
			}
			for i := 0; i < len(got.Plays); i++ {
				if !maps.Equal(got.Plays[i], tt.want.Plays[i]) {
					t.Errorf("plays at index %d do not match, want %v, got %v", i, tt.want.Plays[i], got.Plays[i])
				}
			}
		})
	}
}

func TestGame_IsPlayable(t *testing.T) {

	type fields struct {
		Id    int
		Plays []map[string]int
	}
	type args struct {
		constraints map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "playable",
			fields: fields{
				Id: 1,
				Plays: []map[string]int{
					{
						"blue": 3,
						"red":  4,
					},
					{
						"red":   1,
						"green": 2,
						"blue":  6,
					},
					{
						"green": 2,
					},
				},
			},
			args: args{
				constraints: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			want: true,
		},
		{
			name: "unplayable",
			fields: fields{
				Id: 0,
				// 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
				Plays: []map[string]int{
					{
						"green": 8,
						"blue":  6,
						"red":   20,
					},
					{
						"blue":  5,
						"red":   4,
						"green": 13,
					},
					{
						"green": 5,
						"red":   1,
					},
				},
			},
			args: args{
				constraints: map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				Id:    tt.fields.Id,
				Plays: tt.fields.Plays,
			}
			if got := g.IsPlayable(tt.args.constraints); got != tt.want {
				t.Errorf("IsPlayable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_GetPower(t *testing.T) {
	game := NewGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	want := 48

	power := game.GetPower()

	if power != want {
		t.Errorf("expected power to be %d, but got %d", want, power)
	}
}

func SumOfPower(t *testing.T) {
	games := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	want := 2286

	sumPower := 0
	for _, strGame := range games {
		game := NewGame(strGame)
		sumPower += game.GetPower()
	}

	if sumPower != want {
		t.Fatalf("expected sum of powers to be %d, but got %d", want, sumPower)
	}
}
