package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)
	
	var err error
	//** Use valid paths according to your OS. **//
 
    // CREATING A FILE
 
    // os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
    // it returns a file descriptor which is a pointer to os.File and an error value.
	newFile, err = os.Create("a.txt")

	// error handling
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors
	}

	// TRUNCATING A FILE
	err = os.Truncate("a.txt", 0) // clean bytes in file  // 0 means completely empty the file.
  	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// CLOSING THE FILE
	newFile.Close()

    // OPEN AND CLOSE AN EXISTING FILE
	file, err := os.Open("a.txt") // open in read-only mode

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
   	//OPENING a FILE WITH MORE OPTIONS
	file, err = os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	// We can Use opening attributes individually or combined
    // using an OR between them
    // e.g. os.O_CREATE|os.O_APPEND
    // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
    // os.O_RDONLY // Read only
    // os.O_WRONLY // Write only
    // os.O_RDWR // Read and write
    // os.O_APPEND // Append to end of file
    // os.O_CREATE // Create is none exist
    // os.O_TRUNC // Truncate file when opening

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	/** STATS FILE INFO **/
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println

	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is Directory? ", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat("b.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			// log.Fatal("File does not exist!!")
		}
	}
	// RENAMING AND MOVING A FILE
	oldPath := "a.txt"
	newPath := "aaa.txt"

	err = os.Rename(oldPath, newPath)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	
    // REMOVING A FILE
	err = os.Remove(newPath)
 	// error handling
	if err != nil {
		log.Fatal(err)
	}
}