package main

import (
	"fmt"
	"gorp/gfx"
	"gorp/gobj"
	"gorp/utils"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func handleQuit() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch event.(type) {

		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

func handleFpsCounter(fpsCounter *Entity, start *time.Time, cycles *int) {
	t := time.Now()
	elapsed := t.Sub(*start)
	*cycles++
	if elapsed.Seconds() < 1 {
		return
	}

	fpsString := fmt.Sprintf("%d fps", int(float64(*cycles)/elapsed.Seconds()))
	fpsCounter.text.SetText(fpsString)
	*start = time.Now()
	*cycles = 0
}

func gameLoop(gameRenderer *sdl.Renderer) {
	start := time.Now()
	cycles := 0

	gObjManager := gobj.CreateGameObjectManager()
	gObjManager.FromJSON("./../assets/gobj.json")

	tManager := gfx.CreateTextureManager(gameRenderer)
	tManager.FromJSON("./../assets/textures.json")

	tileMap := generateTileMap(tManager)
	entities := loadEntities(tManager, gObjManager)
	fpsCounter := createFPSCounter()
	entities["fpsCounter"] = fpsCounter
	iHandlerAnimation := createInputHandler()
	iHandlerMovement := createInputHandler()
	mHandler := createMouseHandler()

	loadParticle(entities, gObjManager)

	camera := utils.CreateCamera()
	aRenderer := createARenderer(gameRenderer, camera)

	for {
		aRenderer.clearRenderer()
		aRenderer.handleRendering(tileMap)

		aRenderer.handleRendering(entities)

		aRenderer.present()
		iHandlerAnimation.animationHandler(entities)
		iHandlerMovement.handleMovement(gObjManager)
		mHandler.handleCameraMove(camera)

		handleFpsCounter(fpsCounter, &start, &cycles)

		handleQuit()
	}
}
