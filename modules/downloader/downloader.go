package downloader

import (
	"os"
	"net/http"
	"sync"
	"../fs"
)

type Downloader struct {
	wg sync.WaitGroup
	filePath string
	url string
	fileSize int64
	downloadedBytes int64
}

func (dl* Downloader) dl(start int64, end int32) error {
	client := &http.Client{
		CheckRedirect: nil,
	}
	req, err := http.NewRequest("GET", dl.url, nil)
	req.Header.Add("Range", `bytes=%d-%d`)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, _ := os.Open(dl.filePath + ".tmp")
	file.Seek(start, 0)

	buff := make([]byte, 4096)
	resp.Body.Read(buff)
	file.Write(buff)

	if dl.downloadedBytes == dl.fileSize {
		err = os.Rename(dl.filePath+".tmp", dl.filePath)
		if err != nil {
			return err
		}
	}
	defer dl.wg.Done()

	return nil
}

func (dl* Downloader) Download(url string, filePath string, parts int) error {
	if parts > 0 && parts < 8 {
		dl.url = url
		dl.filePath = filePath
		dl.wg.Add(parts)

		resp, err := http.Get(dl.url)
		if err != nil {
			return err
		}

		dl.fileSize = resp.ContentLength

		fs := fs.Filesystem{}
		fs.CreateEmptyFile(dl.filePath, dl.fileSize)

		go dl.dl(0, 100)
	}
}