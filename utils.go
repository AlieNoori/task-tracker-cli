package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func saveTasks(fp string) {
	os.Remove(fp)
	file, err := os.Create("tasks/tasks.json")
	if err != nil {
		fmt.Println("error crating json file:", err)
		os.Exit(1)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(ts)
}

func clearScreen() {
	time.Sleep(time.Second * 2)
	fmt.Print("\033[H\033[2J")
}

func printScanStd(s string, scanner *bufio.Scanner, v *string) {
	fmt.Print(s)
	*v = ""
	if scanner.Scan() {
		*v += scanner.Text()
	}
}

func strToInt(s string) int {
	s = format(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(errors.New("not a valid option"))
	}
	return i
}

func format(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func printAllTasks() {
	fmt.Println(listTasks(ts))
}
