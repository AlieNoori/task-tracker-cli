package main

import (
	"fmt"
	"slices"
	"strings"
)

type task struct {
	Id       int
	Title    string
	Status   string
	Priority string
}

type tasks []task

func (t *tasks) addTask(title, priority string) string {
	task := task{Title: format(title), Priority: format(priority), Id: lastId + 1, Status: "todo"}
	(*t) = append((*t), task)

	defer saveTasks(FilePath)
	return fmt.Sprintf("Task '%v' added with ID %d", task.Title, task.Id)
}

func (t *tasks) deleteTask(id int) string {
	i := searchTaskById(*t, id)
	if i == len(*t)-1 {
		*t = (*t)[:i]
	} else {
		*t = slices.Delete(*t, i, i+1)
	}

	defer saveTasks(FilePath)
	return fmt.Sprintf("Task ID %d deleted.", (*t)[i].Id)
}

func (t *tasks) updateTask(id int, field string, update string) string {
	i := searchTaskById(*t, id)
	var s string
	switch field {
	case "priority":
		(*t)[i].Priority = update
		s = fmt.Sprintf("Task ID %d updated to %s '%s'", (*t)[i].Id, field, (*t)[i].Priority)
	case "status":
		(*t)[i].Status = update
		s = fmt.Sprintf("Task ID %d updated to %s '%s'", (*t)[i].Id, field, (*t)[i].Status)
	case "title":
		(*t)[i].Title = update
		s = fmt.Sprintf("Task ID %d updated to %s '%s'", (*t)[i].Id, field, (*t)[i].Title)
	}

	defer saveTasks(FilePath)
	return s
}

func (t *tasks) searchTask(s string) tasks {
	var res tasks
	for _, v := range *t {
		if strings.Contains(v.Title, s) {
			res = append(res, v)
		}
	}

	return res
}

func searchTaskById(t tasks, id int) int {
	n, found := slices.BinarySearchFunc(t, task{Id: id, Title: "", Status: "", Priority: ""}, func(a, b task) int {
		if a.Id > b.Id {
			return 1
		} else if a.Id < b.Id {
			return -1
		} else {
			return 0
		}
	})
	if found {
		return n
	}
	return -1
}

func searchTaskByStatus(t tasks, status string) tasks {
	if status == "all" {
		return t
	}

	var res tasks
	for _, v := range t {
		if v.Status == format(status) {
			res = append(res, v)
		}
	}

	return res
}

func listTasks(t tasks) string {
	if len(t) == 0 {
		return "There is no task to show, add one"
	}

	s := "Tasks:\n"
	for _, v := range t {
		s += fmt.Sprintf("ID: %d | Title: %s | Status: %s | Priority: %s\n", v.Id, v.Title, v.Status, v.Priority)
	}

	return s
}

type ById []task

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

type ByTitle []task

func (a ByTitle) Len() int           { return len(a) }
func (a ByTitle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTitle) Less(i, j int) bool { return a[i].Title < a[j].Title }

type ByStatus []task

func (a ByStatus) Len() int           { return len(a) }
func (a ByStatus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStatus) Less(i, j int) bool { return a[i].Status < a[j].Status }

type ByPriority []task

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }
