package vpn

import (
	"github.com/golang-ui/nuklear/nk"
	"log"
)

type ServerGUI struct {
	state State
}

func (sgui* ServerGUI) Init() {
	sgui.state = State{make([]byte, 32), 0}
}

func (sgui* ServerGUI) GfxMain(ctx *nk.Context) {
	// Layout
	bounds := nk.NkRect(400, 0, winWidth, winHeight)
	update := nk.NkBegin(ctx, "VPN Server", bounds,
		sgui.state.hidden | nk.WindowBorder|nk.WindowTitle|nk.WindowClosable|nk.WindowBackground)

	if update > 0 {
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "Server address:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditStringZeroTerminated(ctx, nk.EditBox, sgui.state.buffer, 32, nil)
		}
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "Connection token:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkLabel(ctx, "Server address:", nk.TextLeft)
		}
		nk.NkLayoutRowStatic(ctx, 30, 80, 2)
		{
			if nk.NkButtonLabel(ctx, "Gen new") > 0 {
				log.Println("[INFO] button pressed!")
			}
			if nk.NkButtonLabel(ctx, "Copy") > 0 {
				nk.NkWindowShow(ctx, "Please select VPN mode", nk.Shown)
			}
		}
		nk.NkLayoutRowDynamic(ctx, 20, 1)
		{
			nk.NkLabel(ctx, "Password:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(ctx, 30, 1)
		{
			nk.NkEditStringZeroTerminated(ctx, nk.EditBox, sgui.state.buffer, 32, nil)
		}
		nk.NkLayoutRowStatic(ctx, 30, 80, 2)
		{
			if nk.NkButtonLabel(ctx, "Connect") > 0 {
				log.Println("[INFO] button pressed!")
			}
			if nk.NkButtonLabel(ctx, "Back") > 0 {
				nk.NkWindowShow(ctx, "VPN Server", nk.Hidden)
				nk.NkWindowShow(ctx, "Please select VPN mode", nk.Shown)
			}
		}
	}
	//
	nk.NkEnd(ctx)
}