package main

import "github.com/veandco/go-sdl2/sdl"

func initSDL() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit();
}

func initWindow() (window *sdl.Window) {
	window, err := sdl.CreateWindow("HELLO GO-SDL", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

  return window
}

func createSurface(window *sdl.Window) (surface *sdl.Surface) {
  surface, err := window.GetSurface()

	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)

  return surface;
}

func gameLoop() {
	running := true
	for running {

    for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			
      switch event.(type) {
			
        case *sdl.QuitEvent:
          println("Quit")
          running = false
          break
        
        default:
          break
      }

		}
	
  }
}

func main() {
  initSDL()
  window := initWindow()
  defer window.Destroy();

	createSurface(window);

	window.UpdateSurface()

  gameLoop()
}
