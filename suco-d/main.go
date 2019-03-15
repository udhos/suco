package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

type game struct {
	*application.Application
	mydata int
}

func main() {

	g := game{}

	app, _ := application.Create(application.Options{
		Title:  "Sun Colony Desktop",
		Width:  800,
		Height: 600,
	})

	g.Application = app

	// Create a blue torus and add it to the scene
	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewPhong(math32.NewColor("DarkBlue"))
	torusMesh := graphic.NewMesh(geom, mat)
	g.Application.Scene().Add(torusMesh)

	// Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	g.Application.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	g.Application.Scene().Add(pointLight)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(0.5)
	g.Application.Scene().Add(axis)

	g.Application.CameraPersp().SetPosition(0, 0, 3)
	g.Application.Run()
}
