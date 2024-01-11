/*
	A Labeled Combobox widget:

	-----------------------------
	|            |              |
	|   Food:    |  <dropdown>  |
	|            |              |
	-----------------------------

	The user can then edit the dropdown field and type anything into it
*/
package gadgets

import 	(
	"reflect"
	"strconv"

	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

type BasicCombobox struct {
	ready	bool
	name	string

	parent	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	d	*gui.Node	// dropdown widget

	value	string
	label	string

	values map[string]string

	Custom func()
}

func (d *BasicCombobox) Get() string {
	if ! d.Ready() {return ""}
	return d.value
}

// Returns true if the status is valid
func (d *BasicCombobox) Ready() bool {
	if d == nil {return false}
	return d.ready
}

func (d *BasicCombobox) Enable() {
	if d == nil {return}
	if d.d == nil {return}
	d.d.Enable()
}

func (d *BasicCombobox) Disable() {
	if d == nil {return}
	if d.d == nil {return}
	d.d.Disable()
}

func (d *BasicCombobox) Add(value any) {
	if ! d.Ready() {return}
	log.Log(INFO, "BasicCombobox.Add() =", value)

	var b reflect.Kind
	b = reflect.TypeOf(value).Kind()

	switch b {
	case reflect.Int:
		var i int
		i = value.(int)
		s := strconv.Itoa(i)
		if d.values[s] != "added" {
			d.values[s] = "added"
			d.d.AddDropdownName(s)
		}
	case reflect.String:
		s := value.(string)
		if d.values[s] != "added" {
			d.values[s] = "added"
			d.d.AddDropdownName(s)
		}
	case reflect.Bool:
		if value.(bool) == true {
			s := "true"
			if d.values[s] != "added" {
				d.values[s] = "added"
				d.d.AddDropdownName(s)
			}
		} else {
			s := "false"
			if d.values[s] != "added" {
				d.values[s] = "added"
				d.d.AddDropdownName(s)
			}
		}
	default:
	}
}

func (d *BasicCombobox) Set(value any) bool {
	if ! d.Ready() {return false}
	log.Log(INFO, "BasicCombobox.Set() =", value)

	var b reflect.Kind
	b = reflect.TypeOf(value).Kind()

	switch b {
	case reflect.Int:
		var i int
		i = value.(int)
		s := strconv.Itoa(i)
		d.d.SetText(s)
		d.value = s
	case reflect.String:
		d.d.SetText(value.(string))
		d.value = value.(string)
	case reflect.Bool:
		if value.(bool) == true {
			d.d.SetText("true")
			d.value = "true"
		} else {
			d.d.SetText("false")
			d.value = "false"
		}
	default:
		return false
	}
	return true
}

func NewBasicCombobox(p *gui.Node, name string) *BasicCombobox {
	d := BasicCombobox {
		parent: p,
		name: name,
		ready: false,
	}

	// various timeout settings
	d.l = p.NewLabel(name)
	d.d = p.NewCombobox("")
	d.d.Custom = func() {
		d.value = d.d.GetText()
		log.Warn("BasicCombobox.Custom() user changed value to =", d.value)
		if d.Custom != nil {
			d.Custom()
		}
	}
	d.values = make(map[string]string)
	d.ready = true
	return &d
}
