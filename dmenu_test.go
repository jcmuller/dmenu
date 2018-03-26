package dmenu_test

import (
	"fmt"

	"github.com/jcmuller/dmenu"
)

func ExamplePopup() {
	output, err := dmenu.Popup("Choose:", "One", "Two", "Three")

	if err != nil {
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_Popup_rofi() {
	dmenuRofi := dmenu.New("rofi", "-dmenu -p %s")

	output, err := dmenuRofi.Popup("Choose:", "One", "Two", "Three")

	if err != nil {
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_Popup_mock() {
	dmenuMock := dmenu.New("mockdmenu", "-p 1 %s")

	output, err := dmenuMock.Popup("Choose:", "One", "Two", "Three")

	if err != nil {
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_Popup_zenity() {
	dmenuZenity := dmenu.New("zenity", `--title=%s --list --column=Options`)

	output, err := dmenuZenity.Popup("Choose:", "One", "Two", "Three")

	if err != nil {
		fmt.Printf("%q\n", err)
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}
