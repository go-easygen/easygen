package main

import (
	"fmt"
	"testing"

	"github.com/danverbraganza/varcaser/varcaser"
)

////////////////////////////////////////////////////////////////////////////
// Lists

//==========================================================================
// list0

func TestList0(t *testing.T) {
	t.Log("First and plainest list test")
	const await = "The colors are: red, blue, white, .\n"
	if got := Generate(false, "Test/list0"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// list1

func TestList1Text(t *testing.T) {
	t.Log("Second test, with text template")
	const await = "The quoted colors are: \"red\", \"blue\", \"white\", .\n"
	if got := Generate(false, "Test/list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func TestList1HTML(t *testing.T) {
	t.Log("Second test, with html template")
	const await = "The quoted colors are: &#34;red&#34;, &#34;blue&#34;, &#34;white&#34;, .\n"
	if got := Generate(true, "Test/list1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

//==========================================================================
// listfunc1

func TestListFunc1(t *testing.T) {
	t.Log("Test custom template function - minus1")
	const await = "red, blue, white.\n"
	if got := Generate(false, "Test/listfunc1"); got != await {
		t.Errorf("Mismatch:, got '%s' instead", got)
	}
}

func ExampleFunc1() {
	fmt.Print(Generate(false, "Test/listfunc1"))
	// Output:
	// red, blue, white.
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
	fmt.Print(Generate(false, "Test/commandlineCV"))
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
	fmt.Print(Generate(false, "Test/commandlineCVFull"))
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
	fmt.Printf("%s %s %s %s",
		varcaser.Caser{From: varcaser.LowerCamelCase, To: varcaser.KebabCase}.
			String("someInitMethod"),
		varcaser.Caser{From: varcaser.LowerCamelCase,
			To: varcaser.ScreamingSnakeCase}.
			String("myConstantVariable"),
		ck2lc.String("some-init-method"),
		ck2uc.String("some-init-method"))
	// Output:
	// some-init-method MY_CONSTANT_VARIABLE someInitMethod SomeInitMethod
}
