package easygenapi

import "regexp"

var (
	re  *regexp.Regexp
	rec *regexp.Regexp
)

func TF_strings_init() {
	re = regexp.MustCompile(`(?i)` + Opts.StrFrom)
	// case sensitive string replace
	rec = regexp.MustCompile(Opts.StrFrom)

}

//==========================================================================
// template function

func replace(replStr string) string {
	return re.ReplaceAllString(replStr, Opts.StrTo)
}

// replacec does a case sensitive string replace
func replacec(replStr string) string {
	return rec.ReplaceAllString(replStr, Opts.StrTo)
}
