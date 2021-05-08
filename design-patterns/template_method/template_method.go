package template_method

import "fmt"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Println("prepare downloading")
	t.implement.download()
	t.implement.save()
	fmt.Println("finish downloading")
}

func (t *template) save() {
	fmt.Println("default save")
}
