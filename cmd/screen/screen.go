
package screen

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Usage() {
	fmt.Printf("Usage: pomo [command]\n\n")
  fmt.Println("       Commands           Actions")
  fmt.Println("       start              Begin a simple pomodoro interval session.")
}
