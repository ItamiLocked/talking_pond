package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"os"

	dd "github.com/Robert-Duck-by-BB-SR/talking_pond/internal/duck_dom"
	"golang.org/x/term"

	tpc "github.com/Robert-Duck-by-BB-SR/talking_pond/internal/tps_client"
)

// var frame_chars = []byte{' ', '`', '.', ',', '~', '+', '*', '&', '#', '@'}

// type CharMeDaddy struct {
// 	char, count, r, g, b byte
// }

// func encode_frame(img image.Image) []byte {
// 	orig_bounds := img.Bounds().Max
//
// 	scale_x := orig_bounds.X / 80
// 	scale_y := orig_bounds.Y / 40
// 	new_img_x := orig_bounds.X / scale_x
// 	new_img_y := orig_bounds.Y / scale_y
//
// 	encoded_data := []byte{}
// 	all_rle := []CharMeDaddy{}
//
// 	for y := range new_img_y {
// 		for x := range new_img_x {
// 			r, g, b, _ := img.At(x*scale_x, y*scale_y).RGBA()
// 			lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
// 			indx := lum * uint32(len(frame_chars)) / 256
// 			// sliding window -> 5 bytes
// 			// 0 - char
// 			// 1 - repeat
// 			// 2 - r
// 			// 3 - g
// 			// 4 - b
// 			// 5 - new line
// 			if x == 0 {
// 				all_rle = append(all_rle, CharMeDaddy{frame_chars[indx], 1, uint8(r), uint8(g), uint8(b)})
// 			} else {
// 				curr_rle := &all_rle[len(all_rle)-1]
// 				if frame_chars[indx] == curr_rle.char &&
// 					uint8(r) == curr_rle.r &&
// 					uint8(g) == curr_rle.g &&
// 					uint8(b) == curr_rle.b {
// 					curr_rle.count += 1
// 				} else {
// 					all_rle = append(all_rle, CharMeDaddy{frame_chars[indx], 1, uint8(r), uint8(g), uint8(b)})
// 				}
// 			}
// 		}
// 		for _, el := range all_rle {
// 			encoded_data = append(encoded_data, el.char, el.count, el.r, el.g, el.b)
// 		}
// 		all_rle = []CharMeDaddy{}
// 		encoded_data = append(encoded_data, '\n')
// 	}
// 	return encoded_data
// }

// func decode_frame(enc_data []byte) {
// 	// sliding window -> 5 bytes
// 	// 0 - char
// 	// 1 - repeat
// 	// 2 - r
// 	// 3 - g
// 	// 4 - b
// 	// 5 - new line
// 	fmt.Print("\033[2J\033[H")
// 	for i := 0; i < len(enc_data); i += 5 {
// 		if enc_data[i] == '\n' {
// 			// or i -= 4
// 			i += 1
// 			fmt.Println()
// 			if i >= len(enc_data) {
// 				break
// 			}
// 		}
//
// 		for reps := 0; reps < int(enc_data[i+1]); reps += 1 {
// 			r := enc_data[i+2]
// 			g := enc_data[i+3]
// 			b := enc_data[i+4]
//
// 			var cell string = fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, enc_data[i])
// 			fmt.Print(cell)
// 		}
//
// 	}
// }

// func move_cursor(screen *dd.Screen, item dd.Renderable, direction int) {
// 	new_index := item.ActiveIndex() + direction
// 	if new_index >= 0 && new_index < len(screen.Windows) {
// 		active_item := item.Active()
// 		active_item.SetBackground("")
//
// 		item.SetActive(new_index)
//
// 		next_active_item := item.Active()
// 		next_active_item.SetBackground(dd.INVERT_STYLES)
//
// 		screen.RenderQueue = append(screen.RenderQueue, active_item, next_active_item)
// 		screen.CursorPos = next_active_item.GetPos()
// 	}
// }

