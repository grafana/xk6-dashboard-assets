package assets_test

import (
	"encoding/json"
	"io/fs"
	"regexp"
	"testing"

	assets "github.com/grafana/xk6-dashboard-assets"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	raw := assets.Config()
	if len(raw) == 0 {
		t.Fatal("expected non-empty config")
	}

	var cfg map[string]any

	err := json.Unmarshal(raw, &cfg)
	if err != nil {
		t.Fatalf("failed to unmarshal config: %v", err)
	}
}

func TestReport(t *testing.T) {
	t.Parallel()

	report := assets.Report()
	if len(report) == 0 {
		t.Fatal("expected non-empty report")
	}

	htmlPattern := regexp.MustCompile(`(?is)<html[^>]*>.*</html>`)
	if !htmlPattern.Match(report) {
		t.Fatal("expected report to contain valid HTML with opening and closing <html> tags")
	}
}

func TestUI(t *testing.T) {
	t.Parallel()

	ui := assets.UI()
	entries, err := fs.ReadDir(ui, ".")
	if err != nil {
		t.Fatalf("failed to read UI directory: %v", err)
	}

	if len(entries) == 0 {
		t.Fatal("expected non-empty UI directory")
	}

	// Check for the presence of index.html
	_, err = fs.ReadFile(ui, "index.html")
	if err != nil {
		t.Fatal("expected index.html to be present in UI assets")
	}
}
