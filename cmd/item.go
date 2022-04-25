package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// itemCmd represents the hero command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Verwaltet die Gegenstände",
	Long: `Aufgerufen ohne Parameter werden alle Gegenstände ausgegeben. Das Kommando
unterstützt die folgenden Parameter:
add <Name> - fügt einen neuen Gegenstand hinzu
del <Name> - löscht einen Gegenstand`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(items)
	},
}

// itemAddCmd represents the add command
var itemAddCmd = &cobra.Command{
	Use:   "add <Item> <Material1:Menge1> <Material2:Menge2> ...",
	Short: "Legt einen neuen Gegenstand an",
	Long:  `Zur Anlage eines neuen Gegenstand wird ein Name und die benötigten Materialien angegeben.
Die Art des Materials und seine Anzahl werden durch ein Semikolon getrennt. Beispiel:
acr item add Messer Kupfererz:15 Zinnerz:15`,
	Run: func(cmd *cobra.Command, args []string) {
		var material = make(map[string]int)
		for _, arg := range args[1:] {
			m := strings.Split(arg, ":")
			if len(m) != 1 {
				i, _ := strconv.Atoi(m[1])
				material[m[0]] = i
			}
		}
		db.Write("items", args[0], Item{args[0], material})
	},
}

// itemDelCmd represents the add command
var itemDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Löscht einen Gegenstand",
	Long:  `Zur Löschung eines Gegenstands wird ein Name angegeben.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Delete("items", args[0])
		if err != nil {
			fmt.Printf("Der Gegenstand \"%s\" konnte nicht gelöscht werden.\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(itemCmd)
	itemCmd.AddCommand(itemAddCmd)
	itemCmd.AddCommand(itemDelCmd)
}

func getItem(name string) Item {
	for _, record := range items {
		if record.Name == name {
			return record
		}
	}
	return Item{}
}