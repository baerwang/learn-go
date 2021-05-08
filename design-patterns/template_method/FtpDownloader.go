package template_method

import "fmt"

type FtpDownloader struct {
	*template
}

func NewFtpDownloader() Downloader {
	downloader := &FtpDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FtpDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}
