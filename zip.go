package main

import (
    "archive/zip"
    "log"
    "os"
)

func zipUp(filesToArchive  [] string) {
    // Create a file to write the archive buffer to
    // Could also use an in memory buffer.
    outFile, err := os.Create("test.zip")
    if err != nil {
        log.Fatal(err)
    }
    defer outFile.Close()

    // Create a zip writer on top of the file writer
    zipWriter := zip.NewWriter(outFile)

	// Create and write files to the archive, which in turn
    // are getting written to the underlying writer to the
    // .zip file we created at the beginning
    for i := 1 ; i< len(filesToArchive); i++ {

            _, err := zipWriter.Create(filesToArchive[i])
            if err != nil {
                    log.Fatal(err)
			}
			
            //   _, err = fileWriter.Write([]byte(file.Body))
            //   if err != nil {
            //           log.Fatal(err)
            //   }
    }

    // Clean up
    err = zipWriter.Close()
    if err != nil {
            log.Fatal(err)
    }
}