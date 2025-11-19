package note

import (
	"reflect"
	"testing"
)

func TestGetFilteredTagNoteIds(t *testing.T) {
	tests := []struct {
		name     string
		tagIds   []int
		expected []int
	}{
		{name: "find 1 tag", tagIds: []int{1}, expected: []int{1, 2, 3}},
		{name: "find 2 tags", tagIds: []int{1, 2}, expected: []int{1, 2}},
		{name: "find 2 tags ver2", tagIds: []int{2, 3}, expected: []int{2}},
		{name: "find 3 tags", tagIds: []int{1, 2, 3}, expected: []int{2}},
	}
	tags := map[int][]int{1: {1, 2, 3}, 2: {1, 2}, 3: {2, 3}}

	im := IndexManager{i: Index{Tags: tags}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := im.GetFilteredTagNoteIds(tt.tagIds...)

			equal := reflect.DeepEqual(res, tt.expected)

			if !equal {
				t.Errorf("GetFilteredTagNoteIds(%d) = %d; expected %d", tt.tagIds, res, tt.expected)
			}
		})
	}

}

func BenchmarkGetFilteredTagNoteIds(b *testing.B) {
	tests := []struct {
		name   string
		tagIds []int
	}{
		{name: "find 1 tag", tagIds: []int{1}},
		{name: "find 2 tags", tagIds: []int{1, 2}},
		{name: "find 2 tags ver2", tagIds: []int{2, 3}},
		{name: "find 3 tags", tagIds: []int{1, 2, 3}},
	}

	tags := map[int][]int{1: {1, 2, 3}, 2: {1, 2}, 3: {2, 3}}
	im := IndexManager{i: Index{Tags: tags}}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				im.GetFilteredTagNoteIds(tt.tagIds...)
			}

		})
	}
}
