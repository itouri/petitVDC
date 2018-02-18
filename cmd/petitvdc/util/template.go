package util

import (
	"encoding/json"
	"io"
	"os"

	"github.com/itouri/petitVDC/model"
)

// TODO この構造体の定義を別の所に移す
type TemplateRoot struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`

	RawTemplate json.RawMessage        `json:"template"`
	Template    model.ResourceTemplate `json:"-"`

	// TODO handler handler.ResourceHandler
}

func FetchTemplate(templateSlug string) (*model.ResourceTemplate, error) {
	// TODO あまりにも面倒なのでとりまハードコーディング
	file, err := open(templateSlug)
	if err != nil {
		return nil, err
	}
	return parse(file)
}

func open(templateName string) (*os.File, error) {
	return os.Open("template/centos/7/lxc.json")
}

func parse(reader io.Reader) (*model.ResourceTemplate, error) {
	parseRootTmpl()
	parseResourceTmpl()
}
