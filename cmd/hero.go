package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// heroCmd represents the hero command
var heroCmd = &cobra.Command{
	Use:   "hero",
	Short: "Verwaltet die Helden",
	Long: `Aufgerufen ohne Parameter werden alle Helden ausgegeben. Das Kommando
unterstützt die folgenden Parameter:
add <Name> - fügt einen neuen Helden hinzu
del <Name> - löscht einen Helden`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(heros)
	},
}

// heroAddCmd represents the add command
var heroAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Legt einen neuen Helden an",
	Long: `Zur Anlage eines neuen Helden wird ein Name angegeben.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Write("heros", args[0], Hero{args[0]})
	},
}

// heroDelCmd represents the add command
var heroDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Löscht einen Helden",
	Long: `Zur Löschung eines Helden wird ein Name angegeben.`,
	Run: func(cmd *cobra.Command, args []string) {
		err:= db.Delete("heros", args[0])
		if err != nil {
			fmt.Printf("Der Held \"%s\" konnte nicht gelöscht werden.\n", args[0])
		}
	},
}
func init() {
	rootCmd.AddCommand(heroCmd)
	heroCmd.AddCommand(heroAddCmd)
	heroCmd.AddCommand(heroDelCmd)
}