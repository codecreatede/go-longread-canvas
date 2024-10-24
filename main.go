package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date: 2024-10-24


A long read canvas profiling golang application that allows you to scan for the long reads and
either extract the motifs from the long reads or remove them from the long reads.

A golang implementation of the pacbioHifiFilt


*/

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

var (
	longread    string
	motiflooker string
)

var rootCmd = &cobra.Command{
	Use:  "longread",
	Long: "look for the matching patterns",
	Run:  joinFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&longread, "longread", "L", "path to the long read file", "long read file to be checked")
	rootCmd.Flags().
		StringVarP(&motiflooker, "pattern", "P", "path to the file containing the patterns", "pattern file")
}

func joinFunc(cmd *cobra.Command, args []string) {
	type pacbiofileID struct {
		id string
	}
	type pacbiofileSeq struct {
		seq string
	}
	pacbioIDConstruct := []pacbiofileID{}
	pacbioSeqConstruct := []pacbiofileSeq{}

	fpacbio, err := os.Open(longread)
	if err != nil {
		log.Fatal(err)
	}
	Opacbio := bufio.NewScanner(fpacbio)
	for Opacbio.Scan() {
		line := Opacbio.Text()
		if strings.HasPrefix(string(line), "@") {
			pacbioIDConstruct = append(pacbioIDConstruct, pacbiofileID{
				id: strings.ReplaceAll(strings.Split(string(line), " ")[0], "@", ""),
			})
		}
		if strings.HasPrefix(string(line), "A") || strings.HasPrefix(string(line), "T") ||
			strings.HasPrefix(string(line), "G") ||
			strings.HasPrefix(string(line), "C)") {
			pacbioSeqConstruct = append(pacbioSeqConstruct, pacbiofileSeq{
				seq: string(line),
			})
		}
	}

	patterns := []string{}
	pOpen, err := os.Open(motiflooker)
	if err != nil {
		log.Fatal(err)
	}
	pRead := bufio.NewScanner(pOpen)
	for pRead.Scan() {
		line := pRead.Text()
		patterns = append(patterns, string(line))
	}

	fID := []string{}
	fSeq := []string{}

	for i := range pacbioIDConstruct {
		fID = append(fID, pacbioIDConstruct[i].id)
		fSeq = append(fSeq, pacbioSeqConstruct[i].seq)
	}

	type pSearch struct {
		id    string
		seq   string
		start int
		end   int
	}

	pSearchAppend := []pSearch{}

	for i := range fSeq {
		for j := range patterns {
			pSearchAppend = append(pSearchAppend, pSearch{
				id:    fID[i],
				seq:   fSeq[j],
				start: strings.Index(fSeq[i], patterns[j]),
				end:   strings.Index(fSeq[i], patterns[j]),
			})
		}
	}

	type pAnnotate struct {
		id           string
		seq          string
		priorPsearch string
		postPsearch  string
	}

	pAnnotateAppend := []pAnnotate{}

	for i := range pSearchAppend {
		pAnnotateAppend = append(pAnnotateAppend, pAnnotate{
			id:           pSearchAppend[i].id,
			seq:          pSearchAppend[i].seq,
			priorPsearch: pSearchAppend[i].seq[:pSearchAppend[i].start],
			postPsearch:  pSearchAppend[i].seq[pSearchAppend[i].end:],
		})
	}

	type pJoinA struct {
		id      string
		seq     string
		joinSeq string
	}

	pJoinAppend := []pJoinA{}
	for i := range pAnnotateAppend {
		joinCapture := []string{}
		joinCapture = append(
			joinCapture,
			pAnnotateAppend[i].priorPsearch,
			pAnnotateAppend[i].postPsearch,
		)
		pJoinAppend = append(pJoinAppend, pJoinA{
			id:      pAnnotateAppend[i].id,
			seq:     pAnnotateAppend[i].seq,
			joinSeq: strings.Join(joinCapture, ""),
		})
	}
}
