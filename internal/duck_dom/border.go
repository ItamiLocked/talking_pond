package duckdom

import (
	"fmt"
	"strings"
)

type Border struct {
	Style BorderStyle
	Color string
}

type BorderRenderStyles struct {
	Width int
	Height int
	BorderBackground string
	Border
}

type BorderStyle struct {
	Top          string
	Bottom       string
	Left         string
	Right        string
	TopLeft      string
	TopRight     string
	BottomLeft   string
	BottomRight  string
	MiddleLeft   string
	MiddleRight  string
	Middle       string
	MiddleTop    string
	MiddleBottom string
}

var NoBorder = Border{}

var (
	NormalBorder = BorderStyle{
		Top:          "─",
		Bottom:       "─",
		Left:         "│",
		Right:        "│",
		TopLeft:      "┌",
		TopRight:     "┐",
		BottomLeft:   "└",
		BottomRight:  "┘",
		MiddleLeft:   "├",
		MiddleRight:  "┤",
		Middle:       "┼",
		MiddleTop:    "┬",
		MiddleBottom: "┴",
	}

	BoldBorder = BorderStyle{
		Top:          "━",
		Bottom:       "━",
		Left:         "┃",
		Right:        "┃",
		TopLeft:      "┏",
		TopRight:     "┓",
		BottomLeft:   "┗",
		BottomRight:  "┛",
		MiddleLeft:   "┣",
		MiddleRight:  "┫",
		Middle:       "╋",
		MiddleTop:    "┳",
		MiddleBottom: "┻",
	}

	RoundedBorder = BorderStyle{
		Top:          "─",
		Bottom:       "─",
		Left:         "│",
		Right:        "│",
		TopLeft:      "╭",
		TopRight:     "╮",
		BottomLeft:   "╰",
		BottomRight:  "╯",
		MiddleLeft:   "├",
		MiddleRight:  "┤",
		Middle:       "┼",
		MiddleTop:    "┬",
		MiddleBottom: "┴",
	}
)

// box-sizing: border-box;
// We hate borders
func render_border(border_builder *strings.Builder, position Position, styles BorderRenderStyles) {
	if styles.Border == NoBorder {
		return
	}
	border_style := styles.Border.Style

	middle := strings.Repeat(border_style.Bottom, styles.Width-2)
	top := border_style.TopLeft + middle + border_style.TopRight
	bottom := border_style.BottomLeft + middle + border_style.BottomRight

	border_builder.WriteString(styles.Border.Color)
	border_builder.WriteString(styles.BorderBackground)
	border_builder.WriteString(fmt.Sprintf(MOVE_CURSOR_TO_POSITION, position.Row, position.Col))
	border_builder.WriteString(top)

	for i := 1; i < styles.Height-1; i += 1 {
		left_wall := fmt.Sprintf(MOVE_CURSOR_TO_POSITION, position.Row+i, position.Col)
		right_wall := fmt.Sprintf(MOVE_CURSOR_TO_POSITION, position.Row+i, position.Col+styles.Width-1)
		wall := left_wall + border_style.Left + right_wall + border_style.Right
		border_builder.WriteString(wall)
	}
	border_builder.WriteString(fmt.Sprintf(MOVE_CURSOR_TO_POSITION, styles.Height+position.Row-1, position.Col))
	border_builder.WriteString(bottom)
	border_builder.WriteString(RESET_STYLES)
}
