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
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:     "get <provider-name>",
	Short:   "Generate a token",
	Example: "  moth get okta-dev\n  moth get all",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		matches := match(viper.GetStringMapString(providersKey), args[0])
		clip := getClip && len(matches) == 1
		for p, s := range matches {
			code, err := totp.GenerateCode(s, time.Now())
			if err != nil {
				code = fmt.Sprintf("error generating code: %v", err)
			}
			if clip {
				clipboard.WriteAll(code)
			}
			fmt.Printf("%s | %s\n", code, p)
		}
	},
}

func match(providers map[string]string, q string) map[string]string {
	if q == allKey {
		return providers
	}

	mm := map[string]string{}
	for p, s := range providers {
		if strings.HasPrefix(p, q) {
			mm[p] = s
		}
	}
	return mm
}

var getClip bool

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&getClip, "clip", "c", false, "copy generated token to clipboard")
}
