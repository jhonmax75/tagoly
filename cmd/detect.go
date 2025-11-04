package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// detectCmd analisa uma mensagem de commit e extrai tipo e escopo
var detectCmd = &cobra.Command{
	Use:   "detect [mensagem do commit]",
	Short: "Detecta o tipo e escopo do commit",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]

		// Expressão regular para Conventional Commits
		re := regexp.MustCompile(`^(feat|fix|chore|docs|style|refactor|perf|test|build|ci|revert)(\(([^)]+)\))?:`)
		match := re.FindStringSubmatch(message)

		if match != nil {
			commitType := match[1]
			scope := match[3] // pode ser vazio se não houver escopo
			fmt.Println("Tipo do commit:", commitType)
			if scope != "" {
				fmt.Println("Escopo:", scope)
			} else {
				fmt.Println("Escopo: (nenhum)")
			}
		} else {
			fmt.Println("Mensagem de commit inválida ou não segue o padrão Conventional Commits")
		}
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
}
