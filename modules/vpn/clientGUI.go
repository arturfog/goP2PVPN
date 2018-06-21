package vpn

import (
	"github.com/golang-ui/nuklear/nk"
	"log"
)

type State struct {
	buffer []byte
	hidden nk.Flags
}

const (
	winWidth  = 400
	winHeight = 300
)

type ClientGUI struct {
	state State
}

func (cgui* ClientGUI) Init() {
	cgui.state = State{make([]byte, 32), 0}
}

func (cgui* ClientGUI) GfxMain(ctx *nk.Context) {
	// Layout
	bounds := nk.NkRect(0, 0, winWidth, winHeight)
	update := nk.NkBegin(ctx, "VPN Client", bounds,
		cgui.state.hidden | nk.WindowBorder|nk.WindowTitle|nk.WindowClosable|nk.WindowBackground)

	if update > 0 {
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "Server address:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditStringZeroTerminated(ctx, nk.EditBox, cgui.state.buffer, 32, nil)
		}
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "Token:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditStringZeroTerminated(ctx, nk.EditBox, cgui.state.buffer, 32, nil)
		}
		nk.NkLayoutRowStatic(ctx, 30, 80, 2)
		{
			if nk.NkButtonLabel(ctx, "Paste from clipboard") > 0 {
				log.Println("[INFO] button pressed!")
			}
		}
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "Password:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditStringZeroTerminated(ctx, nk.EditBox, cgui.state.buffer, 32, nil)
		}
		nk.NkLayoutRowStatic(ctx, 30, 80, 2)
		{
			if nk.NkButtonLabel(ctx, "Connect") > 0 {
				log.Println("[INFO] button pressed!")
			}
			if nk.NkButtonLabel(ctx, "Back") > 0 {
				nk.NkWindowShow(ctx, "VPN Client", nk.Hidden)
				nk.NkWindowShow(ctx, "Please select VPN mode", nk.Shown)
			}
		}
	}
	//
	nk.NkEnd(ctx)
}