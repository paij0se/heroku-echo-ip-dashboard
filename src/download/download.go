package download

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	lmmp3 "github.com/paij0se/lmmp3"
)

// This function is going to download the statics files from github
func Download() {
	if _, err := os.Stat("herokudashboard"); os.IsNotExist(err) {
		lmmp3.DownloadFile("herokudashboard.zip", "https://raw.githubusercontent.com/paij0se/heroku-echo-ip-dashboard/main/herokudashboard.zip")
		//lmmp3.DownloadFile("herokudashboard.zip", "https://cdn.discordapp.com/attachments/950041049458438164/978495018467725312/herokudashboard.zip")
		// unzip the file
		dst := "herokudashboard"
		archive, err := zip.OpenReader("herokudashboard.zip")
		if err != nil {
			panic(err)
		}
		defer archive.Close()

		for _, f := range archive.File {
			filePath := filepath.Join(dst, f.Name)
			fmt.Println("unzipping file ", filePath)

			if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
				fmt.Println("invalid file path")
				return
			}
			if f.FileInfo().IsDir() {
				fmt.Println("creating directory...")
				os.MkdirAll(filePath, os.ModePerm)
				continue
			}

			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				panic(err)
			}

			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}

			fileInArchive, err := f.Open()
			if err != nil {
				panic(err)
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				panic(err)
			}

			dstFile.Close()
			fileInArchive.Close()
		}
	} else {
		fmt.Println("herokudashboard already exists")
	}
}
