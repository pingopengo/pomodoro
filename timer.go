package main

import (
	"fmt"
	"time"
)

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
