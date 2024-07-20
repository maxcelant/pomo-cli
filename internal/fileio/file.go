
package fileio

import "fmt"
import "os"
import "gopkg.in/yaml.v2"

type ConfigOptions struct {
  active int `yaml:"active"`
  rest   int `yaml:"rest"`
}

func convertToConfigOptions(opt map[string]interface{}) *ConfigOptions {
  return &ConfigOptions{
    active: opt["active"].(int),
    rest:   opt["rest"].(int),
  }
}

func WriteToLocalYaml(o map[string]interface{}) error {
  opt := convertToConfigOptions(o)
  data, err := yaml.Marshal(&opt)
	if err != nil {
		return fmt.Errorf("could not marshal to YAML: %w", err)
	}

	err = os.WriteFile("config.yaml", data, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
