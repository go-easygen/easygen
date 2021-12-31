
## {{toc 5}}
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## {{.Name}} - Easy to use universal code/text generator

Command `{{.Name}}` is an easy to use universal code/text generator.

It can be used as a text or html generator for _arbitrary_ purposes with _arbitrary_ data and templates. It is a good [GSL](https://github.com/imatix/gsl) replacement, as it

  - is more easy to define driving data, in form of YML instead of XML
  - has more powerful template engine that based on Go template.
    You can also write your own function in Go to customize your template.
  - there are lots of transformation support functions builtin inside it, check out the full list with sample usages and results [here](https://github.com/go-easygen/easygen/issues/25).
  - apart from the above sample usages, it has extensive documentations like [this](https://github.com/go-easygen/easygen/wiki/Docs:-Easygen-usage). Check out the [wiki](https://github.com/go-easygen/easygen/wiki/) for the full list.

You can even use easygen as [a generic Go template testing tool](https://github.com/go-easygen/easygen/wiki/Tip:-Testing-the-templates-on-the-fly) using the `-ts` commandline option, and much more.

Note this document is for `{{.Name}}` versions 4.0+. For historic versions check out the [Different Versions](https://github.com/go-easygen/easygen/wiki/Docs:-Different-Versions) section.


## Usage

### $ {{exec "easygen" | color "sh"}}

## Details

It can be used as a code generator, for example, command line parameter handling code generator, or anything that is structurally repetitive, like the following:

- [Introduction to easygen and its philosophy ](https://suntong.github.io/blogs/2016/01/01/easygen---easy-to-use-universal-code/text-generator)
- [What is the "XSLT" equivalent for JSON? Here is the new answer](https://dev.to/suntong/what-is-the-xslt-equivalent-for-json-here-is-the-new-answer-7la)
- [Transforming json data with easygen](https://dev.to/suntong/transforming-json-data-with-easygen-4g2i)
- [Easygen is now coding itself ](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/)
- [Showcasing the power of easygen with ffcvt ](https://sfxpt.wordpress.com/2015/08/02/showcasing-the-power-of-easygen-with-ffcvt/)
- [Easygen for HTML mock-up ](https://sfxpt.wordpress.com/2015/07/10/easygen-for-mock-up/)
- [Moving beyond code-gen and mock-up, using easygen in real life creating GPT partitions](https://suntong.github.io/blogs/2015/12/26/creating-gpt-partitions-easily-on-the-command-line)

Ready to get started? Then check out [Getting Started](https://github.com/go-easygen/easygen/wiki/Getting-Started) to start building your way to turn your data into any form, any way you want.

## Install Debian/Ubuntu package

    apt install {{.Name}}

