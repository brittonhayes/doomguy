package calc

import (
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"reflect"
	"testing"
)

func TestCalc_Add(t *testing.T) {
	type fields struct {
		Ctx *bot.Context
	}
	type args struct {
		m    *gateway.MessageCreateEvent
		nums []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Basic addition", fields{Ctx: new(bot.Context)}, args{nums: []int{1, 2}, m: new(gateway.MessageCreateEvent)}, "```1 + 2 = 3```", false},
		{"Adding more integers", fields{Ctx: new(bot.Context)}, args{nums: []int{1, 2, 2, 2}, m: new(gateway.MessageCreateEvent)}, "```1 + 2 + 2 + 2 = 7```", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calc{
				Ctx: tt.fields.Ctx,
			}
			got, err := c.Add(tt.args.m, tt.args.nums...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc_Divide(t *testing.T) {
	type fields struct {
		Ctx *bot.Context
	}
	type args struct {
		m *gateway.MessageCreateEvent
		a int
		b int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Basic division", fields{Ctx: new(bot.Context)}, args{a: 4, b: 2, m: new(gateway.MessageCreateEvent)}, "```4 / 2 = 2```", false},
		{"More division", fields{Ctx: new(bot.Context)}, args{a: 12, b: 3, m: new(gateway.MessageCreateEvent)}, "```12 / 3 = 4```", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calc{
				Ctx: tt.fields.Ctx,
			}
			got, err := c.Divide(tt.args.m, tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc_Multiply(t *testing.T) {
	type fields struct {
		Ctx *bot.Context
	}
	type args struct {
		m *gateway.MessageCreateEvent
		a int
		b int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Basic multiplication", fields{Ctx: new(bot.Context)}, args{a: 4, b: 2, m: new(gateway.MessageCreateEvent)}, "```4 * 2 = 8```", false},
		{"More multiplication", fields{Ctx: new(bot.Context)}, args{a: 12, b: 3, m: new(gateway.MessageCreateEvent)}, "```12 * 3 = 36```", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calc{
				Ctx: tt.fields.Ctx,
			}
			got, err := c.Multiply(tt.args.m, tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Multiply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCalc(t *testing.T) {
	tests := []struct {
		name string
		want *Calc
	}{
		{"Initialize calculator", &Calc{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCalc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
