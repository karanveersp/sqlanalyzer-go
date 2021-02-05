package analyzers

import (
	"fmt"
	"github.com/karanveersp/sqlanalyzer-go/util"
	"regexp"
)

var (
	RegexMultiLineComment  = regexp.MustCompile("(?s)(/\\*.*?\\*/)")
	RegexSingleLineComment = regexp.MustCompile("(--.*(?:\\n|$))")
)

func FindAllSingleLineComments(s string) []string {
	return RegexSingleLineComment.FindAllString(s, -1)
}

func FindAllMultiLineComments(s string) []string {
	return RegexMultiLineComment.FindAllString(s, -1)
}

func AnalyzePLSQLFile(programName, sqlPath, outputPath string) {
	fmt.Println(programName, sqlPath, outputPath)
	content := ReadPLSQLFileWithoutComments(sqlPath)
	fmt.Println(content)
}

// ReadFileWithoutComments returns all contents of the file without any single
// line or multi line comments.
func ReadPLSQLFileWithoutComments(path string) string {
	content := util.ReadFileAsString(path)
	withoutSingleLineComments := RegexSingleLineComment.ReplaceAllString(content, "")
	return RegexMultiLineComment.ReplaceAllString(withoutSingleLineComments, "")
}
