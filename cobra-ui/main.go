package cobraUi

import (
	"fyne.io/fyne/widget"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

type CobraUi struct {
	Command *cobra.Command
	// Window		*fyne.Window
	Form *widget.Form
}

const (
	RunCommandText = `Run Command`
)

func CommandNameList(commands []*cobra.Command) []string {
	var list []string
	for _, command := range commands {
		list = append(list, command.Use)
	}
	return list
}
func (cobraui *CobraUi) GetChildrenNames() []string {
	var list []string
	for _, command := range cobraui.GetChildrenCommands() {
		list = append(list, command.Use)
	}
	return list
}

// NewForm generates a new Command Form
func (cobraui *CobraUi) NewForm() *widget.Form {

	form := &widget.Form{}
	// tt := reflect.TypeOf(x).Elem()
	commands := make([]*cobra.Command, 0)
	commands = cobraui.GetChildrenCommands()

	widget.NewSelect(CommandNameList(commands), ChangeCommandSelect)

	// for i := 0; i < tt.NumField(); i++ {
	// 	fld := tt.Field(i)
	// 	tag := fld.Tag.Get("json")
	// 	switch tag {
	// 	case "": // not a display field
	// 	case "img": // special field for images
	// 		we created this in the setup
	// case "num": // special field for ID
	// 	entry := widget.NewEntry()
	// 	x.iDEntry = entry
	// 	form.Append(fld.Name, entry)
	// default:
	// 	form.Append(fld.Name, x.newLabel(tag))
	// }
	// }
	cobraui.Form = form
	return form
}

func (cobraui *CobraUi) GetChildrenCommands() []*cobra.Command {
	cmd := cobraui.Command
	return cmd.Commands()
}

// func GetChildrenCommands(commands []*cobra.Command) {
// 	for _, command := range cobraui.GetChildrenCommands() {
// 		// form.Items = append(form.Items, widget.NewFormItem(command.Use, fyne.))
// 		commands := append(commands, command)
// 	}
// }

func ChangeCommandSelect(changedTo string) {
	util.Logger().Info("ChangeCommandSelect to " + util.ColorInfo(changedTo))
}

func (cobraui *CobraUi) CreateRunCommandButton(addToForm bool) *widget.Button {
	submit := widget.NewButton(RunCommandText, func() {
		if err := cobraui.Execute(); err != nil {
			util.Logger().Fatalf("executing command failed: %s", err)
		}
	})
	submit.Style = widget.PrimaryButton
	if addToForm {
		err := cobraui.AddButtonToForm(submit)
		if err != nil {
			util.Logger().Fatalf("trying to add button to form failed: %s", err)
		}
	}
	return submit
}

func (cobraui *CobraUi) AddButtonToForm(button *widget.Button) error {
	cobraui.Form.AppendItem(&widget.FormItem{
		Text: button.Text,
	})
	// TODO
	return util.ErrorUnimplemented()

}

func (cobraui *CobraUi) SetArgs() error {
	return util.ErrorUnimplemented()
}

func (cobraui *CobraUi) Execute() error {
	// TODO this will require args to be set before running, based on the form above.

	err := cobraui.SetArgs()
	if err != nil {
		return err
	}
	return cobraui.Command.Execute()
}
