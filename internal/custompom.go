package internal

import "fmt"

func pomodoroSelection() {
	fmt.Println("Welcome to the Pomodoro Timer!")
	fmt.Println("How many pomodoros would you like to do?")
	fmt.Scan(&pomodoros)
	pomodoroSelectionDone = true
}

func setCustomTimer() {
	fmt.Println("Enter the duration of the pomodoro in minutes:")
	fmt.Scan(&duration)
	fmt.Println("Enter the duration of the short break in minutes:")
	fmt.Scan(&shortBreak)
	fmt.Println("Enter the duration of the long break in minutes:")
	fmt.Scan(&longBreak)
}
