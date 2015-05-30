package commands

import (
	"github.com/indraniel/srasearch/makeindex"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type MkIdxCmdOpts struct {
	input  string
	output string
}

var MkIdxOpts MkIdxCmdOpts

var cmdMakeIndex = &cobra.Command{
	Use:   "make-index -i <sra-dump.sjd> -o <srasearch.idx>",
	Short: "Generate an search index file from a SRA Dump file",
	Long:  `Generate an search index file from a SRA Dump file`,
	Run: func(cmd *cobra.Command, args []string) {
		MkIdxOpts.processOpts()
		MkIdxOpts.mainRun()
	},
}

func init() {
	cmdMakeIndex.Flags().StringVarP(
		&MkIdxOpts.output,
		"output",
		"o",
		"srasearch.idx",
		"the base index for the search engine",
	)

	cmdMakeIndex.Flags().StringVarP(
		&MkIdxOpts.input,
		"input",
		"i",
		"sradump.sjd.gz",
		"the input gzipped SRA dump file",
	)
}

func (opts MkIdxCmdOpts) mainRun() {
	makeindex.CreateSearchIndex(opts.input, opts.output)
}

func (opts MkIdxCmdOpts) processOpts() {
	if opts.output == "" {
		log.Fatal("Please supply an output file via --output !")
	}

	if opts.input == "" {
		log.Fatal("Please supply an input file via --input !")
	}

	opts.checkExists(opts.input)
}

func (opts MkIdxCmdOpts) checkExists(inputfile string) {
	if _, err := os.Stat(inputfile); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			inputfile, err,
		)
	}
}
