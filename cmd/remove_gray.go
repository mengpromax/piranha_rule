package cmd

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mengpromax/piranha_rule/pkg"
)

const (
	regrexSeparator = "|"
)

// removeGrayCmd represents the removeGray command
var removeGrayCmd = &cobra.Command{
	Use: "remove_gray [-src source_file_path] gray_key_to_be_removed",
	Short: "Remove gray key related logic including marking gray boolean judging to true and " +
		"boolean expressions simplification.",
	Run: func(cmd *cobra.Command, args []string) {
		// 支持传多个 key
		if len(args) == 0 {
			log.Fatalf("should have at least 1 gray key\n")
		}
		grayKeys := args
		log.Printf("target gray key list: %v\n", grayKeys)

		workDirPath, err := os.Getwd()
		if err != nil {
			log.Fatalf("can not get work dir path, err: %v\n", err)
		}

		absoluteSrcDir := path.Join(workDirPath, srcRelativeDir)

		constantMapping, err := pkg.ScanAllConstantsMapping(absoluteSrcDir, grayKeys, nil)
		if err != nil {
			log.Fatalf("scan constant error, err: %v\n", err)
		}

		constantKeyList := make([]string, 0, len(constantMapping))
		for constantKey := range constantMapping {
			constantKeyList = append(constantKeyList, constantKey)
		}

		constantGrayKeyMatchPattern := strings.Join(constantKeyList, regrexSeparator)
		grayKeyMatchPattern := strings.Join(grayKeys, regrexSeparator)

		curExecutable, err := os.Executable()
		if err != nil {
			panic(err)
		}

		curExecutableDir := filepath.Dir(curExecutable)
		err = pkg.RemoveGrayKey(absoluteSrcDir, constantGrayKeyMatchPattern, grayKeyMatchPattern, curExecutableDir)
		if err != nil {
			log.Fatalf("remove gray key executor error, err: %v\n", err)
		}

		log.Printf("remove gray key success, please double check your repository changes!\n")
	},
}

var (
	srcRelativeDir string
)

func init() {
	// FIXME: 支持忽略路径参数
	rootCmd.AddCommand(removeGrayCmd)
	removeGrayCmd.Flags().StringVarP(&srcRelativeDir, "src", "s", ".", "Path of target project to remove gray from.")
}
