package cmd

import (
	"log"

	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"github.com/vlbeaudoin/studybuddy/data"
)

func listAllNotes() {
	notes, err := data.ListNotes()
	if err != nil {
		log.Fatal(err)
	}

	t := tabby.New()
	t.AddHeader("WORD", "DEFINITION", "CATEGORY")
	for _, note := range notes {
		t.AddLine(note.Word, note.Definition, note.Category)
	}
	t.Print()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all studybuddy notes.",
	Run: func(cmd *cobra.Command, args []string) {
		listAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
