package ziptools

import (
	"archive/zip"
	"log"
)


func GetInnerFilesModTime(pathZipFile, delimiter string) string {
	
		
	
		// Open a zip archive for reading.
	
		r, err := zip.OpenReader(pathZipFile)
	
		if err != nil {
	
			log.Fatal(err)
	
		}
	
		defer r.Close()
	
		// Iterate through the files in the archive,
	
		// printing some of their contents.
	
		result := delimiter + delimiter + pathZipFile + delimiter
	
		for _, f := range r.File {
	
			name := " " + f.Name
			t := f.ModTime()
			last_time := t.Format("02-01-2006 15:04:05")
			last_string := "       " + last_time + delimiter
	
			result = result + name + last_string
	
		}
	
		return result
	}