package subcommand

import (
	"fmt"
	"os"
	"strconv"
)

type flag struct {
	datatype string
	name     string
}

var flags = map[string]flag{
	"-a":       {datatype: "int", name: "active"},
	"--active": {datatype: "int", name: "active"},
	"-r":       {datatype: "int", name: "rest"},
	"--rest":   {datatype: "int", name: "rest"},
	"-d":       {datatype: "bool", name: "detach"},
	"--detach": {datatype: "bool", name: "detach"},
}

func handleInt(f flag, value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("value for flag '%s' is not a valid integer: %s", f.name, value)
		os.Exit(1)
	}
	return val
}

func Handler(subcommands []string, out map[string]interface{}) (map[string]interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: input was not given for one of your flags")
			os.Exit(1)
		}
	}()

	for i := 0; i < len(subcommands); i++ {
		cur := subcommands[i]

		f, found := flags[cur]
		if !found {
			return nil, fmt.Errorf("flag '%s' is not a viable flag", cur)
		}

		if (f.datatype == "int" || f.datatype == "string") && i+1 >= len(subcommands) {
			return nil, fmt.Errorf("flag '%s' expects a value but none was provided", cur)
		}

		if f.datatype != "int" && f.datatype != "bool" {
			return nil, fmt.Errorf("datatype '%s' not implemented yet", f.datatype)
		}

		if f.datatype == "int" {
			out[f.name] = handleInt(f, subcommands[i+1])
			i++
		}

		if f.datatype == "bool" {
			out[f.name] = true
		}
	}

	return out, nil
}
