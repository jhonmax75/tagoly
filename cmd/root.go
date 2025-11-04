package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versão global da ferramenta
var Version = "v0.1.0"

// rootCmd representa o comando base quando chamado sem subcomandos
var rootCmd = &cobra.Command{
	Use:   "tagoly",
	Short: "Tagoly - Detecção automática de escopo de commit",
	Long: `Tagoly é uma ferramenta CLI desenvolvida em Go
para identificar automaticamente o escopo de commits Git
e garantir conformidade com o padrão Conventional Commits.`,
	SilenceUsage: true, // evita exibir ajuda em erros
}

// Execute adiciona todos os subcomandos e configurações globais
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Arquivo de configuração (padrão: $HOME/.tagoly.yaml)")
}

var cfgFile string

// initConfig lê variáveis de ambiente e configurações externas, se existirem
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tagoly")
	}

	viper.AutomaticEnv() // lê variáveis de ambiente compatíveis

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Usando arquivo de configuração:", viper.ConfigFileUsed())
	}
}
