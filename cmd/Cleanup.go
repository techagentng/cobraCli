/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"todocli/model"

	"github.com/spf13/cobra"
)

// CleanupCmd represents the Cleanup command
var CleanupCmd = &cobra.Command{
	Use:   "Cleanup",
	Short: "Delete all the Done Tasks from the lists",
	Long: `Clear all the Done/Checked Tasks all at once`,
	Run: func(cmd *cobra.Command, args []string) {
		model.Lst.Cleanup()
	},
}

func init() {
	rootCmd.AddCommand(CleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
