/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys in the file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.PrintErrln("Filename required")
			return
		}

		filename := args[0]
		buf, err := os.ReadFile(filename)
		if err != nil {
			cmd.PrintErrf("%v\n", err)
		}

		// yaml := string(buf)
		// fmt.Printf("%v\n", yaml)
		data := make(map[string]interface{})
		yaml.Unmarshal(buf, &data)

		listKeys(".", data)
	},
}

func listKeys(prefix string, data map[string]interface{}) {
	for key, value := range data {
		fmt.Printf("%v\n", strings.Join([]string{prefix, key}, "/"))
		switch t := value.(type) {
		case string:
			// do nothing
		case map[string]interface{}:
			// log.Println("Recursing")
			listKeys(prefix+"/"+key, t)
		default:
			panic("I don't know which type this is")
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
