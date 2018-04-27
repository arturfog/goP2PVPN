package main

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"
	"log"
	"time"
	"runtime"

	"./modules/vpn"
)

func init() {
	runtime.LockOSThread()
}

const (
	winWidth  = 800
	winHeight = 480

	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
)

func startGUI() {
	clientGUI := vpn.ClientGUI{}
	clientGUI.Init()
	serverGUI := vpn.ServerGUI{}
	serverGUI.Init()

	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "goP2PVPN", nil, nil)

	if err != nil {
		closer.Fatalln(err)
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	log.Printf("glfw: created window %dx%d", width, height)

	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)
	//
	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	nk.NkFontStashEnd()
	//
	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	fpsTicker := time.NewTicker(time.Second / 30)

	for {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			gfxMain(win, ctx, &clientGUI, &serverGUI)
		}
	}
}

func gfxMain(win *glfw.Window, ctx *nk.Context, cgui *vpn.ClientGUI, sgui *vpn.ServerGUI) {
	nk.NkPlatformNewFrame()

	// Layout
	bounds := nk.NkRect(240, 140, 320, 200)
	update := nk.NkBegin(ctx, "Please select VPN mode", bounds,
		nk.WindowBorder|nk.WindowTitle|nk.WindowMinimizable)

	if update > 0 {
		nk.NkLayoutRowStatic(ctx, 145, 145, 2)
		{
			if nk.NkButtonLabel(ctx, "VPN Client") > 0 {
				log.Println("[INFO] starting client ...")
				nk.NkWindowShow(ctx, "VPN Client", nk.Shown)
			}
			if nk.NkButtonLabel(ctx, "VPN Server") > 0 {
				log.Println("[INFO] starting server ...")
				nk.NkWindowShow(ctx, "VPN Client", nk.Hidden)
			}
		}
	}

	nk.NkEnd(ctx)

	cgui.GfxMain(ctx)
	sgui.GfxMain(ctx)

	// Render
	bg := make([]float32, 4)
	width, height := win.GetSize()
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
	nk.NkPlatformRender(nk.AntiAliasingOn, maxVertexBuffer, maxElementBuffer)
	win.SwapBuffers()
}

func main() {
	startGUI()
	//var wg sync.WaitGroup
	//
	//vps := vpn.NewVPNServer(&wg)
	//vps.Connect("127.0.0.1:8081", vps.GenKey())
	//
	//vpc := vpn.NewVPNClient(&wg)
	//vpc.Connect("127.0.0.1:8081", vps.GetKey())
	//wg.Wait()
	//
	//vps.Disconnect()
	//vpc.Disconnect()
}
