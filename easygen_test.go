package easygen

import (
	"fmt"
	"testing"

	"github.com/suntong/easygen"
)

////////////////////////////////////////////////////////////////////////////
// Lists

//==========================================================================
// list0

func TestList0(t *testing.T) {
	t.Log("First and plainest list test")
	const await = "The colors are: red, blue, white, .\n"
	if got := easygen.Generate(false, "test/list0"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// list1

func TestList1Text(t *testing.T) {
	t.Log("Second test, with text template")
	const await = "The quoted colors are: \"red\", \"blue\", \"white\", .\n"
	if got := easygen.Generate(false, "test/list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func TestList1HTML(t *testing.T) {
	t.Log("Second test, with html template")
	const await = "The quoted colors are: &#34;red&#34;, &#34;blue&#34;, &#34;white&#34;, .\n"
	if got := easygen.Generate(true, "test/list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// listfunc1

func TestListFunc1(t *testing.T) {
	t.Log("Test custom template function - minus1")
	const await = "red, blue, white.\n"
	if got := easygen.Generate(false, "test/listfunc1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func ExampleFunc1() {
	fmt.Print(easygen.Generate(false, "test/listfunc1"))
	// Output:
	// red, blue, white.
}

//==========================================================================
// list0 data + listfunc1 template
// I.e.: EasyGen -tf test/listfunc1 test/list0
func ExampleList0Func1() {
	fmt.Print(easygen.Generate2(false, "test/listfunc1", "test/list0"))
	// Output:
	// red, blue, white.
}

//==========================================================================
// list0 data + string template
// I.e.: EasyGen -ts "{{range .Colors}}{{.}}, {{end}}" test/list0
func ExampleList0StrTemplate() {
	fmt.Print(easygen.Generate0(false, "{{range .Colors}}{{.}}, {{end}}", "test/list0"))
	// Output:
	// red, blue, white,
}

//==========================================================================
// listfunc2

func ExampleFunc2() {
	fmt.Print(easygen.Generate(false, "test/listfunc2"))
	// Output:
	// some-init-method 5 5 someInitMethod SomeInitMethod
}

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func ExampleTestExample() {
	// Ref https://groups.google.com/d/msg/golang-nuts/9jKexxD19Js/xrNAQA1ACnMJ
	fmt.Println(`  flags.Bool("debug", false, "Turn on debugging.")`)
	fmt.Println(`  viper.BindPFlag("debug", flags.Lookup("debug"))`)
	// Output:
	//   flags.Bool("debug", false, "Turn on debugging.")
	//   viper.BindPFlag("debug", flags.Lookup("debug"))
}

func ExampleCommandLineCobraViper() {
	// EasyGen commandlineCV | sed 's|^\t|&//&|; s|^$|\t//|'
	fmt.Print(easygen.Generate(false, "test/commandlineCV"))
	// Output:
	//
	//	flags.Bool("debug", false, "Turn on debugging.")
	//	viper.BindPFlag("debug", flags.Lookup("debug"))
	//
	//	flags.String("addr", "localhost:5002", "Address of the service.")
	//	viper.BindPFlag("addr", flags.Lookup("addr"))
	//
	//	flags.String("smtp-addr", "localhost:25", "Address of the SMTP server.")
	//	viper.BindPFlag("smtp-addr", flags.Lookup("smtp-addr"))
	//
	//	flags.String("smtp-user", "", "User for the SMTP server.")
	//	viper.BindPFlag("smtp-user", flags.Lookup("smtp-user"))
	//
	//	flags.String("smtp-password", "", "Password for the SMTP server.")
	//	viper.BindPFlag("smtp-password", flags.Lookup("smtp-password"))
	//
	//	flags.String("email-from", "noreply@abc.com", "The from email address.")
	//	viper.BindPFlag("email-from", flags.Lookup("email-from"))
	//
}

func ExampleCommandLineOptInitFull() {
	// EasyGen commandlineCVFull | sed 's|^|\t// |;'
	fmt.Print(easygen.Generate(false, "test/commandlineCVFull"))
	// Output:
	// func init() {
	//
	//   viper.SetEnvPrefix("DISPATCH")
	//   viper.AutomaticEnv()
	//
	//   /*
	//
	//     When AutomaticEnv called, Viper will check for an environment variable any
	//     time a viper.Get request is made. It will apply the following rules. It
	//     will check for a environment variable with a name matching the key
	//     uppercased and prefixed with the EnvPrefix if set.
	//
	//   */
	//
	//   flags := mainCmd.Flags()
	//
	//   flags.Bool("debug", false, "Turn on debugging.")
	//   viper.BindPFlag("debug", flags.Lookup("debug"))
	//
	//   flags.String("addr", "localhost:5002", "Address of the service.")
	//   viper.BindPFlag("addr", flags.Lookup("addr"))
	//
	//   flags.String("smtp-addr", "localhost:25", "Address of the SMTP server.")
	//   viper.BindPFlag("smtp-addr", flags.Lookup("smtp-addr"))
	//
	//   flags.String("smtp-user", "", "User for the SMTP server.")
	//   viper.BindPFlag("smtp-user", flags.Lookup("smtp-user"))
	//
	//   flags.String("smtp-password", "", "Password for the SMTP server.")
	//   viper.BindPFlag("smtp-password", flags.Lookup("smtp-password"))
	//
	//   flags.String("email-from", "noreply@abc.com", "The from email address.")
	//   viper.BindPFlag("email-from", flags.Lookup("email-from"))
	//
	//   // Viper supports reading from yaml, toml and/or json files. Viper can
	//   // search multiple paths. Paths will be searched in the order they are
	//   // provided. Searches stopped once Config File found.
	//
	//   viper.SetConfigName("CommandLineCV") // name of config file (without extension)
	//
	//   viper.AddConfigPath("/tmp")
	//   viper.AddConfigPath(".")
	//
	//   err := viper.ReadInConfig()
	//   if err != nil {
	//     println("No config file found. Using built-in defaults.")
	//   }
	//
	// }
}

////////////////////////////////////////////////////////////////////////////
// Varible Names

func ExampleVaribleNames() {
}

////////////////////////////////////////////////////////////////////////////
// Strings Test

func ExampleTestStringsCmp() {
	fmt.Print(easygen.Generate0(false, `The {{if eq .StrTest "-AB-axxb- HTML Html html"}}eq says Yea{{else}}eq says Nay{{end}} but {{if eqf .StrTest "-AB-axxb- HTML Html html"}}eqf says Yea{{else}}eqf says Nay{{end}}.`, "test/strings0"))
	// Output:
	// The eq says Nay but eqf says Yea.
}

func ExampleTestStringSplit0() {
	fmt.Print(easygen.Generate0(false, `{{split .Colorlist}}`, "test/list0"))
	// Output:
	// [red blue white]
}

func ExampleTestStringSplit1() {
	fmt.Print(easygen.Generate0(false, `{{range (split .Colorlist)}}{{.}} {{end}}`, "test/list0"))
	// Output:
	// red blue white
}

/*

// Ref https://golang.org/pkg/regexp/#Regexp.ReplaceAllString
$ easygen -rf="a(x*)b" -rt='${1}W' -ts="{{.StrTest}}, {{replace .StrTest}}, {{.StrTest | replace}}" test/strings0
-ab-axxb- HTML Html html, -W-xxW- HTML Html html, -W-xxW- HTML Html html

$ easygen -rf="HTML" -rt='XML' -ts="{{.StrTest}}, {{replacec .StrTest}}, {{replace .StrTest}}" test/strings0
-ab-axxb- HTML Html html, -ab-axxb- XML Html html, -ab-axxb- XML XML XML

$ EASYGEN_RF="HTML" EASYGEN_RT='XML' easygen -ts="{{.StrTest}}, {{replacec .StrTest}}, {{replace .StrTest}}" test/strings0
-ab-axxb- HTML Html html, -ab-axxb- XML Html html, -ab-axxb- XML XML XML

*/

/*
func ExampleTestStrings() {
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	easygen.Opts.StrFrom = "a(x*)b"
	easygen.Opts.StrTo = "${1}W"
	//fmt.Print(easygen.Generate0(false, "{{.StrTest}} {{replace .StrTest}} {{.StrTest | replace}}", "test/strings0"))
	easygen.Opts.TemplateStr = `{{.StrTest}} {{replace .StrTest}} {{.StrTest | replace}}`
	fmt.Print(easygen.Generate(false, "test/strings0"))
	// Output:
	// -ab-axxb- HTML Html html -W-xxW- HTML Html html -W-xxW- HTML Html html
}
*/
