package commands

import (
	"github.com/indraniel/srasearch/merge"
	"github.com/indraniel/srasearch/utils"
	"github.com/spf13/cobra"
	"log"
)

type MergeCmdOpts struct {
	inputMetaData string
	inputUploads  string
	inputDump     string
	output        string
}

var MergeOpts MergeCmdOpts

var cmdMerge = &cobra.Command{
	Use:   "merge -d <old-sra-dump.sjd.gz> -m <sra-incremental.tar.gz> -u <sra-incremental-uploads.gz> -o <output.sjd.gz>",
	Short: "Merge NCBI Batch Telemetry tar incrementals to a SRA Dump",
	Long: `This command merges the raw incremental NCBI Batch Telemetry
         tar file contents into an existing SRA Dump`,
	Run: func(cmd *cobra.Command, args []string) {
		MergeOpts.main()
	},
}

func init() {
	cmdMerge.Flags().StringVarP(
		&MergeOpts.inputMetaData,
		"ncbi-metadata",
		"m",
		"",
		"the incremental SRA tar.gz file",
	)

	cmdMerge.Flags().StringVarP(
		&MergeOpts.inputUploads,
		"ncbi-uploads",
		"u",
		"",
		"the incremental SRA gzipped uploads file",
	)

	cmdMerge.Flags().StringVarP(
		&MergeOpts.inputDump,
		"dump",
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

func (opts MergeCmdOpts) main() {
	opts.processOpts()
	utils.CheckFileExists(opts.inputMetaData)
	utils.CheckFileExists(opts.inputDump)
	merge.RunMerge(opts.inputMetaData, opts.inputUploads, opts.inputDump, opts.output)
}

func (opts MergeCmdOpts) processOpts() {
	if opts.output == "" {
		log.Fatal(
			"Please supply an gzip output file merge into",
			"to via --output !",
		)
	}

	if opts.inputMetaData == "" {
		log.Fatal(
			"Please supply a NCBI incremental input tar file ",
			"to via --ncbi-metadata !",
		)
	}

	if opts.inputUploads == "" {
		log.Fatal(
			"Please supply a NCBI incremental input uploads file ",
			"to via --ncbi-uploads !",
		)
	}

	if opts.inputDump == "" {
		log.Fatal(
			"Please supply an existing SRA Dump file ",
			"to merge against via --input-dump !",
		)
	}
}
