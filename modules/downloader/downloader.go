package downloader

import (
	"os"
	"net/http"
	"io"
)

type Downloader struct {

}


func (dl *Downloader) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func Download(dl* Downloader, url string, filepath string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &Downloader{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}