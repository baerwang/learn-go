package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {

	// NewHttpDownloader().Download("1.txt")

	NewFtpDownloader().Download("2.txt")
}
