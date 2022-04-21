package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// planCmd represents the hero command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Verwaltet die Pläne zur Herstellung von Gegenständen",
	Long: `Aufgerufen ohne Parameter werden alle Pläne ausgegeben. Das Kommando
unterstützt die folgenden Parameter:
add <Name> - fügt einen neuen Plan hinzu
del <Name> - löscht einen Plan
finish <Name> - beendet einen Plan, fügt den Gegenstand seinem Helden hinzu`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(plans)
	},
}

// planAddCmd represents the add command
var planAddCmd = &cobra.Command{
	Use:   "add <Hero> <Item1>  <Item2> ...",
	Short: "Legt einen neuen Plan zur Herstellung eines Gegenstands an",
	Long:  `Zur Anlage eines neuen Plans wird ein Held und die geplanten Items angegeben.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Write("plans", args[0], Plan{args[0], args[1:]})
	},
}

// del is used to delete a plan from the database
// It is called by planDelCmd and planFinishCmd
func del(cmd *cobra.Command, args []string) {
	err := db.Delete("plans", args[0])
	if err != nil {
		fmt.Printf("Der Plan \"%s\" konnte nicht gelöscht werden.\n", args[0])
	}
}

// planDelCmd represents the add command
var planDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Löscht einen Plan",
	Long:  `Zur Löschung eines Plans wird ein Name angegeben.`,
	Run:   del,
}

// planFinishCmd represents the finish command
var planFinishCmd = &cobra.Command{
	Use:   "finish <hero>",
	Short: "Beendet einen Plan und fügt die Items zum Helden hinzu",
	Long:  `Zur Beendigung eines Plans wird der Name des Helden angegeben.`,
	Run: func(cmd *cobra.Command, args []string) {
		var items []string
		for _, p := range plans {
			if p.Hero == args[0] {
				items = p.Items
				break
			}
		}
		db.Write("heros", args[0], Hero{args[0], items})
		del(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.AddCommand(planAddCmd)
	planCmd.AddCommand(planDelCmd)
	planCmd.AddCommand(planFinishCmd)
}
