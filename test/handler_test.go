package test

import (
	"fmt"
	"testing"

	"github.com/zuadi/csvHandler"
	checkressources "github.com/zuadi/csvHandler/checkRessources"
)

func TestCSVHandler(t *testing.T) {

	handler := csvHandler.NewCSVHandler(';')
	if err := handler.LoadAllData("../dist/Bibleindex.csv"); err != nil {
		t.Fatal(err)
	}

	found, notfound := handler.FindValueInColumn(handler.GetRowValuesAsMap(0)["German Schlachter 2000"], "1 Mose 1.1", "1 Mose 1.2", "1 Mose 1.3", "1 Mose 1.4", "1 Mose 1.5")
	if notfound != nil {
		t.Fatal(notfound)
	}

	for i := range found {
		fmt.Println(100, found[i])
	}

	checkressources.PrintMemoryUsage()

}
