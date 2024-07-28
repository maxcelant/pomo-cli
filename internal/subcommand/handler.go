package subcommand

import (
	"fmt"
	"strconv"
)

type Flag struct {
	Datatype string
	Name     string
}

var flags = map[string]Flag{
	"-a":          {Datatype: "int", Name: "active"},
	"--active":    {Datatype: "int", Name: "active"},
	"-r":          {Datatype: "int", Name: "rest"},
	"--rest":      {Datatype: "int", Name: "rest"},
	"-m":          {Datatype: "bool", Name: "minimal"},
	"--minimal":   {Datatype: "bool", Name: "minimal"},
	"-i":          {Datatype: "bool", Name: "intervals"},
	"--intervals": {Datatype: "int", Name: "intervals"},
	"-l":          {Datatype: "string", Name: "link"},
	"--link":      {Datatype: "string", Name: "link"},
}

func handleInt(f Flag, value string) (int, error) {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("value for flag '%s' is not a valid integer: %s", f.Name, value)
	}
	return val, nil
}

func Handler(subcommands []string, out map[string]interface{}) (map[string]interface{}, error) {
	for i := 0; i < len(subcommands); i++ {
		cur := subcommands[i]

		f, found := flags[cur]
		if !found {
			return nil, fmt.Errorf("flag '%s' is not a viable flag\n", cur)
		}

		if (f.Datatype == "int" || f.Datatype == "string") && i+1 >= len(subcommands) {
			return nil, fmt.Errorf("flag '%s' expects a value but none was provided\n", cur)
		}

		switch f.Datatype {
		case "int":
			value, err := handleInt(f, subcommands[i+1])
			if err != nil {
				return nil, err
			}
			out[f.Name] = value
			i++
		case "string":
			out[f.Name] = subcommands[i+1]
			i++
		case "bool":
			out[f.Name] = true
		default:
			return nil, fmt.Errorf("datatype '%s' not implemented yet\n", f.Datatype)
		}
	}

	return out, nil
}
