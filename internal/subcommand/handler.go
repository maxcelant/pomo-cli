package subcommand

import (
	"fmt"
	"os"
	"strconv"
)

type Flag struct {
	datatype string
	name     string
}

var flags = map[string]Flag{
	"-a":       {datatype: "int", name: "active"},
	"--active": {datatype: "int", name: "active"},
	"-r":       {datatype: "int", name: "rest"},
	"--rest":   {datatype: "int", name: "rest"},
}

func Handler(subcommands []string, out map[string]interface{}) (map[string]interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: input was not given for one of your flags")
			os.Exit(1)
		}
	}()

	for i := 0; i < len(subcommands); i++ {
		flag := subcommands[i]

		f, found := flags[flag]
		if !found {
			return nil, fmt.Errorf("flag '%s' is not a viable flag", flag)
		}

		if i+1 >= len(subcommands) {
			return nil, fmt.Errorf("flag '%s' expects a value but none was provided", flag)
		}

		if f.datatype != "int" {
			return nil, fmt.Errorf("datatype '%s' not implemented yet", f.datatype)
		}

		nextValue := subcommands[i+1]
		duration, err := strconv.Atoi(nextValue)
		if err != nil {
			return nil, fmt.Errorf("value for flag '%s' is not a valid integer: %s", flag, nextValue)
		}

		out[f.name] = duration
		i++
	}

	return out, nil
}
