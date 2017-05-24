package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

var (
	cmdEasygen = "easygen"
	dirTest    = "../../test/"
	extRef     = ".ref"
	extGot     = ".got"
)

func testEasygen(t *testing.T, name string, argv ...string) {
	cmd := exec.Command(cmdEasygen, argv...)

	// open the out file for writing
	outfile, err := os.Create(name + extGot)
	if err != nil {
		t.Errorf("write error [%s: %s] %s.", name, argv, err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error [%s: %s] %s.", name, argv, err)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("exit error [%s: %s] %s.", name, argv, err)
	}

	cmd = exec.Command("diff", "-U1", name+extRef, name+extGot)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error %s [%s: %s]", err, name, argv)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("cmp error %s [%s: %s]\n%s", err, name, argv, out)
	}

}

func TestExec(t *testing.T) {
	// var dir string

	// _, filename, _, _ := runtime.Caller(0)
	// fmt.Println("Current test filename: " + filename)
	// dir, _ = os.Getwd()
	// fmt.Println(dir)

	os.Chdir(dirTest)
	// dir, _ = os.Getwd()
	// fmt.Println(dir)

	//Test Functions
	testEasygen(t, "list0", "list0")
	testEasygen(t, "list0", "list0.yaml")
	testEasygen(t, "list0", "-tf", "list0", "list0")
	testEasygen(t, "list0", "-tf", "list0.tmpl", "list0")
	testEasygen(t, "list0", "-tf", "list0.tmpl", "list0.yaml")

	testEasygen(t, "list1", "list1")
	testEasygen(t, "listfunc1", "listfunc1")
	testEasygen(t, "listfunc2", "listfunc2")

	testEasygen(t, "commandlineCLI-024", "commandlineCLI-024")
	testEasygen(t, "commandlineCLI-027", "commandlineCLI-027")
	testEasygen(t, "commandlineCLI-027s", "-tf", "commandlineCLI-027", "commandlineCLI-027s")

	testEasygen(t, "commandlineCVFull", "commandlineCVFull")
	testEasygen(t, "commandlineCV", "commandlineCV")
	testEasygen(t, "commandlineFlag", "commandlineFlag")

}
