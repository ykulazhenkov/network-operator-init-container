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

package config_test

import (
	"encoding/json"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	configPgk "github.com/Mellanox/network-operator-init-container/pkg/config"
)

func createConfig(path string, cfg configPgk.Config) {
	data, err := json.Marshal(cfg)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	ExpectWithOffset(1, os.WriteFile(path, data, 0x744))
}

var _ = Describe("Config test", func() {
	var (
		configPath string
	)
	BeforeEach(func() {
		configPath = filepath.Join(GinkgoT().TempDir(), "config")
	})
	It("Valid - safeDriverLoad disabled", func() {
		createConfig(configPath, configPgk.Config{SafeDriverLoad: configPgk.SafeDriverLoadConfig{
			Enable: false,
		}})
		cfg, err := configPgk.FromFile(configPath)
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg.SafeDriverLoad.Enable).To(BeFalse())
	})
	It("Valid - safeDriverLoad enabled", func() {
		createConfig(configPath, configPgk.Config{SafeDriverLoad: configPgk.SafeDriverLoadConfig{
			Enable:     true,
			Annotation: "something",
		}})
		cfg, err := configPgk.FromFile(configPath)
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg.SafeDriverLoad.Enable).To(BeTrue())
		Expect(cfg.SafeDriverLoad.Annotation).To(Equal("something"))
	})
	It("Failed to read config", func() {
		_, err := configPgk.FromFile(configPath)
		Expect(err).To(HaveOccurred())
	})
	It("Failed to unmarshal config", func() {
		Expect(os.WriteFile(configPath, []byte("invalid\""), 0x744)).NotTo(HaveOccurred())
		_, err := configPgk.FromFile(configPath)
		Expect(err).To(HaveOccurred())
	})
	It("Logical validation failed", func() {
		createConfig(configPath, configPgk.Config{SafeDriverLoad: configPgk.SafeDriverLoadConfig{
			Enable: true,
		}})
		_, err := configPgk.FromFile(configPath)
		Expect(err).To(HaveOccurred())
	})
})
