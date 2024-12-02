package cpalette

import (
	"fmt"
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
	SearchInput        *widget.Editor
	List               *widget.List
	StringList         []string
	StringListFiltered []string
	cursor             int
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
}

type Command struct {
	Name string
	Func func()
	Key  key.Filter
	Icon *widget.Icon
}

func NewCommandPalette() *CommandPalette {
	cp := &CommandPalette{
		SearchInput:        &widget.Editor{SingleLine: true, Submit: true, Alignment: text.Start},
		List:               &widget.List{List: layout.List{Axis: layout.Vertical}},
		StringList:         []string{},
		StringListFiltered: []string{},
		cursor:             0,
		callbacks:          make(map[string]func()),
		shortcutStrings:    make(map[string]string),
		keys:               make(map[key.Filter]string),
		clickables:         make(map[string]*widget.Clickable),
	}
	cp.StringList = []string{}
	cp.StringListFiltered = []string{}
	return cp
}

// this will add a new command to the list of commands
// that the command palette will display
// command string must be unique
func (cp *CommandPalette) RegisterCommand(command string, callback func(), key key.Filter) {
	cp.StringList = append(cp.StringList, command)
	cp.StringListFiltered = cp.StringList
	cp.callbacks[command] = callback
	cp.clickables[command] = &widget.Clickable{}
	if key.Name != "" {
		// log.Printf("[CP] adding key: %v %v to shortcuts", v.Required, v.Name)
		cp.keys[key] = command
		cp.shortcutStrings[command] = fmt.Sprintf("%v %v", key.Required, key.Name)
	}

}

// SetCallback will set the callback for a command
func (cp *CommandPalette) SetCallback(command string, callback func()) {
	// check if the command exists
	if _, ok := cp.callbacks[command]; !ok {
		log.Printf("[CP] command '%v' does not exist", command)
		return
	}
	cp.callbacks[command] = callback
}

func (cp *CommandPalette) InputLayout(gtx C, th *material.Theme) D {
	// layout
	margins := layout.UniformInset(unit.Dp(5))
	return margins.Layout(gtx,
		func(gtx C) D {
			// Wrap the editor in material design
			ed := material.Editor(th, cp.SearchInput, "Search")
			return ed.Layout(gtx)
		},
	)
}

func (cp *CommandPalette) ListLayout(gtx C, th *material.Theme) D {
	// Define insets for the list items
	margins := layout.Inset{Top: unit.Dp(0), Right: unit.Dp(0), Bottom: unit.Dp(5), Left: unit.Dp(5)}
	// layout the list
	return material.List(th, cp.List).Layout(gtx, len(cp.StringListFiltered), func(gtx C, i int) D {
		return margins.Layout(gtx,
			func(gtx C) D {
				th2 := *th
				// th2.Font.Size = unit.Dp(16)
				if i == cp.cursor {
					th2.Palette.Bg = th2.Palette.ContrastBg
					th2.Palette.Fg = th.Palette.ContrastFg
				}
				command := cp.StringListFiltered[i]
				return ActionListItem(&th2, cp.clickables[command], command, cp.shortcutStrings[command]).Layout(gtx)
				// return IconActionListItem(&th2, cp.clickables[command], icons.ContentSave, command).Layout(gtx)
			},
		)
	})
}

// func (cp *CommandPalette) submit(i int) {
func (cp *CommandPalette) submit(command string) {
	// log.Printf("[CP] SUBMIT '%v'", cp.StringListFiltered[cp.cursor])
	log.Printf("[CP] SUBMIT '%v'", command)
	// check if the callback exists and call it
	if callback, ok := cp.callbacks[command]; ok && callback != nil {
		callback()
	}
	if cp.OnSubmit != nil {
		cp.OnSubmit()
	}
}

// handle shortcut keys
func (cp *CommandPalette) HandleShortcutKeys(gtx layout.Context) {
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
				// cp.submit(cp.cursor)
				// // log.Printf("[CP] found command for shortcut %v %v: %v", ev.Modifiers, ev.Name, command)
				// // check if the callback exists and call it
				// if callback, ok := cp.callbacks[command]; ok && callback != nil {
				// 	callback()
				// }
			}
		}
	}
}

// ProcessPointerEvents will process pointer events
func (cp *CommandPalette) ProcessPointerEvents(gtx layout.Context) {
	// loop through filtered list and check for clicks
	for _, command := range cp.StringListFiltered {
		if cp.clickables[command].Clicked(gtx) {
			cp.submit(command)
		}
	}
}

// ProcessKeyEvents will process key events
func (cp *CommandPalette) ProcessKeyEvents(gtx layout.Context) {
	tag := &cp.SearchInput
	event.Op(gtx.Ops, tag)

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
		// key.Filter{Focus: tag, Name: "↓"},
		// key.Filter{Focus: tag, Name: "J", Required: key.ModCtrl},
		// key.Filter{Focus: tag, Name: "K", Required: key.ModCtrl},
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
			log.Printf("[CP] got key.%v", ev.Name)
			// handle enter
			if ev.Name == key.NameReturn {
				// cp.submit(cp.cursor)
				cp.submit(cp.StringListFiltered[cp.cursor])
			}
			// handle escape
			if ev.Name == key.NameEscape {
				log.Println("[CP] escape pressed")
				if cp.OnCancel != nil {
					cp.OnCancel()
				}
			}
			// cursor movement
			if ev.Name == "↓" || ev.Name == "J" {
				cp.cursor = cp.cursor + 1
			}
			if ev.Name == "↑" || ev.Name == "K" {
				cp.cursor = cp.cursor - 1
			}
			if cp.cursor < 0 {
				cp.cursor = 0
			}
			if cp.cursor > len(cp.StringListFiltered)-1 {
				cp.cursor = len(cp.StringListFiltered) - 1
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
			inputUpdated = true
		}
	}

	if inputUpdated {
		trimmedString := strings.TrimSpace(cp.SearchInput.Text())
		cp.StringListFiltered = fuzzy.FindNormalizedFold(trimmedString, cp.StringList)
		cp.cursor = 0
		// log.Println("input changed!", trimmedString)
	}
}

func (cp *CommandPalette) Update(gtx layout.Context) {
	// process pointer events
	cp.ProcessPointerEvents(gtx)

	// process key events
	cp.ProcessKeyEvents(gtx)
}

func (cp *CommandPalette) Layout(gtx layout.Context, th *material.Theme) D {
	gtx.Execute(key.FocusCmd{Tag: cp.SearchInput})
	// process events
	cp.HandleShortcutKeys(gtx)
	cp.Update(gtx)

	// defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
	// paint.Fill(gtx.Ops, th.Palette.Bg)
	// // layout elements
	// d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
	// 	layout.Rigid(func(gtx C) D {
	// 		return cp.InputLayout(gtx, th)
	// 	}),
	// 	layout.Flexed(1, func(gtx layout.Context) D {
	// 		return cp.ListLayout(gtx, th)
	// 	}),
	// )
	// return d

	// apply background
	return layout.Background{}.Layout(gtx,
		func(gtx C) D {
			defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, th.Palette.Bg)
			return D{Size: gtx.Constraints.Min}
		}, func(gtx C) D {
			// layout elements
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return cp.InputLayout(gtx, th)
				}),
				layout.Flexed(1, func(gtx layout.Context) D {
					return cp.ListLayout(gtx, th)
				}),
			)
		})

}
