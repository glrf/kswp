package kswp

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSwap(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestSwap")
	if err != nil {
		t.Fatalf("Could not create temp dir: %s", err)
	}
	defer os.RemoveAll(dir) // clean up

	k := Kswp{
		KubeConf: filepath.Join(dir, "conf"),
		Configs: []KubeConfig{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "buzz"},
			{Name: "none", Path: "/tmp/does/not/exist"},
		},
	}

	for i, c := range k.Configs {
		if c.Path == "" {
			tmpfn := filepath.Join(dir, c.Name)
			if err := ioutil.WriteFile(tmpfn, []byte(c.Name), 0666); err != nil {
				t.Fatalf("Could not create temp file: %s", err)
			}
			k.Configs[i].Path = tmpfn
		}
	}

	for _, c := range k.Configs[0:3] {
		if err := k.Swap(c.Name); err != nil {
			t.Errorf("Swap for %s failed with \n%s", c.Name, err.Error())
		} else {
			file, err := ioutil.ReadFile(k.KubeConf)
			if err != nil {
				t.Errorf("Could not read config file: %s", err)
			} else if string(file) != c.Name {
				t.Errorf("Config not written. Expected %s got %s", c.Name, string(file))
			}
		}
	}

	if err := k.Swap(k.Configs[3].Name); err == nil {
		t.Errorf("Swap for %s should have failed", k.Configs[3].Name)
	}

	if err := k.Swap("wrong name"); err == nil {
		t.Errorf("Swap for %s should have failed", "wrong name")
	}
}
