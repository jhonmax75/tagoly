package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd exibe informações da versão do Tagoly
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Mostra a versão atual do Tagoly",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🧠 Tagoly CLI - Commit Scope Detector")
		fmt.Println("Versão: v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
