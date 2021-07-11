package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/gookit/color" //https://github.com/gookit/color
)

type Config struct {
	ServiceName string
	AppHome     string
}

// Unmarshal yaml file - returns struct
func read_file(filename string) (txt string) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	txt = string(source)
	return txt
}
func main() {
	version := "0.1"
	flag.Usage = func() {
		usagemsg := "epub-fix, version: " + version + " - Fix referneces in ePub"
		fmt.Fprintf(os.Stderr, usagemsg+"%s -x=forord.xml", os.Args[0])
		flag.PrintDefaults()
		color.Error.Println("Prøv igjen!")
	}
	x := flag.String("x", "", "Hvilke xml fil skal du fikse?")

	// imp := flag.Bool("imp", false, "If imp is used - script will be importing and not exporting pdb")

	flag.Parse()
	if *x == "" {
		flag.Usage()
		os.Exit(1)
	} else {
		color.Info.Println("epub-fix version: " + version)
		color.Info.Println("xmlfile:", *x)

	}
	re := regexp.MustCompile(".xml$")
	newfilename := re.ReplaceAllString(*x, "_FIXED.xml")
	oldtxt := read_file(*x)
	matched, _ := regexp.MatchString(`footnote_plugin_reference_`, oldtxt)

	// r := regexp.MustCompile(`(?s)<span class="footnote_referrer"><a role="button" tabindex="0" onkeypress="footnote_moveToReference_\d+_\d+\('footnote_plugin_reference_(\d+_\d+_\d+)'\);"><sup id="footnote_plugin_tooltip_\d+_\d+_\d+" class="footnote_plugin_tooltip_text">\[(\d+)\].*?<span id="footnote_plugin_tooltip_text_\d+_\d+_\d+".*?<span class="footnote_tooltip_continue">Continue reading</span></span></span>`)
	r := regexp.MustCompile(`(?s)<span class="footnote_referrer"><a role="button" tabindex="0" onkeypress="footnote_moveToReference_\d+_\d+\('footnote_plugin_reference_(\d+_\d+_\d+)'\);"><sup id="footnote_plugin_tooltip_\d+_\d+_\d+" class="footnote_plugin_tooltip_text">\[(\d+)\].*?<span id="footnote_plugin_tooltip_text_(\d+_\d+_\d+)".*?</span></span>`)
	r2 := regexp.MustCompile(`<a id="footnote_plugin_reference_(\d+_\d+_\d+)" class="footnote_backlink"><span class="footnote_index_arrow">↑</span>(\d+)</a>`)
	r3 := regexp.MustCompile(`<span class="footnote_tooltip_continue">Continue reading</span>`)
	// footnote := r.FindString(oldtxt)
	// newtxt := r.ReplaceAllString(oldtxt, `<span class="footnote_referrer"><a href="#footnote_plugin_reference_$1"><sup id="footnote_plugin_tooltip_$1" class="footnote_plugin_tooltip_text">[$2]</sup></a></span>`)
	newtxt := r3.ReplaceAllString(oldtxt, ``)
	newtxt = r.ReplaceAllString(newtxt, `<span class="footnote_referrer"><a href="#footnote_plugin_reference_$1"><sup id="footnote_plugin_tooltip_$1" class="footnote_plugin_tooltip_text">[$2]</sup></a></span>`)
	newtxt = r2.ReplaceAllString(newtxt, `<a id="footnote_plugin_reference_$1" class="footnote_backlink" href="#footnote_plugin_tooltip_$1"><span class="footnote_index_arrow">↑</span>$2</a>`)

	// color.Info.Println("Footnote: " + footnote)

	// fmt.Printf("%q\n", r.FindString(oldtxt))
	// re := regexp.MustCompile(`foo.?`)
	// fmt.Printf("%q\n", re.FindString("seafood fool"))
	// fmt.Printf("%q\n", re.FindString("meat"))

	if matched {
		color.Info.Println("Fant fotnote")
		// fmt.Println(newtxt)
		err := ioutil.WriteFile(newfilename, []byte(newtxt), 0777)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		color.Warn.Println("Fant ingen fotnote")
	}

}
