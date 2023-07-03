package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"sync"
)

func crackZipPassword(zipFile, password string, wg *sync.WaitGroup) {
	defer wg.Done()

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Printf("Error opening zip file: %s\n", err)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		f.SetPassword(password)
		rc, err := f.Open()
		if err == nil {
			// Password successfully cracked!
			fmt.Printf("Password found: %s\n", password)
			defer rc.Close()
			break
		}
	}
}

func main() {
	zipFile := "file.zip" // Replace with the name of your zip file
	passwords := []string{
		"password1",
		"password2",
		"password3",
	} // Add more passwords to try

	var wg sync.WaitGroup
	for _, password := range passwords {
		wg.Add(1)
		go crackZipPassword(zipFile, password, &wg)
	}

	wg.Wait()
	fmt.Println("Password cracking completed.")
}

