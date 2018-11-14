package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

type areaHandler struct{}

func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}

func makeServerPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	ipEntry := ui.NewEntry()
	ipEntry.SetReadOnly(true)
	vbox.Append(ipEntry, false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	tokenEntry := ui.NewEntry()
	tokenEntry.SetReadOnly(true)
	vbox.Append(tokenEntry, false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	passEntry := ui.NewEntry()
	passEntry.SetReadOnly(true)
	vbox.Append(passEntry, false)

	button := ui.NewButton("Open File")
	button.OnClicked(func(*ui.Button) {
	})
	vbox.Append(button, false)
	return vbox
}

func makeClientPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	ipEntry := ui.NewEntry()
	ipEntry.SetReadOnly(true)
	vbox.Append(ipEntry, false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	tokenEntry := ui.NewEntry()
	tokenEntry.SetReadOnly(true)
	vbox.Append(tokenEntry, false)

	return vbox
}

func setupUI() {

	mainwin := ui.NewWindow("goP2PVPN", 640, 480, true)
	mainwin.SetMargined(true)
	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Server", makeServerPage())
	tab.SetMargined(0, true)

	tab.Append("Client", makeClientPage())
	tab.SetMargined(1, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
