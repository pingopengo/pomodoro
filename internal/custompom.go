package internal

import "fmt"

func PomodoroSelection() (int, bool) {
	fmt.Println("Welcome to the Pomodoro Timer!")
	fmt.Println("How many pomodoros would you like to do?")
	var pomodoros int
	fmt.Scan(&pomodoros)
	return pomodoros, true
}

func SetCustomTimer() (int, int, int) {
	var duration, shortBreak, longBreak int
	fmt.Println("Enter the duration of the pomodoro in minutes:")
	fmt.Scan(&duration)
	fmt.Println("Enter the duration of the short break in minutes:")
	fmt.Scan(&shortBreak)
	fmt.Println("Enter the duration of the long break in minutes:")
	fmt.Scan(&longBreak)
	return duration, shortBreak, longBreak
}
