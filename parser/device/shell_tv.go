package device

import (
	"path/filepath"

	. "github.com/namkazt/devicedetector/parser"
)

const ParserNameShellTV = `shelltv`
const FixtureFileShellTV = `shell_tv.yml`

func init() {
	RegDeviceParser(ParserNameShellTV,
		func(dir string) DeviceParser {
			return NewShellTV(filepath.Join(dir, FixtureFileShellTV))
		})
}

func NewShellTV(fileName string) *Notebook {
	h := &Notebook{}
	if err := h.Load(fileName); err != nil {
		h.notebookRegx.Regex = `[a-z]+[ _]Shell[ _]\w{6}|tclwebkit(\d+[\.\d]*)`
		return nil
	}
	return h
}

// Device parser for hbbtv detection
type ShellTV struct {
	DeviceParserAbstract
	shellTVRegx Regular
}

func (h *ShellTV) Parse(ua string) *DeviceMatchResult {
	// only parse user agents containing fragments: {brand} shell
	if !h.IsShellTv(ua) {
		return nil
	}
	r := h.DeviceParserAbstract.Parse(ua)
	// always set device type to tv, even if no model/brand could be found
	if r != nil {
		r.Type = ParserNameHbbTv
	}
	return r
}

// Returns if the parsed UA was identified as a HbbTV device
func (h *ShellTV) IsShellTv(ua string) bool {
	return h.shellTVRegx.IsMatchUserAgent(ua)
}
