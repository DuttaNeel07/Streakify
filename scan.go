package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
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
	return folders
}
//This function is used to pass in an empty slice along with folder to start the scanGitFolders() function
func recursiveScanFolder(folder string) []string{
	return scanGitFolders(make ([]string, 0), folder)
}

//creating a function to return the location of dot file to get repo lists, if not present it creates the location for the user

func getDotFilePath() string{
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

func addNewSliceElementsToFile(filepath string, newRepos []string){
	existingRepos:= parseFileLinestoSlice(filepath)
	repos := joinSlices(newRepos, existingRepos)
}

func parseFileLinestoSlice(filepath string) []string{
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil{
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		if err != io.EOF{
			panic(err)
		}
	}

	return lines
}
//This function opens file located at filepath
func openFile(filepath string) *os.File{
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil{
		if os.IsNotExist(err){
			_, err := os.Create(filepath)
			if err != nil{
				panic(err)
			}
		} else{
			panic(err)
		}
	}

	return f
}

func joinSlices(new []string, existing []string) []string{
	for _,i := range new{
		if !sliceContains(existing, i){
			existing = append(existing, i)
		}
	}
return existing
}

func sliceContains(slice []string, value string) bool{
	for _,v := range slice{
		if v == value{
			return true
		}
	}
	return false
}

//This function takes the path and the slice and writes the slice on the file with each line onto a new line

func dumpStringsSlicetoFile(repos []string, filepath string){
	content := strings.Join(repos, "\n")
	os.WriteFile(filepath, []byte(content), 0755)
} 