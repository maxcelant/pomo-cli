
package fileio

import "fmt"
import "os"
import "path/filepath"
import "gopkg.in/yaml.v2"

type ConfigOptions struct {
  Active int `yaml:"active"`
  Rest   int `yaml:"rest"`
  Link   string `yaml:"link"`
}

type PomoConfig struct {
	Pomo ConfigOptions `yaml:"pomo"`
}

func convertToConfigOptions(opt map[string]interface{}) *PomoConfig {
	return &PomoConfig{
		Pomo: ConfigOptions{
			Active: opt["active"].(int),
			Rest:   opt["rest"].(int),
			Link:   opt["link"].(string),
		},
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %w", err)
	}
	configDir := filepath.Join(homeDir, ".pomo")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("could not create config directory: %w", err)
	}
	return filepath.Join(configDir, "pomo.yaml"), nil
}

func WriteToLocalYaml(o map[string]interface{}) error {
  opt := convertToConfigOptions(o)
  data, err := yaml.Marshal(&opt)
	if err != nil {
		return fmt.Errorf("could not marshal to YAML: %w", err)
	}

  configFilePath, err := getConfigFilePath()
  if err != nil {
    return err
  }

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}

func ReadFromLocalYaml(filename string) (*PomoConfig, error) {
  configFilePath, err := getConfigFilePath()
  if err != nil {
    return nil, err
  }

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err 
	}

	var config PomoConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err 
	}

	return &config, nil
}
