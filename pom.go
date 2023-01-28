package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"time"
)

var (
	duration              = 25
	shortBreak            = 5
	longBreak             = 15
	pomodoros             int
	pomodoroSelectionDone bool
)

func main() {
	for {
		if !pomodoroSelectionDone {
			pomodoroSelection()
			n := gosxnotifier.NewNotification("Pomodoro session started! 🕑")
			n.Title = "Pomodoro Timer"
			n.Sound = gosxnotifier.Basso
			n.AppIcon = "pomodoro.png"
			n.Push()
		} else if pomodoroSelectionDone == true {

			fmt.Println("1. Start timer with default settings! 😊")
			fmt.Println("2. Start timer with custom settings! 🤓")
			fmt.Println("3. Exit! 😢")

			var input int
			fmt.Print("Enter your choice: ")
			fmt.Scan(&input)

			switch input {
			case 1:
				startTimer(duration, shortBreak, longBreak)

			case 2:
				setCustomTimer()
				startTimer(duration, shortBreak, longBreak)

			case 3:
				fmt.Println("die stupid! 😡")
				return
			default:
				fmt.Println("Invalid input, try again 😤")
			}
		}
	}

}

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

func startTimer(duration, shortBreak, longBreak int) {
	var pomodorosCompleted int
	for pomodorosCompleted < pomodoros {
		if duration == 1 {
			fmt.Println("Starting timer for", duration, "minute! 🕑")
		} else {
			fmt.Println("Starting timer for", duration, "minutes! 🕑")
		}

		time.Sleep(time.Duration(duration) * time.Minute)
		fmt.Println("Timer ended! 🕑")
		pomodorosCompleted++
		sendNotificationEnd()
		if pomodorosCompleted%4 == 0 {
			fmt.Println("Starting timer for", longBreak, "minutes! 🕑")
			time.Sleep(time.Duration(longBreak) * time.Minute)
			fmt.Println("Timer ended! 🕑")
			sendNotificationEnd()
		} else if pomodorosCompleted == pomodoros {
			fmt.Println("You have completed", pomodoros, "pomodoros! 🎉")
			break
		} else {
			fmt.Println("Starting timer for", shortBreak, "minutes! 🕑")
			time.Sleep(time.Duration(shortBreak) * time.Minute)
			fmt.Println("Timer ended! 🕑")
			sendNotificationEnd()
		}
	}

	fmt.Println("Timer started! 🕑")
	fmt.Println("You have", duration, "minutes to work! 📝")
	fmt.Println("You have", shortBreak, "minutes for a short break! ☕️")
	fmt.Println("You have", longBreak, "minutes for a long break! 🍃")
}

func sendNotificationEnd() {
	n := gosxnotifier.NewNotification("Timer ended! 🕑")
	n.Title = "Pomodoro Timer"
	n.Sound = gosxnotifier.Basso
	n.AppIcon = "pomodoro.png"
	n.Push()
}
