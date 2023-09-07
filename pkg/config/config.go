/*
 Copyright 2023, NVIDIA CORPORATION & AFFILIATES
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package config

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/util/json"
)

// FromFile reads configuration from the file
func FromFile(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read data from the provided file: %v", err)
	}
	cfg := Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal configuration: %v", err)
	}
	if err := cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("configuration is invalid: %v", err)
	}
	return cfg, nil
}

// Config contains configuration for the init container
type Config struct {
	// configuration options for safeDriverLoading feature
	SafeDriverLoad SafeDriverLoadConfig `json:"safeDriverLoad"`
}

// SafeDriverLoadConfig contains configuration options for safeDriverLoading feature
type SafeDriverLoadConfig struct {
	// enable safeDriverLoading feature
	Enable bool `json:"enable"`
	// annotation to use for safeDriverLoading feature
	Annotation string `json:"annotation"`
}

// Validate checks the configuration
func (c *Config) Validate() error {
	if c.SafeDriverLoad.Enable && c.SafeDriverLoad.Annotation == "" {
		return fmt.Errorf(".safeDriverLoad.annotation is required if safeDriverLoad feature is enabled")
	}
	return nil
}

// String returns string representation of the configuration
func (c *Config) String() string {
	//nolint:errchkjson
	data, _ := json.Marshal(c)
	return string(data)
}
