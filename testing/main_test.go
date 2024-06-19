package main

import (
	"testing"
)

func TestMakeTodos(t *testing.T) {
	data := ReadFile("./testData/2024-06-13.md")
	got := MakeTodos(data)
	var cases []Todo
	cases = append(cases, Todo{
		task: "- [ ] test todo",
		info: []string{"    - info for todo"},
	})
	cases = append(cases, Todo{
		task: "- [ ] test todo",
		info: nil,
	})
	cases = append(cases, Todo{
		task: "- [ ] make todo app",
		info: []string{"    - how to make todo app"},
	})
	cases = append(cases, Todo{
		task: "- [ ] make todo app",
		info: []string{
			"    - info for todo",
			"    - how to make todo app",
		},
	})

	for i, want := range cases {
		if i >= len(got) {
			t.Error("Not enougth todos")
		}

		if got[i].task != want.task {
			t.Errorf("got %q want %q", got[i].task, want.task)
		}

		if len(got[i].info) != len(want.info) {
			t.Errorf("got info len: %q want info len: %q", len(got[i].info), len(want.info))
			continue
		}

		for j, wantInfo := range want.info {
			if got[i].info[j] != wantInfo {
				t.Errorf("got %q want %q", got[i].info[j], wantInfo)
			}
		}
	}
}
