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
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:     "add <provider-name> <secret>",
	Short:   "Add a provider",
	Example: "  moth add okta-dev ABCD1234ABCD1234",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		provider, secret := args[0], args[1]

		if provider == allKey {
			failf("the provider 'all' is reserved\n")
		}

		code, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			failf("secret validation failed: %v\n", err)
		}

		viper.Set(pkey(provider), secret)
		writeCfg()

		if addClip {
			clipboard.WriteAll(code)
		}
		fmt.Printf("%s | %s\n", code, provider)
	},
}

var addClip bool

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&addClip, "clip", "c", false, "copy generated token to clipboard")
}
