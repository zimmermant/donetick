package model

import (
	"strings"
	"testing"
)

func TestNestedAncestors(t *testing.T) {
	got := NestedAncestors("#home / pets / max")
	want := []string{"home", "home/pets", "home/pets/max"}

	if len(got) != len(want) {
		t.Fatalf("NestedAncestors length = %d, want %d: %#v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("NestedAncestors[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestNestedNameMatches(t *testing.T) {
	label := "home/pets/max"

	for _, query := range []string{"home", "#home", "home/pets", "HOME/PETS/MAX"} {
		if !NestedNameMatches(label, query) {
			t.Fatalf("NestedNameMatches(%q, %q) = false, want true", label, query)
		}
	}

	if NestedNameMatches(label, "home/pets/max/toys") {
		t.Fatalf("NestedNameMatches(%q, %q) = true, want false", label, "home/pets/max/toys")
	}
}

func TestNestedSearchText(t *testing.T) {
	text := NestedSearchText([]Label{{Name: "home/pets/max"}})

	for _, want := range []string{"home", "#home", "home/pets", "#home/pets", "home/pets/max", "#home/pets/max"} {
		if !stringsContainsToken(text, want) {
			t.Fatalf("NestedSearchText missing %q in %q", want, text)
		}
	}
}

func stringsContainsToken(text string, token string) bool {
	for _, part := range strings.Fields(text) {
		if part == token {
			return true
		}
	}
	return false
}
