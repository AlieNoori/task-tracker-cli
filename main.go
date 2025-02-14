package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ts       tasks
	lastId   int
	filePath string = "tasks/tasks.json"
)

func init() {
	f, _ := os.Open(filePath)
	defer f.Close()
	json.NewDecoder(f).Decode(&ts)
	if len(ts) > 0 {
		lastId = ts[len(ts)-1].Id
	}
}

func main() {
	defer saveTasks(filePath)

	var res string
	scanner := bufio.NewScanner(os.Stdin)

	for {

		printScanStd("Enhanced Task Tracker CLI\n1. Add Task\n2. Update Task Status\n3. List Tasks\n4. Delete Task\n5. Search Tasks\n6. Sort Tasks\n7. Exit\nChoose an option:", scanner, &res)

		option := strToInt(res)

		switch option {
		case 1:
			add(scanner)
		case 2:
			update(scanner)
		case 3:
			list(scanner)
		case 4:
			del(scanner)
		case 5:
			search(scanner)
		case 6:
			sortList(scanner)
		case 7:
			exit()
			return
		default:
			fmt.Println("not a valid option")
		}
	}
}

func add(scanner *bufio.Scanner) {
	var title, priority string

	printScanStd("Enter title of the task:", scanner, &title)

	printScanStd("Enter priority (High, Medium, Low):", scanner, &priority)

	fmt.Println(ts.addTask(title, priority))

	clearScreen()
}

func update(scanner *bufio.Scanner) {
	var res, update string

	printAllTasks()
	printScanStd("Enter task ID to update:", scanner, &res)
	id := strToInt(res)
	printScanStd("Which part do do you want to update:\n1. Title\n2. Status\n3. Priority\nChoose an option:", scanner, &res)
	option := strToInt(format(res))

	switch option {
	case 1:
		printScanStd("Enter new title:", scanner, &update)
	case 2:
		printScanStd("Enter new status (todo, in-progress, done):", scanner, &update)
	case 3:
		printScanStd("Enter new priority (High, Medium, Low):", scanner, &update)
	default:
		fmt.Println("There is no such field")
	}

	fmt.Println(ts.updateTask(id, res, update))
	clearScreen()
}

func list(scanner *bufio.Scanner) {
	var status string

	printScanStd("Enter status to filter by (todo, in-progress, done, all):", scanner, &status)

	fmt.Println(listTasks(searchTaskByStatus(ts, format(status))))
}

func del(scanner *bufio.Scanner) {
	var res, conf string

	printAllTasks()

	printScanStd("Enter task ID to delete:", scanner, &res)
	id := strToInt(res)

	printScanStd("Are you sure about to delete this task?(y/n):", scanner, &conf)
	if strings.ContainsAny(strings.ToLower(conf), "y") {
		fmt.Println(ts.deleteTask(id))
		clearScreen()
	} else if strings.ContainsAny(strings.ToLower(conf), "n") {
		return
	}
}

func search(scanner *bufio.Scanner) {
	var query string
	printScanStd("Enter keyword to search for tasks:", scanner, &query)
	fmt.Println(listTasks(ts.searchTask(query)))
}

func sortList(scanner *bufio.Scanner) {
	var res string
	printScanStd("Sort tasks by:\n1. ID\n2. Title\n3. Status\n4. Priority\nChoose an option:", scanner, &res)
	opt := strToInt(res)

	switch opt {
	case 1:
		sort.Sort(ById(ts))
	case 2:
		sort.Sort(ByTitle(ts))
	case 3:
		sort.Sort(ById(ts))
	case 4:
		sort.Sort(ByPriority(ts))
	}
	fmt.Println("Tasks sorted successfully")

	clearScreen()
}

func exit() {
	fmt.Println("Exit")
	os.Exit(0)
}
