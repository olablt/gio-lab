package ui

import (
	"fmt"
	"image"
	"log"
	"strings"

	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type CommandPalette struct {
	SearchInput      *widget.Editor
	List             *widget.List
	Commands         []Command
	CommandsFiltered []Command
	cursor           int
	// commands to callbacks
	callbacks map[string]func()
	OnSubmit  func()
	OnCancel  func()
	// commands to key strings
	shortcutStrings map[string]string
	// keys to commands
	keys map[key.Filter]string
	// commands to clickable widgets
	clickables map[string]*widget.Clickable
	Visible    bool
	//
	KeyPress       bool
	Key            key.Name
	ClickableLayer *widget.Clickable
}

type Command struct {
	Category string
	Name     string
	Func     func()
	Key      key.Filter
	Icon     *widget.Icon
}

// NewCommandPalette creates a new command palette with default settings
func NewCommandPalette() *CommandPalette {
	cp := &CommandPalette{
		SearchInput:      &widget.Editor{SingleLine: true, Submit: true, Alignment: text.Start},
		List:             &widget.List{List: layout.List{Axis: layout.Vertical}},
		Commands:         []Command{},
		CommandsFiltered: []Command{},
		cursor:           -1,
		callbacks:        make(map[string]func()),
		shortcutStrings:  make(map[string]string),
		keys:             make(map[key.Filter]string),
		clickables:       make(map[string]*widget.Clickable),
		ClickableLayer:   &widget.Clickable{},
	}
	return cp
}

// Modify RegisterCommand to accept a Command struct
// RegisterCommand adds a new command to the palette that can be searched and executed
func (cp *CommandPalette) RegisterCommand(cmd Command) {
	cp.Commands = append(cp.Commands, cmd)
	cp.CommandsFiltered = cp.Commands
	cp.callbacks[cmd.Name] = cmd.Func
	cp.clickables[cmd.Name] = &widget.Clickable{}
	if cmd.Key.Name != "" {
		cp.keys[cmd.Key] = cmd.Name
		cp.shortcutStrings[cmd.Name] = fmt.Sprintf("%v %v", cmd.Key.Required, cmd.Key.Name)
	}
}

// UpdateCommands filters the command list based on search text
// If there's a category (text before ':'), it filters by category first
// Then it uses fuzzy search to find matching command names
func (cp *CommandPalette) UpdateCommands(selectFirst bool) {
	searchText := cp.SearchInput.Text()
	searchText = strings.TrimSpace(searchText)

	// Check if search contains colon for category filtering
	colonIdx := strings.Index(searchText, ":")
	if colonIdx >= 0 {
		category := searchText[:colonIdx]
		searchAfterColon := searchText[colonIdx+1:]

		// First filter by category
		categoryFiltered := []Command{}
		for _, cmd := range cp.Commands {
			if strings.EqualFold(cmd.Category, category) {
				categoryFiltered = append(categoryFiltered, cmd)
			}
		}

		// Then apply fuzzy search on names only (not categories)
		cp.CommandsFiltered = []Command{}
		names := make([]string, len(categoryFiltered))
		for i, cmd := range categoryFiltered {
			names[i] = cmd.Name // Search only names when category filtered
		}
		matchingNames := fuzzy.FindNormalizedFold(searchAfterColon, names)
		for _, name := range matchingNames {
			for _, cmd := range categoryFiltered {
				if cmd.Name == name {
					cp.CommandsFiltered = append(cp.CommandsFiltered, cmd)
				}
			}
		}
	} else {
		// Normal search without category filtering - search both category and name
		cp.CommandsFiltered = []Command{}
		searchTerms := make([]string, len(cp.Commands))
		for i, cmd := range cp.Commands {
			searchTerms[i] = fmt.Sprintf("%s: %s", cmd.Category, cmd.Name)
		}
		matchingTerms := fuzzy.FindNormalizedFold(searchText, searchTerms)
		for _, term := range matchingTerms {
			for _, cmd := range cp.Commands {
				if fmt.Sprintf("%s: %s", cmd.Category, cmd.Name) == term {
					cp.CommandsFiltered = append(cp.CommandsFiltered, cmd)
				}
			}
		}
	}

	if selectFirst {
		cp.cursor = 0
	}
}

// SetCallback changes what happens when a specific command is executed
func (cp *CommandPalette) SetCallback(command string, callback func()) {
	// check if the command exists
	if _, ok := cp.callbacks[command]; !ok {
		log.Printf("[CP] command '%v' does not exist", command)
		return
	}
	cp.callbacks[command] = callback
}

// Call executes the function associated with a command
func (cp *CommandPalette) Call(command string) {
	// check if the command exists
	if call, ok := cp.callbacks[command]; !ok || call == nil {
		log.Printf("[CP] command '%v' does not exist", command)
		return
	}
	cp.callbacks[command]()
}

// submit executes a command and resets the palette
func (cp *CommandPalette) submit(command string) {
	// log.Printf("[CP] SUBMIT '%v'", cp.CommandsFiltered[cp.cursor])
	log.Printf("[CP] SUBMIT '%v'", command)
	// if the callback exists, call it
	cp.Call(command)
	// if OnSubmit is set, call it
	if cp.OnSubmit != nil {
		cp.OnSubmit()
	}
	cp.Reset()
}

// HandleShortcutKeys checks for keyboard shortcuts and executes their commands
func (cp *CommandPalette) HandleShortcutKeys(gtx layout.Context) {
	if cp.Visible {
		return
	}
	// tag := &cp.SearchInput
	// event.Op(gtx.Ops, tag)
	filters := []event.Filter{}
	for key, command := range cp.keys {
		_ = command
		if key.Name != "" {
			// log.Printf("[CP] adding key: %v %v to filters", v.Required, v.Name)
			filters = append(filters, key)
			// filters = append(filters, key.Filter{Name: v.Name, Required: v.Required})
		}
	}
	// check for new key events
	for {
		event, ok := gtx.Event(filters...)
		if !ok {
			break
		}
		ev, ok := event.(key.Event)
		if !ok {
			continue
		}
		// handle ev
		if ev.State == key.Press {
			// log.Printf("[CP] got shortcut %v %v", ev.Modifiers, ev.Name)
			filter := key.Filter{Name: ev.Name, Required: ev.Modifiers}
			if command, ok := cp.keys[filter]; ok {
				cp.submit(command)
			}
		}
	}
}

// In ProcessPointerEvents method
// ProcessPointerEvents handles mouse clicks on the command palette
func (cp *CommandPalette) ProcessPointerEvents(gtx layout.Context) {
	if cp.ClickableLayer.Clicked(gtx) {
		cp.Reset()
	}
	// loop through filtered list and check for clicks
	for _, command := range cp.CommandsFiltered {
		if cp.clickables[command.Name].Clicked(gtx) {
			cp.submit(command.Name) // Use command.Name instead of command
		}
	}
}

// Reset clears the search and hides the command palette
func (cp *CommandPalette) Reset() {
	cp.cursor = -1
	// log.Println("cp.cursor", cp.cursor)
	cp.SearchInput.SetText("")
	cp.CommandsFiltered = cp.Commands
	cp.Visible = false
}

// Show makes the command palette visible with initial search text
func (cp *CommandPalette) Show(txt string) {
	cp.SearchInput.SetText(txt)
	cp.SearchInput.SetCaret(20, 20)
	cp.UpdateCommands(false)
	cp.cursor = -1
	// cp.CommandsFiltered = cp.Commands
	cp.Visible = true
}

// ProcessKeyEvents handles keyboard input for navigation and selection
func (cp *CommandPalette) ProcessKeyEvents(gtx layout.Context) {
	// tag := &cp.SearchInput
	// event.Op(gtx.Ops, tag)

	// handle key events
	filters := []event.Filter{
		key.Filter{Name: "↑"},
		key.Filter{Name: "↓"},
		key.Filter{Name: "J", Required: key.ModCtrl},
		key.Filter{Name: "K", Required: key.ModCtrl},
		key.Filter{Name: key.NameReturn},
		key.Filter{Name: key.NameEscape},
		// key.FocusFilter{Target: tag},
		// key.Filter{Focus: tag, Name: "↑"},
	}
	// check for new key events
	cp.KeyPress = false
	cp.Key = ""
	for {
		event, ok := gtx.Event(filters...)
		if !ok {
			break
		}
		ev, ok := event.(key.Event)
		if !ok {
			continue
		}
		// handle ev
		if ev.State == key.Press {
			log.Printf("[CP] got key.%v", ev.Name)
			cp.KeyPress = true
			cp.Key = ev.Name

			// handle enter
			if ev.Name == key.NameReturn {
				if cp.cursor >= 0 {
					cp.submit(cp.CommandsFiltered[cp.cursor].Name) // Use .Name here
					cp.Reset()                                     // first submit and then reset
				}
			}
			// handle escape
			if ev.Name == key.NameEscape {
				log.Println("[CP] escape pressed")
				cp.Reset()
				if cp.OnCancel != nil {
					cp.OnCancel()
				}
			}
			// cursor movement
			if ev.Name == "↓" || ev.Name == "J" {
				cp.cursor = cp.cursor + 1
				if cp.cursor > len(cp.CommandsFiltered)-1 {
					cp.cursor = len(cp.CommandsFiltered) - 1
				}
			}
			if ev.Name == "↑" || ev.Name == "K" {
				cp.cursor = cp.cursor - 1
				if cp.cursor < 0 {
					cp.cursor = -1
				}
			}
		}
	}

	// process input events
	inputUpdated := false
	for {
		ev, ok := cp.SearchInput.Update(gtx)
		if !ok {
			break
		}
		// log.Println("got event", ev, reflect.TypeOf(ev))
		_, ok = ev.(widget.ChangeEvent)
		if ok {
			// log.Println("got widget.ChangeEvent")
			inputUpdated = true
		}
	}

	if inputUpdated {
		cp.UpdateCommands(true)
	}
}

// SetCursor changes which item is highlighted in the command list
func (cp *CommandPalette) SetCursor(i int) {
	cp.cursor = i
}

// Update processes all events (keyboard, mouse, etc) for the command palette
func (cp *CommandPalette) Update(gtx layout.Context) {
	cp.HandleShortcutKeys(gtx)
	// process pointer events
	cp.ProcessPointerEvents(gtx)
	// process key events
	cp.ProcessKeyEvents(gtx)
}

// Layout draws the command palette on screen
func (cp *CommandPalette) Layout(gtx layout.Context, th *material.Theme) D {
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			// semi transparent background
			return cp.ClickableLayer.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				dims := ColorBox(gtx, gtx.Constraints.Min, Alpha(BgColor, 250))
				return dims
			})
		},
		func(gtx layout.Context) layout.Dimensions {
			w := gtx.Dp(500)
			h := gtx.Dp(300)
			gtx.Constraints.Max = image.Pt(w, h)
			gtx.Constraints.Min = image.Pt(w, h)
			// fill the background
			paint.FillShape(gtx.Ops, th.Palette.Bg, clip.Rect{Max: gtx.Constraints.Min}.Op())
			Rows(
				Rigid(func(gtx layout.Context) layout.Dimensions {
					return cp.InputLayout(gtx, th)

				}),
				Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return cp.ListLayout(gtx, th)
				}),
			)(gtx)
			return layout.Dimensions{Size: image.Point{w, h}}
		})

}

