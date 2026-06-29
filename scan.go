package main

import (
	"fmt"
	"log"
	"strings"
	"os"
)

func scan(folder string){
	fmt.Printf("Folders found: \n\n")
	repositiories:= recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewSliceElementsToFile(filepath, repositiories)
	fmt.Printf("\n\nSucessfully Added!\n\n")
}
//This is a recursive function to get into all directories to find .git folder
func scanGitFolders(folders []string, folder string) []string{
	// to trim the last '/' in folder path
	folder = strings.TrimSuffix(folder, "/")

	f,err := os.Open(folder)
	if err != nil{
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil{
		log.Fatal(err)
	}

	var path string
	//To go through all subdirs and add .git folders path and skip node_modules folder as they are too large also not meaningful for this
	for _, file:= range files{
		if file.IsDir(){
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "./git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}else if file.Name() == "node_modules"{
				continue
			}

			folders = scanGitFolders(folders, path)
			
		}
	}
}
//This function is used to pass in an empty slice along with folder to start the scanGitFolders() function
func recursiveScanFolder(folder string) []string{
	return scanGitFolders(make ([]string, 0), folder)
}

