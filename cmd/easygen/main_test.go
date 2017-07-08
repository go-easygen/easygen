package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	cmdEasygen = "easygen"
	dirTest    = "../../test/"
	extRef     = ".ref" // extension for reference file
	extGot     = ".got" // extension for generated file
)

// testEasygen runs @cmdEasyGen with @argv and compares the generated
// output for @name with the corresponding @extRef
func testEasygen(t *testing.T, name string, argv ...string) {
	var (
		diffOut         bytes.Buffer
		generatedOutput = name + extGot
		cmd             = exec.Command(cmdEasygen, argv...)
	)

	t.Logf("Testing %s: `%s %s`", name, cmdEasygen, strings.Join(argv, " "))

	// open the out file for writing
	outfile, err := os.Create(generatedOutput)
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

	cmd = exec.Command("diff", "-U1", name+extRef, generatedOutput)
	cmd.Stdout = &diffOut

	err = cmd.Start()
	if err != nil {
		t.Errorf("start error %s [%s: %s]", err, name, argv)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("cmp error %s [%s: %s]\n%s", err, name, argv, diffOut.String())
	}
	os.Remove(generatedOutput)
}

func TestExec(t *testing.T) {
	os.Chdir(dirTest)

	//Test Basic Functions
	testEasygen(t, "list0", "list0")
	// Filename suffixes are optional
	testEasygen(t, "list0", "list0.yaml")
	testEasygen(t, "list0", "-tf", "list0", "list0")
	testEasygen(t, "list0", "-tf", "list0.tmpl", "list0")
	testEasygen(t, "list0", "-tf", "list0.tmpl", "list0.yaml")

	testEasygen(t, "list1", "list1")
	testEasygen(t, "listfunc1", "listfunc1")
	// TODO: fix all //XX: lines
	//XX: testEasygen(t, "listfunc2", "listfunc2")

	//Test String Functions
	testEasygen(t, "strings0", "-rf", `a(x*)b`, "-rt", `${1}W`, "strings0")
	testEasygen(t, "strings1", "-rf", "HTML", "-rt", "XML", "-tf", "strings1", "strings0")
	// varcaser string functions
	testEasygen(t, "var0", "-ts", "{{.Name}}", "var0")
	//XX: testEasygen(t, "var1", "-ts", "{{ck2uc .Name}}", "var0")
	//XX: testEasygen(t, "var2", "-ts", "{{ck2ss .Name}}", "var0")

	//Test Bigger files
	testEasygen(t, "commandlineCLI-024", "commandlineCLI-024")
	testEasygen(t, "commandlineCLI-027", "commandlineCLI-027")
	testEasygen(t, "commandlineCLI-027s", "-tf", "commandlineCLI-027", "commandlineCLI-027s")

	testEasygen(t, "commandlineCVFull", "commandlineCVFull")
	testEasygen(t, "commandlineCV", "commandlineCV")
	//XX: testEasygen(t, "commandlineFlag", "commandlineFlag")

	// Enum generation: (a) run template with multiple data inputs,
	//                  (b) run the same input with multiple template files:
	testEasygen(t, "enum_multiple_data_files", "-tf", "enum_c-header", "raid_type", "raid_driver")
	testEasygen(t, "enum_multiple_template_files", "-tf", "enum_c-header,enum_c-source", "raid_type.yaml")
	testEasygen(t, "enum_multiple_template_and_data", "-tf", "enum_c-header,enum_c-to_str", "raid_type", "raid_driver.yaml")
}
