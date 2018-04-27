package vpn

import (
	"github.com/golang-ui/nuklear/nk"
	"log"
)

type ServerGUI struct {
	state State
}

func (cgui* ServerGUI) Init() {
	cgui.state = State{make([]byte, 32), 0}
}

func (cgui* ServerGUI) Show(state bool) {
	if state == false {
		cgui.state.hidden = nk.WindowHidden
	} else {
		cgui.state.hidden = 0
	}

}

func (cgui* ServerGUI) GfxMain(ctx *nk.Context) {
	// Layout
	bounds := nk.NkRect(0, 0, winWidth, winHeight)
	update := nk.NkBegin(ctx, "VPN Server", bounds,
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
		nk.NkLayoutRowStatic(ctx, 30, 80, 1)
		{
			if nk.NkButtonLabel(ctx, "Connect") > 0 {
				log.Println("[INFO] button pressed!")
			}
		}
	}
	//
	nk.NkEnd(ctx)
}