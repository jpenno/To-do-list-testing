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
		info: "    - info for todo",
	})
	cases = append(cases, Todo{
		task: "- [ ] test todo",
	})
	cases = append(cases, Todo{
		task: "- [ ] make todo app",
		info: "    - how to make todo app",
	})
	cases = append(cases, Todo{
		task: "- [ ] make todo app",
		info: "    - info for todo    - how to make todo app",
	})

	for i, want := range cases {
		if got[i].task != want.task {
			t.Errorf("got %q want %q", got[i].task, want.task)
		}

		if got[i].info != want.info {
			t.Errorf("got %q want %q", got[i].info, want.info)
		}
	}
}
