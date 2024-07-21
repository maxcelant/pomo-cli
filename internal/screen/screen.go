package screen

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Usage() {
	fmt.Printf("Usage: pomo [command] [options]\n\n")
	fmt.Println("Commands and their options:")
	fmt.Println("  start                Begin a simple pomodoro interval session.")
	fmt.Println("    -m, --minimal      Runs minimal command line display.")
	fmt.Println("    -i, --intervals    Set the number of intervals to accomplish.")
	fmt.Println("    -l, --log          Prompt to log accomplishments after every rest.")
	fmt.Println()
	fmt.Println("  config               Configure active and rest times.")
	fmt.Println("    -a, --active       Set the amount of time to work (e.g., 45m).")
	fmt.Println("    -r, --rest         Set the amount of time to rest (e.g., 15m).")
	fmt.Println("    -l, --link         Link to your obsidian notes (e.g., /link/to/obsidian).")
	fmt.Println("    -f, --file         Optionally send a pomo.yaml file with your options.")
}
