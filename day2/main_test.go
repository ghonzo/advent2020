// Advent of Code 2020, Day 2
package main

import "testing"

func TestPasswordPolicy_valid(t *testing.T) {
	type fields struct {
		Min      int
		Max      int
		Letter   byte
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"a", fields{1, 3, 'a', "abcde"}, true},
		{"b", fields{1, 3, 'b', "cdefg"}, false},
		{"c", fields{2, 9, 'c', "ccccccccc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := passwordPolicy{
				Min:      tt.fields.Min,
				Max:      tt.fields.Max,
				Letter:   tt.fields.Letter,
				Password: tt.fields.Password,
			}
			if got := p.valid(); got != tt.want {
				t.Errorf("PasswordPolicy.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordPolicy_valid2(t *testing.T) {
	type fields struct {
		Min      int
		Max      int
		Letter   byte
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"a", fields{1, 3, 'a', "abcde"}, true},
		{"b", fields{1, 3, 'b', "cdefg"}, false},
		{"c", fields{2, 9, 'c', "ccccccccc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := passwordPolicy{
				Min:      tt.fields.Min,
				Max:      tt.fields.Max,
				Letter:   tt.fields.Letter,
				Password: tt.fields.Password,
			}
			if got := p.valid2(); got != tt.want {
				t.Errorf("PasswordPolicy.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
