package duckdom

import (
	"fmt"
	"strings"
)

type Component struct {
	Position
	Styles
	Content string
	Buffer  string
	// NOTE: we should really think about it
	// maybe it would be better if we just made a bunch of functions
	// that take *Component as an input and does some actions with it
	ChildComponents []Component
	Action func()
}

func (self *Component) ExecuteAction() {
	self.Action()
}

// FIXME: NODARIO definetely needs a name or logic change
// FIXME: The fuck you mean?
func (self *Component) Render() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf(MOVE_CURSOR_TO_POSITION, self.Position.StartingRow, self.Position.StartingCol))
	builder.WriteString(self.Styles.Compile())
	builder.WriteString(self.Buffer)
	builder.WriteString(RESET_STYLES)
	self.Content = builder.String()
	return self.Content
}
