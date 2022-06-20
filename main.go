package main

import (
	"fmt"
	"os"

	"github.com/freggy/foodr/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:              "foodr",
		Short:            "Give me food",
		TraverseChildren: true,
	}
	rootCmd.AddCommand(cmd.CreateGenerateCommand())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
