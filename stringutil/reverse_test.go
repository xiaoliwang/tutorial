package stringutil

import "testing"

func TestReverse(t *testing.T) {
    cases := [] struct {
        in, want string
    } {
        {"Hellow, world", "dlrow ,wolleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}