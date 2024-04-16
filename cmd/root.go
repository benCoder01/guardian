/*
Copyright © 2024 Benedikt Ricken <contact@benediktricken.de>
*/
package cmd

import (
	"benedikt/infratester/tester"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "infratest",
	Short: "Test your infrastructure",
	Long: `Tests your infrastructure based on a provided .yaml file. Helps you any unwanted errors in your website configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		configuration := tester.ParseYMLToConfig(configFile)
		conditions := []tester.Condition{}

		for _,successConfig := range configuration.SuccessConditionsConfiguration {
			conditions = append(conditions, tester.SuccessCondition{Name: successConfig.Name, Url: successConfig.Url})
		}

		e := tester.Evaluator{Conditions: conditions}
		e.Run()
	},
	Args: cobra.MinimumNArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringVarP(&configFile, "config","c","", "config.yaml")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


