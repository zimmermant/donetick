package chore

import (
	"testing"

	chModel "donetick.com/core/internal/chore/model"
	lModel "donetick.com/core/internal/label/model"
)

func TestFilterChoresByNestedLabel(t *testing.T) {
	chores := []*chModel.Chore{
		{ID: 1, Name: "Feed Max", LabelsV2: &[]lModel.Label{{Name: "home/pets/max"}}},
		{ID: 2, Name: "Water fern", LabelsV2: &[]lModel.Label{{Name: "home/plants"}}},
	}

	got := filterChoresByNestedLabel(chores, "#home")
	if len(got) != 2 {
		t.Fatalf("filterChoresByNestedLabel(%q) = %#v, want both home chores", "#home", got)
	}

	for _, query := range []string{"home/pets", "home/pets/max"} {
		got := filterChoresByNestedLabel(chores, query)
		if len(got) != 1 || got[0].ID != 1 {
			t.Fatalf("filterChoresByNestedLabel(%q) = %#v, want only chore 1", query, got)
		}
	}
}

func TestSearchChoresByNestedLabel(t *testing.T) {
	chores := []*chModel.Chore{
		{ID: 1, Name: "Feed Max", LabelsV2: &[]lModel.Label{{Name: "home/pets/max"}}},
		{ID: 2, Name: "Water fern", LabelsV2: &[]lModel.Label{{Name: "home/plants"}}},
	}

	got := searchChoresByNestedLabel(chores, "#home/pets")
	if len(got) != 1 || got[0].ID != 1 {
		t.Fatalf("searchChoresByNestedLabel(%q) = %#v, want only chore 1", "#home/pets", got)
	}
}
