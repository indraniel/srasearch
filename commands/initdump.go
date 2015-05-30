package commands

import (
	"github.com/indraniel/srasearch/initdump"
	"github.com/indraniel/srasearch/utils"
	"github.com/spf13/cobra"
	"log"
)

type InitDumpCmdOpts struct {
	output   string
	metadata string
	uploads  string
}

var InitDumpOpts InitDumpCmdOpts

var cmdInitDump = &cobra.Command{
	Use:   "init-dump -o <output> -m <ncbi-metadata.tar.gz> -u <ncbi-uploads.gz>",
	Short: "Transform the NCBI Batch Telemetry tar files to a set of JSON Docs",
	Long: `This command transforms the raw NCBI Batch Telemetry tar file
         contents into a intermediary file of custom JSON Documents`,
	Run: func(cmd *cobra.Command, args []string) {
		InitDumpOpts.main()
	},
}

func init() {
	cmdInitDump.Flags().StringVarP(
		&InitDumpOpts.output,
		"output",
		"o",
		"2015-05-01-dump.dat.gz",
		"the output file to dump the serialized JSON Documents to",
	)

	cmdInitDump.Flags().StringVarP(
		&InitDumpOpts.metadata,
		"ncbi-metadata",
		"m",
		"NCBI_SRA_Metadata_Full_WUGSC_20150101.tar.gz",
		"the full NCBI SRA Metadata tar.gz file (NCBI_SRA_Metadata_Full_*.tar.gz)",
	)

	cmdInitDump.Flags().StringVarP(
		&InitDumpOpts.uploads,
		"ncbi-uploads",
		"u",
		"NCBI_SRA_Files_Full_WUGSC_20150501.gz",
		"the full NCBI SRA Uploads gzip file (NCBI_SRA_Files_Full_*.gz)",
	)
}

func (opts InitDumpCmdOpts) main() {
	InitDumpOpts.processOpts()
	utils.CheckFileExists(opts.metadata)
	utils.CheckFileExists(opts.uploads)
	initdump.Main(opts.metadata, opts.uploads, opts.output)
}

func (opts InitDumpCmdOpts) processOpts() {
	if opts.output == "" {
		log.Fatal(
			"Please supply an gzip output file to dump",
			"to via --output= !",
		)
	}

	if opts.metadata == "" {
		log.Fatal(
			"Please supply the NCBI SRA Metadata file",
			"via --metadata= !",
		)
	}

	if opts.uploads == "" {
		log.Fatal(
			"Please supply the NCBI SRA Uploads file",
			"via --uploads= !",
		)
	}
}
