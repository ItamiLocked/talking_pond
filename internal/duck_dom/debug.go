package duckdom

import (
	"fmt"
	"os"
)

/* Debug a value of one statement */
func DebugMeDaddy(screen *Screen, content string) {
	fmt.Printf(MOVE_CURSOR_TO_POSITION, screen.Height, 1)
	fmt.Printf(CLEAR_ROW)
	fmt.Printf(MOVE_CURSOR_TO_POSITION+DEBUG_STYLES+"%s"+RESET_STYLES, screen.Height, 1, "DebugDuck: "+content)
}
func FileDebugMeDaddy(content string) {
	err := os.WriteFile("debug_result.txt", []byte(content), 0755)
	if err != nil {
		fmt.Printf("unable to write file: %s", err)
	}
}
