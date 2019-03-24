# easygen

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-easygen/easygen?status.svg)](http://godoc.org/github.com/go-easygen/easygen)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-easygen/easygen)](https://goreportcard.com/report/github.com/go-easygen/easygen)
[![travis Status](https://travis-ci.org/go-easygen/easygen.svg?branch=master)](https://travis-ci.org/go-easygen/easygen)
[ ![Codeship Status for go-easygen/easygen](https://codeship.com/projects/4f9d9b30-b8ad-0133-b733-0e8881fc1b37/status?branch=master)](https://codeship.com/projects/135255)

## TOC
- [easygen - Easy to use universal code/text generator](#easygen---easy-to-use-universal-codetext-generator)
  - [Usage](#usage)
    - [$ easygen](#-easygen)
  - [Install](#install)
    - [Install as Debian/Unbuntu package](#install-as-debianunbuntu-package)
    - [Install from source](#install-from-source)
  - [Test](#test)
  - [Details](#details)
  - [Command line flag handling code auto-generation](#command-line-flag-handling-code-auto-generation)
    - [`easygen` itself](#`easygen`-itself)
    - [`cli` based](#`cli`-based)
  - [The `easygen` usage](#the-`easygen`-usage)
    - [Command line](#command-line)
    - [The library ](#the-library-)
      - [> cmd/easygen/main.go](#-cmdeasygenmaingo)
    - [Different Versions](#different-versions)
  - [Author(s) & Contributor(s)](#author(s)-&-contributor(s))

# easygen - Easy to use universal code/text generator

Command `easygen` is an easy to use universal code/text generator.

It can be used as a text or html generator for _arbitrary_ purposes with _arbitrary_ data and templates. It is a good [GSL](https://github.com/imatix/gsl) replacement, as it

  - is more easy to define driving data, in form of YML instead of XML
  - has more powerful template engine that based on Go template.
    You can also write your own function in Go to customize your template.

You can even use easygen as a generic Go template testing tool using the `-ts` commandline option, and much more.

Note this document is for `easygen` versions 4.0+. For historic versions check out the [Different Versions](#different-versions) section.


## Usage

### $ easygen
```sh
easygen version 4.0.0

Usage:
 easygen [flags] template_name [data_filename [data_filename...]]

Flags:

  -debug level
    	debugging level
  -ej extension
    	extension of json file (default ".json")
  -et extension
    	extension of template file (default ".tmpl")
  -ey extension
    	extension of yaml file (default ".yaml")
  -ts string
    	template string (in text)

data_filename(s): The name for the .yaml or .json data.
 - If omitted derive from the template_name.
 - Can have the extension or without it. If withot extension,
   will try .yaml first then .json
 - Can include the path as well.

template_name: The name for the template file.
 - Can have the extension or without it.
 - Can include the path as well.
 - Can be a comma-separated list giving many template files, in which case
   at least one data_filename must be given.

Flag defaults can be overridden by corresponding environment variable, e.g.:
  EASYGEN_EY=.yml EASYGEN_ET=.tpl easygen ...
```

## Install

### Install as Debian/Unbuntu package


    apt install easygen


### Install from source

	go get github.com/go-easygen/easygen/...
	ls -l $GOPATH/bin

You should find an `easygen` executable newly created in there. 

## Test

	export PATH=$PATH:$GOPATH/bin

	$ easygen $GOPATH/src/github.com/go-easygen/easygen/test/list0
	The colors are: red, blue, white, .

	cd $GOPATH/src/github.com/go-easygen/easygen

	$ easygen test/list1 
	The quoted colors are: "red", "blue", "white", .

	$ easygen test/listfunc1 test/list0
	red, blue, white.


And also check out the provided [more examples](https://godoc.org/github.com/go-easygen/easygen#pkg-examples) in the [![Go Doc](https://img.shields.io/badge/godoc-reference-4b68a3.svg)](https://godoc.org/github.com/go-easygen/easygen) document.


## Details

It can be used as a code generator, for example, command line parameter handling code generator, or anything that is structurally repetitive, like the following:


- [Introduction to easygen and its philosophy ](https://suntong.github.io/blogs/2016/01/01/easygen---easy-to-use-universal-code/text-generator)
- [Easygen is now coding itself ](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/)
- [Showcasing the power of easygen with ffcvt ](https://sfxpt.wordpress.com/2015/08/02/showcasing-the-power-of-easygen-with-ffcvt/)
- [Easygen for HTML mock-up ](https://sfxpt.wordpress.com/2015/07/10/easygen-for-mock-up/)
- [Moving beyond code-gen and mock-up, using easygen in real life creating GPT partitions](https://suntong.github.io/blogs/2015/12/26/creating-gpt-partitions-easily-on-the-command-line)

Ready to get started? Then check out [Getting Started](https://github.com/go-easygen/easygen/wiki/Getting-Started) to start building your way to turn your data into any form, any way you want.

<a name="clfhcag" />

## Command line flag handling code auto-generation

[ ](https://suntong.github.io/blogs/)

As explained above, one practical use of `easygen` is to auto-generating Go code for command line parameter handling, for both [`viper` and `cobra`](https://github.com/suntong/blog/blob/master/GoOptP7-easygen.md), and Go's [built-in `flag` package](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/).


### `easygen` itself

Currently, `easygen`'s command line parameter handling is built on top of Go's built-in `flag` package, and the handling code is entirely generated by `easygen` itself. Thus, showing how `easygen` is handling the command line parameters itself also means showing what functionality the auto-generated command line parameter handling code can do for you. 

Currently, there are three tiers program parameters can be given:

1. Default values defined within the program, so that program parameters can have meaningful defaults to start with
1. Values defined in environment variables
1. Values passed from command line 

The latter will have higher priority and will override values defined formerly. I.e., the values from command line will override that in environment variables, which in turn override program defaults.

We will use the `-ts`, template string, as an example to illustrate. The program defaults is empty, which means using the `.tmpl` template file the same as the `.yaml` data file. We will override that first by environment variable, then from command line, illustrated in next section.

### `cli` based

See,

[A cookbook on how to jump-start a `cli` based command line handling program](cli-project.md#cookbook)

## The `easygen` usage

### Command line

[Check here for more on using `easygen` the command line tool](using_easygen.md).

### The library 

The `easygen` is a library as well as a command line tools. Not only it is super easy to use, it is super easy to extend as well.

The [restructured `easygen`](https://github.com/go-easygen/easygen/issues/10) can now be a building block that people can easily extend, any extra functionalities, or extra feature that it depends on, or any external dependencies are now moved out to sub modules. Thus the library users can now pick and choose exactly what they want from the library.

- The [egVar package example](http://godoc.org/github.com/go-easygen/easygen/egVar#example-package) shows how to add the variable name manipulation on top of the default library.
- The [egCal  package example](http://godoc.org/github.com/go-easygen/easygen/egCal#example-package) shows how to add the variable name manipulation and generic calculation functionalities, together with the default functions, all at the same time.

To put them all together, check out the `easygen`'s `main.go`:

#### > cmd/easygen/main.go

```go
package main

import (
	"flag"
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egCal"
	"github.com/go-easygen/easygen/egVar"
)

//go:generate sh -v easygen.gen.sh

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = Usage
	flag.Parse()

	// One mandatory non-flag arguments
	if flag.NArg() < 1 {
		Usage()
	}

	tmpl0 := easygen.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).
		Funcs(egVar.FuncDefs()).Funcs(egCal.FuncDefs())

	args := flag.Args()
	if len(easygen.Opts.TemplateStr) > 0 {
		easygen.Process0(tmpl, os.Stdout, easygen.Opts.TemplateStr, args...)
	} else {
		easygen.Process(tmpl, os.Stdout, args...)
	}
}
```

It has been as simple as this up until version 3. I.e., it's quite simple to make use of `easygen` as a package. 


### Different Versions

The `easygen` has gone through four different versions whose API are a bit different between them.

To always stay at the latest version, `import`

    "github.com/go-easygen/easygen"

in your Go code. However, to stay within a certain version, `import` the following package respectively to what you need:

- V4: "[gopkg.in/easygen.v4](https://gopkg.in/easygen.v4)"
- V3: "[gopkg.in/easygen.v3](https://gopkg.in/easygen.v3)"
- V2: "[gopkg.in/easygen.v2](https://gopkg.in/easygen.v2)"
- V1: "[gopkg.in/easygen.v1](https://gopkg.in/easygen.v1)"

To see the differences between them, check out

- [V4 vs V3](https://github.com/go-easygen/easygen/wiki/V4-vs-V3)
- [V3 vs V2](https://github.com/go-easygen/easygen/wiki/V3-vs-V2)
- [V2 vs V1](https://github.com/go-easygen/easygen/wiki/V2-vs-V1)

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

Gerrit Renker  
https://github.com/grrtrr

All patches welcome. 
