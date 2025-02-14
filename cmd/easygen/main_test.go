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

	t.Logf("Testing %s:\n\t%s %s", name, cmdEasygen, strings.Join(argv, " "))

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
	//os.Remove(generatedOutput)
}

func TestExec(t *testing.T) {
	os.Chdir(dirTest)

	//Test Basic Functions
	testEasygen(t, "list0", "list0")
	// Filename suffixes are optional
	testEasygen(t, "list0", "list0.yaml")
	testEasygen(t, "list0", "list0", "list0")
	testEasygen(t, "list0", "list0.tmpl", "list0")
	testEasygen(t, "list0", "list0.tmpl", "list0.yaml")
	//testEasygen(t, "list0E", "list0E", "list0")

	testEasygen(t, "list1", "list1")
	testEasygen(t, "listfunc1", "listfunc1")
	testEasygen(t, "listfunc2", "listfunc2")
	testEasygen(t, "tf-calc", "tf-calc")

	//Test Basic Json Functions
	testEasygen(t, "list0j", "list0j")
	testEasygen(t, "list0j", "list0j", "list0j")
	testEasygen(t, "list0j", "list0j", "list0j.json")

	// template_name can be a comma-separated list
	testEasygen(t, "list01", "list0,listfunc1", "list0")

	//Test HTML
	testEasygen(t, "list1HTML", "list1HTML", "list1")

	//Test lowercase keys
	testEasygen(t, "example-lowercase", "example-lowercase")

	// varcaser string functions
	testEasygen(t, "var0", "-ts", "{{.Name}}", "var0")
	testEasygen(t, "var1", "-ts", "{{clk2uc .Name}}", "var0")
	testEasygen(t, "var2", "-ts", "{{clk2ss .Name}}", "var0")

	//Test Bigger files
	testEasygen(t, "commandlineCLI-024", "commandlineCLI-024")
	testEasygen(t, "commandlineCLI-027", "commandlineCLI-027")
	testEasygen(t, "commandlineCLI-027s", "commandlineCLI-027", "commandlineCLI-027s")

	testEasygen(t, "commandlineCVFull", "commandlineCVFull")
	testEasygen(t, "commandlineCV", "commandlineCV")
	testEasygen(t, "commandlineFlag", "commandlineFlag")

	// Enum generation: (a) run template with multiple data inputs,
	//                  (b) run the same input with multiple template files:
	testEasygen(t, "enum_multiple_data_files", "enum_c-header", "raid_type", "raid_driver")
	testEasygen(t, "enum_multiple_template_files", "enum_c-header,enum_c-source", "raid_type.yaml")
	testEasygen(t, "enum_multiple_template_and_data", "enum_c-header,enum_c-to_str", "raid_type", "raid_driver.yaml")

	// Test with multiple data inputs files:
	testEasygen(t, "lists", "lists", "list0.yaml,listfunc2.yaml")
	testEasygen(t, "lists", "lists", "list0,listfunc2")
	testEasygen(t, "lists", "lists", "listfunc2.yaml,list0.yaml")

	//Test nested templates
	//testEasygen(t, "nested_header_footer", "nested_header.tmpl,nested_footer.tmpl,nested_thanks.tmpl", "nested_data.yaml")
	testEasygen(t, "nested_demo_argsa", "nested_demo_argsa", "nested_data.yaml")
	testEasygen(t, "nested_demo_argsm", "nested_demo_argsm", "nested_data.yaml")
	testEasygen(t, "nested_demo_argsm_iterate", "nested_demo_argsm_iterate", "nested_data.yaml")
}
