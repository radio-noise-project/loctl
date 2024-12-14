package run

import (
	"archive/tar"
	"bytes"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/facebookgo/symwalk"
	"github.com/radio-noise-project/loctl/internal/client"
)

func StartContainer(sisterId string, dirname string) {
	cli := client.NewConfiguration()
	sendTarball(sisterId, dirname, cli)
}

func sendTarball(sisterId string, dirname string, cli *client.Config) {
	buf := new(bytes.Buffer)
	var sendBuf bytes.Buffer
	tw := tar.NewWriter(buf)
	defer tw.Close()

	if err := symwalk.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		intake_path := strings.Replace(path, dirname, "", 1)
		tar_mode := info.Mode()
		tar_size := info.Size()

		if info.Mode()&fs.ModeSymlink != 0 {
			path_sym, err := filepath.EvalSymlinks(path)
			if err != nil {
				return err
			}
			path = path_sym
			stat, err := os.Stat(path_sym)
			if err != nil {
				return err
			}
			tar_mode = stat.Mode()
			tar_size = stat.Size()
		}

		if err := tw.WriteHeader(&tar.Header{
			Name:    intake_path,
			Mode:    int64(tar_mode),
			ModTime: info.ModTime(),
			Size:    tar_size,
		}); err != nil {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		return nil

	}); err != nil {
		panic(err)
	}

	writer := multipart.NewWriter(&sendBuf)
	filewriter, err := writer.CreateFormFile("file", "file")
	if err != nil {
		panic(err)
	}

	io.Copy(filewriter, buf)
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
