package dmenu_test

import (
	"fmt"

	"github.com/jcmuller/dmenu"
)

func ExamplePopup() {
	output, err := dmenu.Popup("Choose an option:", "One word", "Two", "Three things")

	if err != nil {
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_Popup_rofi() {
	dmenuRofi := dmenu.New("rofi", "-dmenu", "-p", "%s")

	output, err := dmenuRofi.Popup("Choose an option:", "One word", "Two", "Three things")

	if err != nil {
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_Popup_mock() {
	dmenuMock := dmenu.New("mockdmenu", "-p", "1", "%s")

	output, err := dmenuMock.Popup("Choose an option:", "One word", "Two", "Three things")

	if err != nil {
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_Popup_zenity() {
	dmenuZenity := dmenu.NewZenityList()
	output, err := dmenuZenity.Popup("Choose an option:", "One word", "Two", "Three things")

	if err != nil {
		fmt.Printf("%q\n", err)
		if err, ok := err.(*dmenu.EmptySelectionError); !ok {
			fmt.Println(fmt.Errorf("Error getting output: %s", err))
		}
	}

	fmt.Println(output)
	// Output: Two
}

func ExampleDmenu_YesNo_zenity() {
	dmenuZenity := dmenu.NewZenityYesNo()
	answer, err := dmenuZenity.YesNo("Are you sure you want to do that?")

	if err != nil {
		panic(err)
	}

	fmt.Println(answer)
	// Output: true
}
