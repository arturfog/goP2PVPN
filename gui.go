package main

import (
	"strings"

	"./modules/vpn"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window
var vps *vpn.VPNServer
var vpc *vpn.VPNClient

func makeServerPage() ui.Control {
	vps = vpn.NewVPNServer()

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Server address:"), false)
	ipEntry := ui.NewEntry()
	ipEntry.SetText("127.0.0.1:8081")
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
		vps.Connect(strings.Trim(ipEntry.Text(), "\x00"), strings.Trim(tokenEntry.Text(), "\x00"))
	})
	vbox.Append(connectButton, false)

	vbox.Append(ui.NewLabel("Status:"), false)
	connectionStatus := ui.NewMultilineEntry()
	connectionStatus.SetReadOnly(true)
	vbox.Append(connectionStatus, false)

	return vbox
}

func makeClientPage() ui.Control {
	vpc = vpn.NewVPNClient()

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Server address:"), false)
	ipEntry := ui.NewEntry()
	ipEntry.SetText("127.0.0.1:8081")
	vbox.Append(ipEntry, false)

	vbox.Append(ui.NewLabel("Token:"), false)
	tokenEntry := ui.NewEntry()
	vbox.Append(tokenEntry, false)

	vbox.Append(ui.NewLabel("Password:"), false)
	passEntry := ui.NewEntry()
	vbox.Append(passEntry, false)

	button := ui.NewButton("Connect")
	button.OnClicked(func(*ui.Button) {
		vpc.Connect(strings.Trim(ipEntry.Text(), "\x00"), strings.Trim(tokenEntry.Text(), "\x00"))
	})
	vbox.Append(button, false)

	vbox.Append(ui.NewLabel("Status:"), false)
	connectionStatus := ui.NewMultilineEntry()
	connectionStatus.SetReadOnly(true)
	vbox.Append(connectionStatus, false)

	return vbox
}

func makeShellPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Shell:"), false)
	shellOutput := ui.NewMultilineEntry()
	shellOutput.SetReadOnly(true)
	vbox.Append(shellOutput, false)

	clrBtn := ui.NewButton("Clear")
	vbox.Append(clrBtn, false)

	vbox.Append(ui.NewLabel("Input:"), false)
	shellInput := ui.NewEntry()
	vbox.Append(shellInput, false)

	execBtn := ui.NewButton("Execute")
	execBtn.OnClicked(func(*ui.Button) {
		command := strings.Trim(shellInput.Text(), "\x00")
		cmdBytes := []byte(command)
		bytesToSend := []byte{vpn.CMD_EXEC_SHELL}
		bytesToSend = append(bytesToSend, cmdBytes...)

		vpc.Conn.WriteToUDP(bytesToSend, vpc.Peer)
	})
	vbox.Append(execBtn, false)

	return vbox
}

func makeTransferPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Upload:"), false)
	uploadEntry := ui.NewEntry()
	uploadEntry.SetReadOnly(true)
	vbox.Append(uploadEntry, false)

	selectBtn := ui.NewButton("Select file ...")
	selectBtn.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		uploadEntry.SetText(filename)
	})
	vbox.Append(selectBtn, false)

	upBtn := ui.NewButton("Upload")
	vbox.Append(upBtn, false)

	vbox.Append(ui.NewLabel("Download:"), false)
	dlEntry := ui.NewEntry()
	vbox.Append(dlEntry, false)

	downBtn := ui.NewButton("Download")
	vbox.Append(downBtn, false)

	return vbox
}

func makeProxyPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	enableBtn := ui.NewButton("Enable")
	vbox.Append(enableBtn, false)

	disableBtn := ui.NewButton("Disable")
	vbox.Append(disableBtn, false)

	vbox.Append(ui.NewLabel("Status:"), false)
	statusOutput := ui.NewMultilineEntry()
	statusOutput.SetReadOnly(true)
	vbox.Append(statusOutput, false)

	return vbox
}

func makeRDPPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	enableBtn := ui.NewButton("Enable")
	vbox.Append(enableBtn, false)

	disableBtn := ui.NewButton("Disable")
	vbox.Append(disableBtn, false)

	vbox.Append(ui.NewLabel("Status:"), false)
	statusOutput := ui.NewMultilineEntry()
	statusOutput.SetReadOnly(true)
	vbox.Append(statusOutput, false)

	return vbox
}

func setupUI() {

	mainwin = ui.NewWindow("goP2PVPN", 640, 480, true)
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

	tab.Append("Shell", makeShellPage())
	tab.SetMargined(2, true)

	tab.Append("File transfer", makeTransferPage())
	tab.SetMargined(3, true)

	tab.Append("Proxy", makeProxyPage())
	tab.SetMargined(4, true)

	tab.Append("RDP", makeRDPPage())
	tab.SetMargined(5, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
