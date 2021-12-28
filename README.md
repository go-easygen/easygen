# easygen
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-6-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-easygen/easygen?status.svg)](http://godoc.org/github.com/go-easygen/easygen)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-easygen/easygen)](https://goreportcard.com/report/github.com/go-easygen/easygen)
[![Build Status](https://github.com/go-easygen/easygen/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/go-easygen/easygen/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)



## TOC
- [easygen - Easy to use universal code/text generator](#easygen---easy-to-use-universal-codetext-generator)
- [Usage](#usage)
  - [$ easygen](#-easygen)
- [Details](#details)
- [Install Debian/Ubuntu package](#install-debianubuntu-package)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## easygen - Easy to use universal code/text generator

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
easygen version 5.1.8

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

template_name: The name for the template file.
 - Can have the extension (specified by -et) or without it.
 - Can include the path as well.
 - Can be a comma-separated list giving many template files, in which case
   at least one data_filename must be given.

data_filename(s): The name for the .yaml or .json data.
 - If omitted derive from the template_name.
 - Can have the extension or without it. If withot extension,
   will try the .yaml file first then .json
 - Can include the path as well.
 - If there is a single data_filename, and it is "-",
   then read the data (.yaml or .json format) from stdin.

Flag defaults can be overridden by corresponding environment variable, e.g.:
  EASYGEN_EY=.yml EASYGEN_ET=.tpl easygen ...
```

## Details

It can be used as a code generator, for example, command line parameter handling code generator, or anything that is structurally repetitive, like the following:

- [Introduction to easygen and its philosophy ](https://suntong.github.io/blogs/2016/01/01/easygen---easy-to-use-universal-code/text-generator)
- [Easygen is now coding itself ](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/)
- [Showcasing the power of easygen with ffcvt ](https://sfxpt.wordpress.com/2015/08/02/showcasing-the-power-of-easygen-with-ffcvt/)
- [Easygen for HTML mock-up ](https://sfxpt.wordpress.com/2015/07/10/easygen-for-mock-up/)
- [Moving beyond code-gen and mock-up, using easygen in real life creating GPT partitions](https://suntong.github.io/blogs/2015/12/26/creating-gpt-partitions-easily-on-the-command-line)

Ready to get started? Then check out [Getting Started](https://github.com/go-easygen/easygen/wiki/Getting-Started) to start building your way to turn your data into any form, any way you want.

## Install Debian/Ubuntu package

    apt install easygen

## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/go-easygen/easygen/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `easygen_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf easygen_*_linux_amd64.tar.gz
sudo mv -v easygen_*_linux_amd64/easygen /usr/local/bin/
rmdir -v easygen_*_linux_amd64
```


### Distro package

- Packages available for Linux distros are
  * [Alpine Linux](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-alpine)
  * [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb)
  * [RedHat](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-rpm)

The repo setup instruction url has been given above.
For example, for [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb) --

### Debian package


```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/suntong/repo/setup.deb.sh' \
  | sudo -E bash

# That's it. You then can do your normal operations, like

sudo apt-get update
apt-cache policy easygen

sudo apt-get install -y easygen
```

## Install Source

To install the source code instead:

```
go get -v -u github.com/go-easygen/easygen
```

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe)  
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
the _one-stop wire-framing solution_ for Go cli based projects, from _init_ to _deploy_.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/suntong"><img src="https://avatars.githubusercontent.com/u/422244?v=4?s=100" width="100px;" alt=""/><br /><sub><b>suntong</b></sub></a><br /><a href="https://github.com/go-easygen/easygen/commits?author=suntong" title="Code">💻</a> <a href="#ideas-suntong" title="Ideas, Planning, & Feedback">🤔</a> <a href="#design-suntong" title="Design">🎨</a> <a href="#data-suntong" title="Data">🔣</a> <a href="https://github.com/go-easygen/easygen/commits?author=suntong" title="Tests">⚠️</a> <a href="https://github.com/go-easygen/easygen/issues?q=author%3Asuntong" title="Bug reports">🐛</a> <a href="https://github.com/go-easygen/easygen/commits?author=suntong" title="Documentation">📖</a> <a href="#blog-suntong" title="Blogposts">📝</a> <a href="#example-suntong" title="Examples">💡</a> <a href="#tutorial-suntong" title="Tutorials">✅</a> <a href="#tool-suntong" title="Tools">🔧</a> <a href="#platform-suntong" title="Packaging/porting to new platform">📦</a> <a href="https://github.com/go-easygen/easygen/pulls?q=is%3Apr+reviewed-by%3Asuntong" title="Reviewed Pull Requests">👀</a> <a href="#question-suntong" title="Answering Questions">💬</a> <a href="#maintenance-suntong" title="Maintenance">🚧</a> <a href="#infra-suntong" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
    <td align="center"><a href="http://gerrit.sdf.org/"><img src="https://avatars.githubusercontent.com/u/5132989?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Gerrit Renker</b></sub></a><br /><a href="https://github.com/go-easygen/easygen/commits?author=grrtrr" title="Code">💻</a> <a href="#ideas-grrtrr" title="Ideas, Planning, & Feedback">🤔</a> <a href="https://github.com/go-easygen/easygen/issues?q=author%3Agrrtrr" title="Bug reports">🐛</a> <a href="#userTesting-grrtrr" title="User Testing">📓</a> <a href="#talk-grrtrr" title="Talks">📢</a> <a href="#content-grrtrr" title="Content">🖋</a> <a href="#blog-grrtrr" title="Blogposts">📝</a></td>
    <td align="center"><a href="https://github.com/bruston"><img src="https://avatars.githubusercontent.com/u/3519911?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Benjamin Ruston</b></sub></a><br /><a href="https://github.com/go-easygen/easygen/commits?author=bruston" title="Code">💻</a> <a href="https://github.com/go-easygen/easygen/issues?q=author%3Abruston" title="Bug reports">🐛</a> <a href="#userTesting-bruston" title="User Testing">📓</a></td>
    <td align="center"><a href="https://github.com/sanjaymsh"><img src="https://avatars.githubusercontent.com/u/66668807?v=4?s=100" width="100px;" alt=""/><br /><sub><b>sanjaymsh</b></sub></a><br /><a href="#platform-sanjaymsh" title="Packaging/porting to new platform">📦</a></td>
    <td align="center"><a href="https://wiki.debian.org/AnthonyFok"><img src="https://avatars.githubusercontent.com/u/1274764?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Anthony Fok</b></sub></a><br /><a href="https://github.com/go-easygen/easygen/issues?q=author%3Aanthonyfok" title="Bug reports">🐛</a> <a href="https://github.com/go-easygen/easygen/pulls?q=is%3Apr+reviewed-by%3Aanthonyfok" title="Reviewed Pull Requests">👀</a> <a href="#maintenance-anthonyfok" title="Maintenance">🚧</a> <a href="#userTesting-anthonyfok" title="User Testing">📓</a></td>
    <td align="center"><a href="https://github.com/ghost"><img src="https://avatars.githubusercontent.com/u/10137?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Deleted user</b></sub></a><br /><a href="https://github.com/go-easygen/easygen/issues?q=author%3Aghost" title="Bug reports">🐛</a> <a href="#ideas-ghost" title="Ideas, Planning, & Feedback">🤔</a> <a href="#userTesting-ghost" title="User Testing">📓</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
