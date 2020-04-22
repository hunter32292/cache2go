package cmd

import (
	"fmt"

	"github.com/hunter32292/cache2go/pkg/cache"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Service configuration state",
	Long:  `Service configuration state`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		setupService(configPath)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

}

func setupService(configPath string) {
	fmt.Println(configPath)
	cache.InitCacheService("localhost", "8080")
}
