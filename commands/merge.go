package commands

import (
	"github.com/indraniel/srasearch/merge"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type MergeCmdOpts struct {
	inputTar  string
	inputDump string
	output    string
}

var MergeOpts MergeCmdOpts

var cmdMerge = &cobra.Command{
	Use:   "merge -d <old-sra-dump.sjd.gz> -t <sra-incremental.tar.gz> -o <output.sjd.gz>",
	Short: "Merge NCBI Batch Telemetry tar incrementals to a SRA Dump",
	Long: `This command merges the raw incremental NCBI Batch Telemetry
         tar file contents into an existing SRA Dump`,
	Run: func(cmd *cobra.Command, args []string) {
		MergeOpts.processOpts()
		checkFileExists(MergeOpts.inputTar)
		checkFileExists(MergeOpts.inputDump)
		MergeOpts.mainRun()
	},
}

func init() {
	cmdMerge.Flags().StringVarP(
		&MergeOpts.inputTar,
		"input-tar",
		"t",
		"",
		"the incremental SRA tar file",
	)

	cmdMerge.Flags().StringVarP(
		&MergeOpts.inputDump,
		"input-dump",
		"d",
		"sradump.sjd.gz",
		"the existing SRA dump file to merge into",
	)

	cmdMerge.Flags().StringVarP(
		&MergeOpts.output,
		"output",
		"o",
		"merged.sjd.gz",
		"the merged output file to dump the serialized JSON Documents to",
	)
}

func (opts MergeCmdOpts) mainRun() {
	merge.RunMerge(opts.inputTar, opts.inputDump, opts.output)
}

func (opts MergeCmdOpts) processOpts() {
	if opts.output == "" {
		log.Fatal(
			"Please supply an gzip output file merge into",
			"to via --output !",
		)
	}

	if opts.inputTar == "" {
		log.Fatal(
			"Please supply an incremental input tar file ",
			"to via --input-tar !",
		)
	}

	if opts.inputDump == "" {
		log.Fatal(
			"Please supply an existing SRA Dump file ",
			"to merge against via --input-dump !",
		)
	}
}

func checkFileExists(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			file, err,
		)
	}
}
