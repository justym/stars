package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/vector"
	"image/color"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth = 600
	screenHeight = 400
	scale = 1
	amount = 500
	speed = 0.2
	coefficient = 1
)

// Type Star has own point values X,Y,Z,PZ (Previous Z)
type Star struct {
	X float32
	Y float32
	Z float32
	PZ float32
}

// Set Method sets own values
func (s *Star) Set(){
	s.X = randRange(-1 * screenWidth,screenWidth)
	s.Y = randRange(-1 * screenHeight,screenHeight)
	s.Z = randRange(0,screenWidth)
	s.PZ = s.Z
}

// Update Method updates own values
func (s *Star) Update(){
	s.Z = s.Z - speed
	if s.Z < 1 {
		s.X = randRange(-1 * screenWidth,screenWidth)
		s.Y = randRange(-1 * screenHeight, screenHeight)
		s.Z = randRange(0,screenWidth)
		s.PZ = s.Z
	}
}

// Draw Method draw own star in screen
func (s *Star) Draw(screen *ebiten.Image) error {
	var path vector.Path
	var op = &vector.DrawPathOptions{
		LineWidth:   2,
		StrokeColor: color.White,
	}

	var sx,sy,px,py float32
	sx = calculate(s.X,s.Z)
	sy = calculate(s.Y,s.Z)
	px = calculate(s.X,s.PZ)
	py = calculate(s.Y,s.PZ)
	s.PZ = s.Z

	path.MoveTo(sx,sy)
	path.LineTo(px,py)
	path.Draw(screen,op)

	return nil
}
// Return rand in range
func randRange(min,max int) float32 {
	return rand.Float32() * float32(max - min) + float32(min)
}

// Calculate sx,sy,px,py
func calculate(value, z float32)float32{
	if value < 0 {
		return (coefficient * (value / z)) - value
	}
	return (coefficient * (value / z)) + value
}

var stars = make([]Star,amount)

// Init stars
func init() {
	for i := 0; i < amount; i++ {
		stars[i].Set()
	}
}

//Update screen image
func screenUpdate(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped(){
		return nil
	}

	if err := screen.Fill(color.Black); err != nil{
		return err
	}

	for i := 0; i < amount; i++{
		stars[i].Update()
		if err := stars[i].Draw(screen); err != nil{
			return err
		}
	}

	//Debug Print fps, tps
	fps,tps := ebiten.CurrentFPS(),ebiten.CurrentTPS()
	if err := ebitenutil.DebugPrint(screen,fmt.Sprintf("FPS: %0.2f \nTPS: %0.2f \n",fps,tps)); err != nil{
		return err
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.Run(screenUpdate,screenWidth,screenHeight,scale,"Stars"); err != nil {
		log.Fatal(err)
	}
}



