package device

import (
	"path/filepath"

	. "github.com/namkazt/devicedetector/parser"
)

const ParserNameNotebook = `notebook`
const FixtureFileNotebook = `notebooks.yml`

func init() {
	RegDeviceParser(ParserNameNotebook,
		func(dir string) DeviceParser {
			return NewNotebook(filepath.Join(dir, FixtureFileNotebook))
		})
}

func NewNotebook(fileName string) *Notebook {
	h := &Notebook{}
	if err := h.Load(fileName); err != nil {
		h.notebookRegx.Regex = `FBMD/`
		return nil
	}
	return h
}

// Device parser for hbbtv detection
type Notebook struct {
	DeviceParserAbstract
	notebookRegx Regular
}

func (h *Notebook) Parse(ua string) *DeviceMatchResult {
	// only parse user agents containing FBMD fragment
	if !h.notebookRegx.IsMatchUserAgent(ua) {
		return nil
	}
	return h.DeviceParserAbstract.Parse(ua)
}
