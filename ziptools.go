package ziptools

import (
	"archive/zip"
	"log"
)

//Выдает список файлов (ввиде одной строки) из архива + последнее время изменение этих файлов . 
// можно указать разделитель \n - для строки или <br> - для HTML формата. 
func GetInnerFilesModTime(pathZipFile, delimiter string) string {
	
		

		// Open a zip archive for reading.
		r, err := zip.OpenReader(pathZipFile)
	
		if err != nil {
	
			log.Fatal(err)
	
		}
	
		defer r.Close()
	
		//создаем шаблонную строку для вывода
		result := delimiter + delimiter + pathZipFile + delimiter
	
		//пробигаем по списку файлов архива и добавляем информацию в шаблонную строку
		for _, f := range r.File {
	
			name := " " + f.Name
			t := f.ModTime()
			last_time := t.Format("02-01-2006 15:04:05")
			last_string := "       " + last_time + delimiter
	
			result = result + name + last_string
	
		}
	
		return result
	}