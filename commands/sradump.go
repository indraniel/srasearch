package commands

import (
	"github.com/indraniel/srasearch/jdoc"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cmdTransform = &cobra.Command{
	Use:   "transform [path/to/NCBIDownloadTarFile]",
	Short: "Transform the NCBI Batch Telemetry tar file to a JSON Document",
	Long: `This command transforms the raw NCBI Batch Telemetry tar file
         contents into a file of custom JSON Documents`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here I am in the 'transform' command!")
		tarfile := getTarFile(args)
		checkTarExists(tarfile)
		jdoc.ProcessNCBITarFile(tarfile)
	},
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
