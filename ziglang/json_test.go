package ziglang

import (
	"bytes"
	"os"
	"testing"

	"github.com/goccy/go-json"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalJSON(t *testing.T) {
	data, err := os.ReadFile("testdata/index.json")
	if err != nil {
		t.Fatal(err)
	}
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	var versions Versions
	if err := decoder.Decode(&versions); err != nil {
		t.Fatal(err)
	}

	// check master version
	assert.Equal(t, "0.12.0-dev.3496+a2df84d0f", versions["master"].Version)
	assert.Equal(t, "2024-03-30", versions["master"].Date)
	assert.Equal(t, "https://ziglang.org/documentation/master/", versions["master"].Docs)
	assert.Equal(t, "https://ziglang.org/documentation/master/std/", versions["master"].StdDocs)
	assert.Len(t, versions["master"].Platforms, 13)
	assert.Equal(t, versions["master"].Platforms["armv7a-linux"].Size, "42601528")

	// check some random values
	assert.Equal(t, "", versions["0.8.0"].Version)
	assert.Equal(t, "845cb17562978af0cf67e3993f4e33330525eaf01ead9386df9105111e3bc519", versions["0.7.1"].Platforms["x86_64-macos"].Shasum)
}

func BenchmarkUnmarshalVersions(b *testing.B) {
	data, err := os.ReadFile("testdata/index.json")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)
		decoder := json.NewDecoder(reader)
		var versions Versions
		if err := decoder.Decode(&versions); err != nil {
			b.Fatal(err)
		}
	}
}
