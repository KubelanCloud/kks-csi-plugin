//go:build linux

package driver

import (
	"os"
	"os/exec"
	"testing"
)

func TestHostMountRequiresNsenter(t *testing.T) {
	if _, err := exec.LookPath("nsenter"); err != nil {
		t.Skip("nsenter not available")
	}
	if os.Geteuid() != 0 {
		t.Skip("host mount namespace tests require root")
	}

	tmp := t.TempDir()
	if out, err := hostMount("mkdir", "-p", tmp); err != nil {
		t.Fatalf("hostMount mkdir failed: %v: %s", err, string(out))
	}
	if _, err := os.Stat(tmp); err != nil {
		t.Fatalf("expected host-visible directory: %v", err)
	}
}