// InputLayout draws the search input box
func (cp *CommandPalette) InputLayout(gtx C, th *material.Theme) D {
	// layout
	margins := layout.UniformInset(unit.Dp(5))
	return margins.Layout(gtx,
		TextInput(cp.SearchInput, "Text Input"),
	)
}

// ListLayout draws the filtered list of commands
func (cp *CommandPalette) ListLayout(gtx C, th *material.Theme) D {
	margins := layout.Inset{Top: unit.Dp(0), Right: unit.Dp(0), Bottom: unit.Dp(5), Left: unit.Dp(5)}
	// Check if we're in category filter mode
	searchText := strings.TrimSpace(cp.SearchInput.Text())
	inCategoryMode := strings.Contains(searchText, ":")
	return material.List(th, cp.List).Layout(gtx, len(cp.CommandsFiltered), func(gtx C, i int) D {
		return margins.Layout(gtx,
			func(gtx C) D {
				th2 := *th
				if i == cp.cursor {
					th2.Palette.Bg = th2.Palette.ContrastBg
					th2.Palette.Fg = th.Palette.ContrastFg
				}
				cmd := cp.CommandsFiltered[i]
				displayName := cmd.Name
				if !inCategoryMode {
					displayName = fmt.Sprintf("%s: %s", cmd.Category, cmd.Name)
				}
				return ActionListItem(&th2, cp.clickables[cmd.Name], displayName, cp.shortcutStrings[cmd.Name]).Layout(gtx)
			},
		)
	})
}
