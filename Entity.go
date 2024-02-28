package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/exp/maps"
)

type Entity struct {
	entityType int32
	sprite     *Sprite
	text       *Text
	position   *sdl.Rect
	speed      int32 //per tickrate
	physics    int
}

func (entity *Entity) renderSprite(renderer *sdl.Renderer) {
	renderer.Copy(entity.sprite.texture, entity.sprite.getFrame(), entity.position)
}

func (entity *Entity) renderText(renderer *sdl.Renderer) {
	entity.text.renderText(renderer, entity.position)
}

func (entity *Entity) render(renderer *sdl.Renderer) {
	switch entity.entityType {
	case 0:
		entity.renderSprite(renderer)
	case 1:
		entity.renderText(renderer)
	}

}

func renderEntities(entitiesMap map[string]*Entity, renderer *sdl.Renderer) {
	entities := maps.Values(entitiesMap)
	for i := 0; i < len(entities); i++ {
		entity := entities[i]

		entity.render(renderer)
	}
}

func (entity *Entity) moveLeft(elapsed time.Duration) {
	entity.position.X -= entity.speed
}

func (entity *Entity) moveRight(elapsed time.Duration) {
	entity.position.X += entity.speed
}

func (entity *Entity) moveUp(elapsed time.Duration) {
	entity.position.Y -= entity.speed
}

func (entity *Entity) moveDown(elapsed time.Duration) {
	entity.position.Y += entity.speed
}
