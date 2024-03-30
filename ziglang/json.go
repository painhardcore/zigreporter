// package ziglang provides a data structure for unmarshalling bizzare format zig versions from json
// that located here https://ziglang.org/download/index.json
package ziglang

import (
	"github.com/goccy/go-json"
)

type Platform struct {
	Tarball string `json:"tarball"`
	Shasum  string `json:"shasum"`
	Size    string `json:"size"`
}

type Version struct {
	Date      string              `json:"date"`
	Docs      string              `json:"docs"`
	StdDocs   string              `json:"stdDocs"`
	Notes     string              `json:"notes"`
	Version   string              `json:"version"`
	Platforms map[string]Platform `json:"platforms"`
}

type Versions map[string]Version

func (v *Version) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	platforms := make(map[string]Platform)
	for key, value := range raw {
		if key == "date" || key == "docs" || key == "stdDocs" || key == "notes" || key == "version" {
			var str string
			if err := json.Unmarshal(value, &str); err != nil {
				return err
			}
			switch key {
			case "version":
				v.Version = str
			case "date":
				v.Date = str
			case "docs":
				v.Docs = str
			case "stdDocs":
				v.StdDocs = str
			case "notes":
				v.Notes = str
			}
			continue
		}

		var platform Platform
		if err := json.Unmarshal(value, &platform); err != nil {
			return err
		}
		platforms[key] = platform
	}
	v.Platforms = platforms
	return nil
}
