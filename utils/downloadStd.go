package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadStd() error {
	stdFile, err := os.Create("standard.txt")
	if err != nil {
		return err
	}
	defer stdFile.Close()

	stdDownload, err1 := http.Get("https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt")
	if err1 != nil{
		return err1
	}
	defer stdDownload.Body.Close()

	if stdDownload.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status %s", stdDownload.Status)
	}

	_,err2 := io.Copy(stdFile, stdDownload.Body)
	if err2 != nil{
		return err2
	}

	return nil

}