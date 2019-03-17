package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var arguments = struct {
	Input        string
	Concurrency  int
	TrackersFile string
}{}

func parseArgs(args []string) {
	// Create new parser object
	parser := argparse.NewParser("TorrentDeprivatiser", "Replace the Announce URL with Public ones and remove the Private Bit on a folder full of .torrent")

	// Create flags
	input := parser.String("i", "input", &argparse.Options{
		Required: true,
		Help:     "Input directory"})

	concurrency := parser.Int("j", "concurrency", &argparse.Options{
		Required: false,
		Help:     "Concurrency",
		Default:  4})

	trackersFile := parser.String("t", "trackers", &argparse.Options{
		Required: true,
		Help:     "Tracker list file"})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	// Fill arguments structure
	arguments.Input = *input
	arguments.Concurrency = *concurrency
	arguments.TrackersFile = *trackersFile
}
