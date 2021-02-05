package main

import (
	"flag"
	"fmt"
	"github.com/karanveersp/sqlanalyzer-go/analyzers"
	"os"
	"path/filepath"
)

func validateParams(programName, sqlPath, outputPath string) {
	if programName == "" || sqlPath == "" || outputPath == "" {
		fmt.Println("Program name, sql file path and output path are all required.")
		flag.Usage()
		os.Exit(1)
	}
}

func getInput(programName, sqlPath, outputPath string) (string, string, string) {
	fmt.Println("Enter Program Name:")
	_, _ = fmt.Scanln(&programName)

	fmt.Println("Enter PL/SQL file path:")
	_, _ = fmt.Scanln(&sqlPath)

	fmt.Println("Enter output xlsx path:")
	_, _ = fmt.Scanln(&outputPath)

	return programName, filepath.Clean(sqlPath), filepath.Clean(outputPath)
}

func getArgs() (string, string, string) {
	// declare variables
	var programName, sqlPath, outputPath string
	var isInteractiveMode bool

	// flags declaration
	flag.StringVar(&programName, "pname", "", "Program name to fill in output")
	flag.StringVar(&sqlPath, "sql", "", "Path to PL/SQL file to analyze")
	flag.StringVar(&outputPath, "out", "", "Path to output .xlsx")
	flag.BoolVar(&isInteractiveMode, "i", false, "Whether to run in interactive mode")

	flag.Parse()

	if isInteractiveMode {
		return getInput(programName, sqlPath, outputPath)
	}

	validateParams(programName, sqlPath, outputPath)

	return programName, filepath.Clean(sqlPath), filepath.Clean(outputPath)
}

func main() {
	programName, sqlPath, outputPath := getArgs()
	analyzers.AnalyzePLSQLFile(programName, sqlPath, outputPath)
}
