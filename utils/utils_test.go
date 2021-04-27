package utils

import "testing"

func TestOperating(t *testing.T) {
	PackageZip(`/baerwang/wxx`, `haha.gz`)
	UnZip(`/baerwang/wxx/utils/haha.gz`)
}

func TestAbc(t *testing.T) {
	PrintAbc()
}
