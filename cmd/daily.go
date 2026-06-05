package main

import (
	"fmt"

	"github.com/spf13/cobra"
	destroyingasteroids "github.com/0226zy/lc-go/solutions/daily/2126_destroying_asteroids"
)

func init() {
	rootCmd.AddCommand(destroyingAsteroidsCmd)
}

var destroyingAsteroidsCmd = &cobra.Command{
	Use:     "destroying-asteroids",
	Short:   "摧毁小行星 (Destroying Asteroids)",
	Aliases: []string{"2126"},
	Run: func(cmd *cobra.Command, args []string) {
		mass := 10
		asteroids := []int{3, 9, 19, 5, 21}
		result := destroyingasteroids.AsteroidsDestroyed(mass, asteroids)
		fmt.Printf("AsteroidsDestroyed(%d, %v) = %v\n", mass, asteroids, result)
	},
}