func create_main_window(screen *dd.Screen) {
	width, height, _ := term.GetSize(int(os.Stdin.Fd()))
	screen.Width = width
	screen.Height = height

	sidebar := dd.CreateWindow(dd.Styles{
		Width:      50,
		Height:     screen.Height - 1,
		Background: dd.PRIMARY_THEME.PrimaryBg,
		Paddding:   1,
		Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.PRIMARY_THEME.SecondaryTextColor},
	})

	sidebar.AddComponent(
		dd.CreateComponent("Deez nuts123123", dd.Styles{
			MaxWidth:  10,
			TextColor: dd.PRIMARY_THEME.SecondaryTextColor,
			Paddding:  1,
		}),
	)

	sidebar.AddComponent(
		dd.CreateComponent("Deez nuts", dd.Styles{
			MaxWidth:   10,
			Background: dd.MakeRGBBackground(250, 0, 0),
			TextColor:  dd.MakeRGBTextColor(0, 0, 0),
			Paddding:   1,
			Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.MakeRGBTextColor(100, 100, 100)},
		},
		),
	)
	screen.AddWindow(sidebar)

	content := dd.CreateWindow(dd.Styles{
		Width:      screen.Width - sidebar.Styles.Width - 1,
		Height:     int(float32(screen.Height)*0.9) + 1,
		Background: dd.PRIMARY_THEME.PrimaryBg,
		Paddding:   1,
		Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.PRIMARY_THEME.SecondaryTextColor},
		Direction:  dd.INLINE,
	})

	content.AddComponent(
		dd.CreateComponent(
			"|SIMD|",
			dd.Styles{
				MaxWidth:   10,
				Background: dd.MakeRGBBackground(80, 40, 100),
				Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.MakeRGBTextColor(100, 100, 100)},
			},
		))
	content.AddComponent(
		dd.CreateComponent(
			"LIGMA???",
			dd.Styles{
				MaxWidth:   10,
				Background: dd.MakeRGBBackground(80, 40, 100),
				Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.MakeRGBTextColor(100, 100, 100)},
			},
		))

	screen.AddWindow(content)

	input_bar := dd.CreateWindow(
		dd.Styles{
			Width:      screen.Width - sidebar.Styles.Width - 1,
			Height:     int(float32(screen.Height)*0.1) - 1,
			Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.PRIMARY_THEME.SecondaryTextColor},
			Background: dd.PRIMARY_THEME.PrimaryBg,
		},
	)

	input_bar.AddComponent(
		dd.CreateComponent(
			"",
			dd.Styles{
				MinWidth:   10,
				MaxWidth:   input_bar.Width - 2,
				Background: dd.MakeRGBBackground(100, 40, 100),
			},
		))

	input_bar.Components[0].Inputable = true
	screen.AddWindow(input_bar)

	screen.StatusBar = dd.Window{
		Position: dd.Position{Row: screen.Height, Col: 1},
		Styles: dd.Styles{
			Width:      screen.Width,
			Height:     1,
			Background: dd.MakeRGBBackground(80, 40, 100),
		},
	}
	screen.StatusBar.Components = []*dd.Component{
		{
			Parent: &screen.StatusBar,
			Buffer: dd.NORMAL,
			Styles: dd.Styles{
				MaxWidth: screen.Width,
				Height:   1,
			},
		},
	}
	screen.Render()
	screen.Activate()
}

func create_login_screen(screen *dd.Screen) {
	width, height, _ := term.GetSize(int(os.Stdin.Fd()))
	screen.Width = width
	screen.Height = height

	login := dd.CreateWindow(
		dd.Styles{
			Width:      40,
			Height:     10,
			Paddding:   1,
			Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.PRIMARY_THEME.SecondaryTextColor},
			Background: dd.PRIMARY_THEME.PrimaryBg,
		},
	)
	login.Position = dd.Position{Row: screen.Height/2 - 5, Col: screen.Width/2 - 20}

	login.AddComponent(
		dd.CreateComponent(
			"",
			dd.Styles{
				MinWidth:   10,
				MaxWidth:   login.Width - 2,
				Background: dd.MakeRGBBackground(100, 40, 100),
			},
		))
	login.AddComponent(
		dd.CreateComponent(
			"",
			dd.Styles{
				MinWidth:   10,
				MaxWidth:   login.Width - 2,
				Background: dd.MakeRGBBackground(100, 40, 100),
			},
		))
	login.AddComponent(
		dd.CreateComponent(
			"connect",
			dd.Styles{
				MinWidth:   10,
				MaxWidth:   login.Width / 2,
				Background: dd.MakeRGBBackground(100, 40, 100),
				Border:     dd.Border{Style: dd.RoundedBorder, Color: dd.PRIMARY_THEME.SecondaryBg},
			},
		))

	login.Components[0].Inputable = true
	login.Components[1].Inputable = true
	screen.AddWindow(login)

	screen.StatusBar = dd.Window{
		Position: dd.Position{Row: screen.Height, Col: 1},
		Styles: dd.Styles{
			Width:      screen.Width,
			Height:     1,
			Background: dd.MakeRGBBackground(80, 40, 100),
		},
	}
	screen.StatusBar.Components = []*dd.Component{
		{
			Parent: &screen.StatusBar,
			Buffer: dd.NORMAL,
			Styles: dd.Styles{
				MaxWidth: screen.Width,
				Height:   1,
			},
		},
	}
	screen.Render()
	screen.Activate()
}

func main() {
	old_state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error enabling raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), old_state)

	dd.ClearScreen()
	screen := dd.Screen{State: &dd.Normal, EventLoopIsRunning: true}

	client := tpc.Client{}

	if !client.LoadClient() {
		create_login_screen(&screen)
	} else {
		create_main_window(&screen)
	}

	// TODO: Check if its possible to accept more than one byte
	stdin_buffer := make([]byte, 1)
	for screen.EventLoopIsRunning {
		for len(screen.RenderQueue) > 0 {
			item_to_render := screen.RenderQueue[0]
			fmt.Print(item_to_render)
			screen.RenderQueue = screen.RenderQueue[1:]
		}

		fmt.Printf(dd.MOVE_CURSOR_TO_POSITION, screen.CursorPosition.Row, screen.CursorPosition.Col)

		_, err := os.Stdin.Read(stdin_buffer)
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		screen.State.HandleKeypress(&screen, stdin_buffer)
	}
	// restart to default settings
	fmt.Print(dd.SHOW_CURSOR)
	// TODO: any assert should have show cursor
}
