package terraplug

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Downloader struct {
	PluginDirectory string
	Target          Target
}

func (d *Downloader) Download(url string, plugin PluginVersion) error {
	os.MkdirAll(d.PluginDirectory, 0700)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return errors.New("")
	}

	// extract tar.gz and zip, locate binary file

	file, err := os.Create(filepath.Join(
		d.PluginDirectory,
		plugin.Path(d.Target),
	))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)

	return err
}
