package internal

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"time"
)

func StartTimer(pomodoros, duration, shortBreak, longBreak int, sendNotificationEnd func()) {
	var pomodorosCompleted int

	for pomodorosCompleted < pomodoros {

		if duration == 1 {
			fmt.Println("Starting timer for", duration, "minute! 🕑")
		} else {
			fmt.Println("Starting timer for", duration, "minutes! 🕑")
		}
		// progress bar
		bar := progressbar.Default(int64(duration))
		for i := 0; i < duration; i++ {
			time.Sleep(time.Second * 1)
			err := bar.Add(1)
			if err != nil {
				return
			}

			// time.Sleep(time.Duration(duration) * time.Minute)
		}
		time.Sleep(time.Duration(duration) * time.Minute)
		fmt.Println("Timer ended! 🕑")
		err := bar.Finish()
		if err != nil {
			return
		}
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
