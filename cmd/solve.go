package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ddonukis/advent-of-code-golang/internal/router"
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve Advent of Code puzzle for specified year and day.",
	Long: `Solve Advent of Code puzzle for specified year and day.

Example:
> aoc-go solve 2024 1
Part1: 123
Part2: 456`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("solve accepts %d arg(s), received %d", 2, len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("year: %s day: %s\n", args[0], args[1])

		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("first argument must be a valid interger representing a year, recieved: %s\n", args[0])
			os.Exit(1)
		}
		day, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("second argument must be a valid interger representing a day, recieved: %s\n", args[1])
			os.Exit(1)
		}

		inputFilePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalln(err)
		}

		if inputFilePath == "" {
			inputFilePath, err = defaultFilePath(year, day)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("default data path:", inputFilePath)
		}

		router.RunSolver(year, day, inputFilePath)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	solveCmd.Flags().StringP("file", "f", "", "path to the input file")

}

func defaultFilePath(year, day int) (path string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path = filepath.Join(
		dir,
		"data",
		strconv.Itoa(year),
		fmt.Sprintf("%02d", day),
		"input.txt",
	)
	return
}
