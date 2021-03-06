package radius

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestRFCBuild(t *testing.T) {
	t.Parallel()

	var packages []string

	f, err := os.Open(".")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "rfc") {
			packages = append(packages, entry.Name())
		}
	}

	for _, pkg := range packages {
		func(pkg string) {
			t.Run(pkg, func(t *testing.T) {
				t.Parallel()

				cmd := exec.Command("go", "build", "layeh.com/radius/"+pkg)
				output, err := cmd.CombinedOutput()
				if err != nil {
					t.Errorf("%s: %s\n", err, output)
				}
			})
		}(pkg)
	}
}
