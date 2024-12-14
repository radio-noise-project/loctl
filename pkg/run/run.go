package run

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/radio-noise-project/loctl/internal/client"
)

func StartContainer(sisterId string, dirname string) {
	cli := client.NewConfiguration()
	sendTarball(sisterId, dirname, cli)
}

func sendTarball(sisterId string, dirname string, cli *client.Config) {
	var buf bytes.Buffer
	var sendBuf bytes.Buffer
	gzw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gzw)

	err := filepath.Walk(dirname, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(file)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if fi.Mode().IsRegular() {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.Copy(tw, f); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if err := tw.Close(); err != nil {
		panic(err)
	}

	if err := gzw.Close(); err != nil {
		panic(err)
	}

	writer := multipart.NewWriter(&sendBuf)
	filewriter, err := writer.CreateFormFile("file", "test.tar.gz")
	if err != nil {
		panic(err)
	}

	io.Copy(filewriter, &buf)
	writer.Close()

	resp, err := http.Post(cli.DestinationUrl+"/api/v0/run?sisterId="+sisterId, writer.FormDataContentType(), &sendBuf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("failed to upload tarball")
	}
}
