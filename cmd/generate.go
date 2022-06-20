package cmd

import (
	"github.com/spf13/cobra"
)

func CreateGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate recipe list",
		RunE:  RunGenerate,
	}
	cmd.Flags().String("recipes", "recipes.json", "path to the recipes json file")
	cmd.MarkFlagRequired("recipes")
	return cmd
}

func RunGenerate(cmd *cobra.Command, args []string) error {
	if path, err := cmd.Flags().GetString("recipes"); err != nil {
		return err
	}
	// TODO: read recipes file and save to cli state
	return nil
}
