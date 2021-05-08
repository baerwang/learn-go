package template_method

import "fmt"

type HttpDownloader struct {
	*template
}

func NewHttpDownloader() Downloader {
	downloader := &HttpDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HttpDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HttpDownloader) save() {
	fmt.Println("http save")
}
