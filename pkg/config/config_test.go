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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	configPgk "github.com/Mellanox/network-operator-init-container/pkg/config"
)

func createConfig(cfg configPgk.Config) string {
	data, err := json.Marshal(cfg)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return string(data)
}

var _ = Describe("Config test", func() {
	It("Valid - safeDriverLoad disabled", func() {
		cfg, err := configPgk.Load(createConfig(configPgk.Config{SafeDriverLoad: configPgk.SafeDriverLoadConfig{
			Enable: false,
		}}))
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg.SafeDriverLoad.Enable).To(BeFalse())
	})
	It("Valid - safeDriverLoad enabled", func() {
		cfg, err := configPgk.Load(createConfig(configPgk.Config{SafeDriverLoad: configPgk.SafeDriverLoadConfig{
			Enable:     true,
			Annotation: "something",
		}}))
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg.SafeDriverLoad.Enable).To(BeTrue())
		Expect(cfg.SafeDriverLoad.Annotation).To(Equal("something"))
	})
	It("Failed to unmarshal config", func() {
		_, err := configPgk.Load("invalid\"")
		Expect(err).To(HaveOccurred())
	})
	It("Logical validation failed - no annotation", func() {
		_, err := configPgk.Load(createConfig(configPgk.Config{SafeDriverLoad: configPgk.SafeDriverLoadConfig{
			Enable: true,
		}}))
		Expect(err).To(HaveOccurred())
	})
})
