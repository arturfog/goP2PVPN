package main

import (
	"./modules/vpn"
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
	vps := vpn.NewVPNServer()

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Server address:"), false)
	ipEntry := ui.NewEntry()
	vbox.Append(ipEntry, false)

	vbox.Append(ui.NewLabel("Token:"), false)
	tokenEntry := ui.NewEntry()
	tokenEntry.SetReadOnly(true)
	vbox.Append(tokenEntry, false)

	tokenButton := ui.NewButton("Generate token")
	tokenButton.OnClicked(func(*ui.Button) {
		key := vps.GenKey()
		tokenEntry.SetText(key)
	})
	vbox.Append(tokenButton, false)

	vbox.Append(ui.NewLabel("Password:"), false)
	passEntry := ui.NewEntry()
	vbox.Append(passEntry, false)

	connectButton := ui.NewButton("Connect")
	connectButton.OnClicked(func(*ui.Button) {
		vps.Connect(ipEntry.Text(), tokenEntry.Text())
	})
	vbox.Append(connectButton, false)

	vbox.Append(ui.NewLabel("Status:"), false)
	connectionStatus := ui.NewEntry()
	connectionStatus.SetReadOnly(true)
	vbox.Append(connectionStatus, false)

	return vbox
}

func makeClientPage() ui.Control {
	vpc := vpn.NewVPNClient()

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Server address:"), false)
	ipEntry := ui.NewEntry()
	vbox.Append(ipEntry, false)

	vbox.Append(ui.NewLabel("Token:"), false)
	tokenEntry := ui.NewEntry()
	vbox.Append(tokenEntry, false)

	vbox.Append(ui.NewLabel("Password:"), false)
	passEntry := ui.NewEntry()
	vbox.Append(passEntry, false)

	button := ui.NewButton("Connect")
	button.OnClicked(func(*ui.Button) {
		vpc.Connect(ipEntry.Text(), tokenEntry.Text())
	})
	vbox.Append(button, false)

	vbox.Append(ui.NewLabel("Status:"), false)
	connectionStatus := ui.NewEntry()
	connectionStatus.SetReadOnly(true)
	vbox.Append(connectionStatus, false)

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

	tab.Append("Settings", makeClientPage())
	tab.SetMargined(2, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
