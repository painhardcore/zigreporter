package ziglang

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

const ZigLangVerionsURL = "https://ziglang.org/download/index.json"

// GetLatestMasteVersion for those who want to know the latest version of Zig
func GetLatestMasteVersion() (string, error) {
	resp, err := http.Get(ZigLangVerionsURL)
	if err != nil {
		return "", fmt.Errorf("Failed to fetch %s: %w", ZigLangVerionsURL, err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var versionData Versions
	if err := decoder.Decode(&versionData); err != nil {
		return "", fmt.Errorf("Failed to decode JSON with versions: %w", err)
	}
	latest, ok := versionData["master"]
	if !ok {
		return "", fmt.Errorf("Failed to find 'master' version in JSON, %v", versionData)
	}
	return latest.Version, nil
}
