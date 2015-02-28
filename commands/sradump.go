package commands

import (
	"github.com/indraniel/srasearch/sradump"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cmdTransform = &cobra.Command{
	Use:   "sra-dump [path/to/NCBIDownloadTarFile]",
	Short: "Transform the NCBI Batch Telemetry tar files to a set of JSON Docs",
	Long: `This command transforms the raw NCBI Batch Telemetry tar file
         contents into a intermediary file of custom JSON Documents`,
	Run: func(cmd *cobra.Command, args []string) {
		tarfile := getTarFile(args)
		checkTarExists(tarfile)
		makeSraDump(tarfile)
	},
}

func makeSraDump(tarfile string) {
	db := sradump.CollectAccessionStats(tarfile)
	sradump.ProcessTarXMLs(tarfile, db)
}

func getTarFile(args []string) (tarfile string) {
	if len(args) == 0 {
		log.Fatal("Please supply a tar file as an argument!")
	}

	tarfile = args[0]

	if tarfile == "" {
		log.Fatal("Please supply a tar file as an argument!")
	}

	return
}

func checkTarExists(tarfile string) {
	if _, err := os.Stat(tarfile); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			tarfile, err,
		)
	}
}
