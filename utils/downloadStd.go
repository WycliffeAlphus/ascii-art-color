package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
This DownloadStd function downloads a file from a specified URL and saves it as "standard.txt" in the local file system.
It handles errors that may occur during the process of creating the file, downloading the content,
checking the HTTP response status, and copying the content to the file.
*/
func DownloadStd() error {
	stdFile, err := os.Create("standard.txt")
	if err != nil {
		return err
	}
	defer stdFile.Close()

	stdDownload, err1 := http.Get("https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt")
	if err1 != nil {
		return err1
	}
	defer stdDownload.Body.Close()

	if stdDownload.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status %s", stdDownload.Status)
	}

	_, err2 := io.Copy(stdFile, stdDownload.Body)
	if err2 != nil {
		return err2
	}

	return nil

}
