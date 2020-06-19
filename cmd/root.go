// Copyright © 2018 Adam Gilat <adamg@traiana.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgDefault     = ".moth"
	cfgDefaultFile = cfgDefault + ".yaml"

	allKey       = "all"
	providersKey = "providers"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "moth",
	Short: "Take the 'oof' out of twofactor",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		failf("%v\n", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		fmt.Sprintf("config file (default is $HOME/%s)", cfgDefaultFile),
	)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			failf("%v\n", err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgDefault)
		cfgFile = path.Join(home, cfgDefaultFile)
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func pkey(provider string) string {
	return fmt.Sprintf("%s.%s", providersKey, provider)
}

func writeCfg() {
	if err := viper.WriteConfigAs(cfgFile); err != nil {
		failf("failed to write config: %v\n", err)
	}
}

func failf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
