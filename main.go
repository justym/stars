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

//Type Star has own point values X,Y,Z,PZ (Previous Z)
type Star struct {
	X int
	Y int
	Z int
	PZ int
}

//Set Method sets own values
func (s *Star) Set(){
	s.X = rand.Intn(screenWidth) - rand.Intn(screenWidth)
	s.Y = rand.Intn(screenHeight) - rand.Intn(screenHeight)
	s.Z = rand.Intn(screenWidth)
	s.PZ = s.Z
}

//Update Method updates own values
func (s *Star) Update(){
	s.Z = s.Z - speed
	if s.Z < 1 {
		s.Z = screenWidth
		s.X = rand.Intn(screenWidth) - rand.Intn(screenWidth)
		s.Y = rand.Intn(screenHeight) - rand.Intn(screenHeight)
		s.PZ = s.Z
	}
}

//Draw Method draw own star in screen
func (s *Star) Draw(screen *ebiten.Image) error {
	var path vector.Path
	var drawOption = &vector.DrawPathOptions{
		LineWidth:   5,
		StrokeColor: color.White,
	}

	sx := float32(s.X / s.Z) + float32(s.X)
	sy := float32(s.Y / s.Z) + float32(s.Y)

	px := float32(s.X / s.PZ) + float32(s.X)
	py := float32(s.Y / s.PZ) + float32(s.Y)

	s.PZ = s.Z

	path.MoveTo(sx,sy)
	path.LineTo(px,py)

	path.Draw(screen,drawOption)

	return nil
}

const (
	screenWidth = 600
	screenHeight = 400
	scale = 1
	starCount = 500
	speed = 1
)

var stars = make([]Star,starCount)

func init() {
	for i := 0; i < starCount; i++ {
		stars[i].Set()
	}
}

func screenUpdate(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped(){
		return nil
	}

	if err := screen.Fill(color.Black); err != nil{
		return err
	}

	for i := 0; i < starCount; i++{
		stars[i].Update()
		if err := stars[i].Draw(screen); err != nil{
			return err
		}
	}

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



