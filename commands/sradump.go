package commands

import (
	"github.com/indraniel/srasearch/sradump"
	"github.com/indraniel/srasearch/utils"
	"github.com/spf13/cobra"
	"log"
)

type SraDumpCmdOpts struct {
	output   string
	metadata string
	uploads  string
}

var SraDumpOpts SraDumpCmdOpts

var cmdSraDump = &cobra.Command{
	Use:   "sra-dump -o <output> -m <ncbi-metadata.tar.gz> -u <ncbi-uploads.gz>",
	Short: "Transform the NCBI Batch Telemetry tar files to a set of JSON Docs",
	Long: `This command transforms the raw NCBI Batch Telemetry tar file
         contents into a intermediary file of custom JSON Documents`,
	Run: func(cmd *cobra.Command, args []string) {
		SraDumpOpts.main()
	},
}

func init() {
	cmdSraDump.Flags().StringVarP(
		&SraDumpOpts.output,
		"output",
		"o",
		"sradump.sjd.gz",
		"the output file to dump the serialized JSON Documents to",
	)

	cmdSraDump.Flags().StringVarP(
		&SraDumpOpts.metadata,
		"ncbi-metadata",
		"m",
		"sradump.sjd.gz",
		"the NCBI SRA Metadata tar.gz file",
	)

	cmdSraDump.Flags().StringVarP(
		&SraDumpOpts.uploads,
		"ncbi-uploads",
		"u",
		"sradump.sjd.gz",
		"teh NCBI SRA Uploads gzip file",
	)
}

func (opts SraDumpCmdOpts) main() {
	SraDumpOpts.processOpts()
	utils.CheckFileExists(opts.metadata)
	utils.CheckFileExists(opts.uploads)
	sradump.RunSraDump(opts.metadata, opts.uploads, opts.output)
}

func (opts SraDumpCmdOpts) processOpts() {
	if opts.output == "" {
		log.Fatal(
			"Please supply an gzip output file to dump",
			"to via --output !",
		)
	}

	if opts.metadata == "" {
		log.Fatal(
			"Please supply the NCBI SRA Metadata file",
			"via --metadata !",
		)
	}

	if opts.uploads == "" {
		log.Fatal(
			"Please supply the NCBI SRA Uploads file",
			"via --uploads !",
		)
	}
}
