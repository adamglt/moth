// Copyright © 2018 adamglt <adamg@traiana.com>
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var delYes bool

var deleteCmd = &cobra.Command{
	Use:     "delete <provider-name>",
	Short:   "Delete a provider",
	Example: "  moth delete okta-dev",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]

		if provider == allKey {
			if !delYes {
				failf("to delete all, run the command with the '-yes' flag\n")
			}
			viper.Set(providersKey, map[string]string{})
			writeCfg()
		} else {
			m := viper.GetStringMap(providersKey)
			delete(m, provider)
			viper.Set(providersKey, m)
			writeCfg()
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVar(&delYes, "yes", false, "confirm 'all' deletion")
}
