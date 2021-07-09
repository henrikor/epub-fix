package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/gookit/color" //https://github.com/gookit/color
)

func main() {
	flag.Usage = func() {
		usagemsg := "Fix referneces in ePub"
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
		color.Info.Println("xmlfile:", *x)

	}


	// matched, _ := regexp.MatchString(`,`, *g)

}