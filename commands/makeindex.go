package commands

import (
	"github.com/indraniel/srasearch/makeindex"
	"github.com/indraniel/srasearch/utils"

	"github.com/spf13/cobra"

	"log"
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
		MkIdxOpts.main()
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

func (opts MkIdxCmdOpts) main() {
	opts.processOpts()
	makeindex.CreateSearchIndex(opts.input, opts.output)
}

func (opts MkIdxCmdOpts) processOpts() {
	if opts.output == "" {
		log.Fatal("Please supply an output file via --output !")
	}

	if opts.input == "" {
		log.Fatal("Please supply an input file via --input !")
	}

	utils.CheckFileExists(opts.input)
}
