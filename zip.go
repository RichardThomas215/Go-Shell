package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
)

func zipUp(filesToArchive []string) {

	// Create a file to write the archive buffer to
	outFile, err := os.Create("files.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Create a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outFile)

	// Create and write files to the archive
	for i := 1; i < len(filesToArchive); i++ {

		fileWriter, err := zipWriter.Create(filesToArchive[i])
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadFile(filesToArchive[i])

		_, err = fileWriter.Write([]byte(data))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Clean up
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
