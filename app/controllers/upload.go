package controllers

import (
	"github.com/revel/revel"
)

type Upload struct {
	*revel.Controller
}

type FileInfo struct {
	ContentType string
	Filename    string
	RealFormat  string `json:",omitempty"`
	Resolution  string `json:",omitempty"`
	Size        int
	Status      string `json:",omitempty"`
}

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func (c *Upload) Upload() revel.Result {
	return c.Render()
}

func (c *Upload) HandleUpload(bundle []byte) revel.Result {
	// Validation rules.
	c.Validation.Required(bundle)
	c.Validation.MaxSize(bundle, 2*MB).
		Message("File cannot be larger than 2MB")

	// Check format of the file.
	//conf, format, err := image.DecodeConfig(bytes.NewReader(avatar))

	return c.RenderJson(FileInfo{
		ContentType: c.Params.Files["bundle"][0].Header.Get("Content-Type"),
		Filename:    c.Params.Files["bundle"][0].Filename,
		Size:        len(bundle),
		Status:      "Successfully uploaded",
	})
}
