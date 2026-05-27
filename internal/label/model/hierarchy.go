package model

import "strings"

// NormalizeNestedName canonicalizes a label path while preserving its display casing.
func NormalizeNestedName(name string) string {
	parts := strings.Split(strings.TrimSpace(strings.TrimPrefix(name, "#")), "/")
	normalized := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			normalized = append(normalized, part)
		}
	}
	return strings.Join(normalized, "/")
}

// NestedAncestors returns every searchable level for a label path.
func NestedAncestors(name string) []string {
	normalized := NormalizeNestedName(name)
	if normalized == "" {
		return nil
	}

	parts := strings.Split(normalized, "/")
	ancestors := make([]string, 0, len(parts))
	for i := range parts {
		ancestors = append(ancestors, strings.Join(parts[:i+1], "/"))
	}
	return ancestors
}

// NestedNameMatches reports whether a label is equal to the query or sits below it.
func NestedNameMatches(labelName string, query string) bool {
	normalizedQuery := strings.ToLower(NormalizeNestedName(query))
	if normalizedQuery == "" {
		return false
	}

	for _, ancestor := range NestedAncestors(labelName) {
		if strings.ToLower(ancestor) == normalizedQuery {
			return true
		}
	}
	return false
}

// NestedSearchText returns label terms that can be indexed for text search.
func NestedSearchText(labels []Label) string {
	terms := make([]string, 0, len(labels)*2)
	seen := make(map[string]bool)
	for _, label := range labels {
		for _, ancestor := range NestedAncestors(label.Name) {
			key := strings.ToLower(ancestor)
			if seen[key] {
				continue
			}
			seen[key] = true
			terms = append(terms, ancestor, "#"+ancestor)
		}
	}
	return strings.Join(terms, " ")
}
