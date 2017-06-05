# Jump start the cli based project

This is on how to jump-start a [`cli`](https://github.com/mkideal/cli/) based project.

The [`cli`](https://github.com/mkideal/cli/) package really make it easy to code your command line handling program, especially when you need to deal with _sub commands_ like `apt-get`, `git` etc do. Now, want to make it even simpler? 

## Usage

Yes! Use the code auto generation. It can make things even more simple. Run the following on command line so that you can see for yourself:

```
go get github.com/go-easygen/easygen
ls -l $GOPATH/bin
export PATH=$PATH:$GOPATH/bin

easygen $GOPATH/src/github.com/go-easygen/easygen/test/commandlineCLI-024
```

## Result

You should see the following Go code:

```go
// -*- go -*-
////////////////////////////////////////////////////////////////////////////
// Program: gogo
// Purpose: Golang package manager
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

var app = &cli.Register(&cli.Command{
	Name: "gogo",
	Desc: "Golang package manager",
	Text: "  gogo is a new golang package manager\n  very very good",
	Argv: func() interface{} { return new(gogoT) },
	Fn:   gogo,

	NumArg:      cli.ExactN(1),
})

type gogoT struct {
	cli.Helper
	Version	bool	`cli:"v,version" usage:"display version"`
	List	bool	`cli:"l,list" usage:"list all sub commands or not"`
}

func gogo(ctx *cli.Context) error {
	argv := ctx.Argv().(*gogoT)
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	ctx.String("[gogo]: %v\n", ctx.Args())

	return nil
}


////////////////////////////////////////////////////////////////////////////
// Program: build
// Purpose: Build golang application
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

var buildCmd = app.Register(&cli.Command{
	Name: "build",
	Desc: "Build golang application",
	Text: "Usage:\n  gogo build [Options] Arch(i386|amd64)",
	Argv: func() interface{} { return new(buildT) },
	Fn:   build,

	NumArg:      cli.ExactN(1),
	CanSubRoute: true,
})

type buildT struct {
	cli.Helper
	Dir	string	`cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix	string	`cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out	string	`cli:"o,out" usage:"output filename"`
}

func build(ctx *cli.Context) error {
	argv := ctx.Argv().(*buildT)
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	ctx.String("[build]: %v\n", ctx.Args())

	return nil
}

////////////////////////////////////////////////////////////////////////////
// Program: install
// Purpose: Install golang application
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

var installCmd = app.Register(&cli.Command{
	Name: "install",
	Desc: "Install golang application",
	Text: "Usage:\n  gogo install [Options] package [package...]",
	Argv: func() interface{} { return new(installT) },
	Fn:   install,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
})

type installT struct {
	cli.Helper
	Dir	string	`cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix	string	`cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out	string	`cli:"o,out" usage:"output filename"`
}

func install(ctx *cli.Context) error {
	argv := ctx.Argv().(*installT)
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	ctx.String("[install]: %v\n", ctx.Args())

	return nil
}

////////////////////////////////////////////////////////////////////////////
// Program: publish
// Purpose: Publish golang application
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

var publishCmd = app.Register(&cli.Command{
	Name: "publish",
	Desc: "Publish golang application",
	Argv: func() interface{} { return new(publishT) },
	Fn:   publish,
})

type publishT struct {
	cli.Helper
	Dir	string	`cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix	string	`cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out	string	`cli:"o,out" usage:"output filename"`
	List	bool	`cli:"l,list" usage:"list all sub commands"`
}

func publish(ctx *cli.Context) error {
	argv := ctx.Argv().(*publishT)
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	ctx.String("[publish]: %v\n", ctx.Args())

	return nil
}

```   

As you can see the skeleton of the program has been automatically generated. 

## Source

So what's behind the scene of the "magic"? Take a look at the driving [Yaml file](https://github.com/go-easygen/easygen/blob/master/test/commandlineCLI-024.yaml), that's everything you specify to make the magic happen. I'm republishing it below for your reference:

```yaml
# program name, name for the executable
ProgramName: gogo

# package name
# - For standalone program that does not belong to any package, e.g., 
#   https://github.com/suntong/easygen/blob/7791e4f0e5605543d27da1671a21376cdb9dcf2a/easygen/easygen.go
#   just ignore the first line, the `package` output, and copy the rest
# - If you don't mind using a separated file to handle commandline paramters,
#   then name the package as "main". see the spin-out "TF-minus1.go" file under
#   https://github.com/suntong/easygen/tree/d1ab0b5fe80ddac57fe9ef51f6ccb3ab998cd5ee
# - If you are using it in a pacakge, look no further than
#   https://github.com/suntong/easygen/blob/master/easygenapi/config.go
#   which was a direct dump: easygen test/commandlineFlag | gofmt > easygenapi/config.go
#
PackageName: main

Name: gogo
Var: app
Desc: "Golang package manager"
Text: '  gogo is a new golang package manager\n  very very good'
NumArg: cli.ExactN(1)

Options:
  - Name: Version
    Type: bool
    Flag: v,version
    Usage: display version

  - Name: List
    Type: bool
    Flag: l,list
    Usage: list all sub commands or not


Command:

  - Name: build
    Desc: "Build golang application"
    Text: 'Usage:\n  gogo build [Options] Arch(i386|amd64)'
    NumArg: cli.ExactN(1)
    
    Options:
      - Name: Dir
        Type: string
        Flag: dir
        Value: '"./"'
        Usage: source code root dir

      - Name: Suffix
        Type: string
        Flag: suffix
        Value: '".go,.c,.s"'
        Usage: "source file suffix"

      - Name: Out
        Type: string
        Flag: o,out
        Usage: "output filename"

  - Name: install
    Desc: "Install golang application"
    Text: 'Usage:\n  gogo install [Options] package [package...]'
    NumArg: cli.AtLeast(1)
    
    Options:
      - Name: Dir
        Type: string
        Flag: dir
        Value: '"./"'
        Usage: source code root dir

      - Name: Suffix
        Type: string
        Flag: suffix
        Value: '".go,.c,.s"'
        Usage: "source file suffix"

      - Name: Out
        Type: string
        Flag: o,out
        Usage: "output filename"

  - Name: publish
    Desc: "Publish golang application"
    
    Options:
      - Name: Dir
        Type: string
        Flag: dir
        Value: '"./"'
        Usage: source code root dir

      - Name: Suffix
        Type: string
        Flag: suffix
        Value: '".go,.c,.s"'
        Usage: "source file suffix"

      - Name: Out
        Type: string
        Flag: o,out
        Usage: "output filename"

      - Name: List
        Type: bool
        Flag: l,list
        Usage: "list all sub commands"

```

Yes, you are right -- that's all you need to have the above code skeleton generated. 

## Takeaway

### Understanding how to use easygen for cli based project

Now, the question is, how it can help your own project. Assuming you have never code a `cli` based project before, what you need to do are, 

- Take a look at the above [Yaml file](https://github.com/go-easygen/easygen/blob/master/test/commandlineCLI-024.yaml) and see the _generated Go code_, and see their relationship.
- Take a look at the [orginal finished Go code](https://github.com/mkideal/cli/tree/master/_examples/024-multi-command) that such code generation is for, and see what else you need to do to make it a working code.
- The generated Go code was meant to be cut into each corresponding Go code files, e.g., the `main.go`, `build.go`, `install.go` and `publish.go` as in the above [finished code](https://github.com/mkideal/cli/tree/master/_examples/024-multi-command)
- At this point, you have a good starting point already, because the code skeletons are already in place. 

### Round-trip generation

This is for illustration only. However, for round-trip generation, I.e., you generate and use the generated code, and you find yourself need to add more options later, then such manual cut & paste (into each corresponding Go code) method will be flimsy. So what shall we do? 

- Take a look at the [template file](https://github.com/go-easygen/easygen/blob/master/test/commandlineCLI.tmpl) that is being used to generate the Go code. This is where the _true_ magic happens. 
- I recommend removing all the repeating part, and the sample implementation part out. I.e., your own `build.go`, `install.go` and `publish.go` contains the true implementation, and use the code-gen to automatically generate all the declaration part. This way the automatically-generated code can always been overwritten (with the latest definition) without affecting your existing implementation.
- For details on code automatic generation, refer to [easygen](https://github.com/go-easygen/easygen/) for all other information.

### Practical Example

There is a new `commandlineCLI-027` set of data and template, that is built after [code sample 027](https://github.com/mkideal/cli/tree/master/_examples/027-global-option). Replace the above `commandlineCLI-024` with `commandlineCLI-027`, you will get code for the corresponding samples. Further more, t

The `commandlineCLI-027` is built with the above *"round-trip generation"* principle in mind. Take a look at [`redoCmds.go`](https://github.com/suntong/lang/blob/master/lang/Go/src/sys/CLI/027-global-redo/redoCmds.go) in the [027-global-redo](https://github.com/suntong/lang/tree/master/lang/Go/src/sys/CLI/027-global-redo) project, it is completely generated by [easygen](https://github.com/go-easygen/easygen/), and can be regenerated as many times as you want, and whenever you want, even after several releases. 


## Cookbook

Since I've written the above how-to, I've followed it to generate lots of my code wire-frames, polishing and fine-tuning my approaches along the way. Now, the procedure has mostly been stabilized, so I think it is a good idea to share
how I am doing it, generating my cli-based code wire-frames. All the following code-gen are based on my [cli-t0.tmpl](https://github.com/suntong/cli/blob/78b8d5e8b97d0950faec5e7844f220a572577ba0/cli-t0.tmpl) template. 

### Simple case, no sub-commands

Let's start with the simpler case, generating cli handling code that has no sub-commands. This is the way tradition Unix tools use, and there are many tradition *nix tools that are doing this way, e.g., `cp`, `rm`, `find`, `rsync` etc, to name a few.

Here is how to generate a cli handling code-base for this case, using [my picv v0.1.0 ](https://github.com/suntong/picv/tree/0.1.0) as the example:

1. Provide a `.yaml` file -- [picv.yaml](https://github.com/suntong/picv/blob/0.1.0/picv.yaml)
1. Provide a code-gen shell script -- [picvCLIGen.sh](https://github.com/suntong/picv/blob/0.1.0/picvCLIGen.sh) and run it to generate [picvCLIDef.go](https://github.com/suntong/picv/blob/0.1.0/picvCLIDef.go)
1. Extract the commented out code to [picvMain.go](https://github.com/suntong/picv/blob/0.1.0/picvMain.go) and [picvCLICmd.go](https://github.com/suntong/picv/blob/0.1.0/picvCLICmd.go)
1. Now the code wire-frame is done, and ready for compilation and execution.

```
$ picv -h
picture vault
built on 2017-06-03

Tool to deal with camera pictures and put them in vault

Usage:
  picv [Options] dir [dirs...]

Options:

  -h, --help       display help information
  -g, --gap[=5]    Gap in days to be considered as different group/vault
  -p, --pod[=15]   Minimum number of picture to have before splitting to a different group/vault
  -v, --verbose    Verbose mode (Multiple -v options increase the verbosity.)

$ picv 
{"Help":false,"Gap":5,"Pod":15,"Verbose":{}}{"Help":false,"Gap":5,"Pod":15,"Verbose":{}}
```

If you need to tweak the `.yaml` file further, just

1. Update the `.yaml` file
1. Run `go generate -x` to update the [picvCLIDef.go](https://github.com/suntong/picv/blob/0.1.0/picvCLIDef.go) file
1. And then do more editing to [picvCLICmd.go](https://github.com/suntong/picv/blob/0.1.0/picvCLICmd.go) to incorporate the new changes


### With sub-commands case

The _sub commands_ are modern way of providing cli interfaces, like `apt-get`, `git` etc. As a comparison, the tradition rcs (Revision Control System) tool uses different commands for version repo management, check-in, check-out, and view logs (`rcs`, `ci`, `co`, `rlog` respectively), but `git` provides these functionalities all under the same `git` command, but different sub commands instead (`clone`/`init` etc, `commit`, `checkout`, `log`).

How to generate cli handling code that has sub-commands? The procedure is exactly the same as above. I'm using [my picv v0.2.0 ](https://github.com/suntong/picv/tree/0.2.0) this time as the example:

1. Provide a `.yaml` file -- [picv.yaml](https://github.com/suntong/picv/blob/0.2.0/picv.yaml). You can see that the `.yaml` file only changes in its organization, while the content are reused nearly 100%, only that [new sub-commands are added](https://github.com/suntong/picv/blob/10daaa84d9545ffb129c4909de7019dc896fd0be/picv.yaml) and [options have been shifted around](https://github.com/suntong/picv/commit/10daaa84d9545ffb129c4909de7019dc896fd0be#diff-78301e4359962421c827625b6393ab5b).
1. Provide a code-gen shell script -- [picvCLIGen.sh](https://github.com/suntong/picv/blob/0.2.0/picvCLIGen.sh) and run it to generate [picvCLIDef.go](https://github.com/suntong/picv/blob/0.2.0/picvCLIDef.go). This step is exactly as above. 
1. Extract the commented out code to [picvMain.go](https://github.com/suntong/picv/blob/0.2.0/picvMain.go), [cmdCut.go](https://github.com/suntong/picv/blob/0.2.0/cmdCut.go) and the new [cmdArch.go](https://github.com/suntong/picv/blob/0.2.0/cmdArch.go). Note that I leave the _Main dispatcher_ in `picvMain.go` in this case, because it is clearer, as each sub commands just need to deal with their own handling. 
1. Now the code wire-frame is done, and ready for compilation and execution.

```
$ picv -h
picture vault
built on 2017-06-03

Tool to deal with camera pictures and put them in vault

Options:

  -h, --help      display help information
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity.)

Commands:

  cut    Separate picture into groups
  arch   Archive groups of picture into vaults


$ picv cut -h 
Separate picture into groups

Options:

  -h, --help       display help information
  -v, --verbose    Verbose mode (Multiple -v options increase the verbosity.)
  -g, --gap[=5]    Gap in days to be considered as different group/vault
  -p, --pod[=15]   Minimum number of picture to have before splitting to a different group/vault

$ picv cut 
{"Gap":5,"Pod":15,"Verbose":0}

$ picv cut -vv 
{"Gap":5,"Pod":15,"Verbose":2}

```

The rest are the same  as above as well.

## Further abstraction

Note the output of my picv [v0.1.0 ](https://github.com/suntong/picv/tree/0.1.0) and [v0.2.0 ](https://github.com/suntong/picv/tree/0.2.0) are different. This is because I've done a further abstraction step (manually, but could also be somewhat automated if you want) in [v0.1.1 ](https://github.com/suntong/picv/tree/0.1.1), which defined an option structure that is independent of the `mkideal/cli` package.

I view this further abstraction a _logical_ representation of the options the tool can use, while the option structure defined in [picvCLIDef.go](https://github.com/suntong/picv/blob/master/picvCLIDef.go) is a _"physical"_ representation of the options, because its foreign-package-depending nature. 

This step is not absolutely necessary, but I found such practice more organized. Note [how it is masking/hiding the implementation detail that `-v` is from global option, and how easy it is to make the corresponding changes when the `-v` option changed its definition location](https://github.com/suntong/picv/commit/10daaa84d9545ffb129c4909de7019dc896fd0be#diff-e7c16ee367258e94c2fc12e8b273d921R32).
