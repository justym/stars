package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Star struct {
	Image *ebiten.Image
	Color color.Color
	X int
	Y int
	Z int
}

func (s *Star) Init(x,y int){
	s.X = x
	s.Y = y
	s.Z = x
	s.Image,_ = ebiten.NewImage(screenWidth,screenHeight,ebiten.FilterDefault)
	s.Color = color.White
	s.Image.Set(s.X,s.Y,s.Color)
}

func (s *Star) Update(){
	/*err := s.Image.Clear()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	s.Image.Set(s.X-10,s.Y-10,s.Color)*/
}

//////////////////////ENTRY POINT FROM HERE///////////////////////////
//var offscreen *ebiten.Image
var stars = make([]Star,100)

func init() {
	//offscreen, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	for i := 0; i < len(stars); i++ {
		x := rand.Intn(screenWidth)
		y := rand.Intn(screenHeight)
		stars[i].Init(x, y)
	}
}

//var index = 0
var starFrame = 0
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	// Draw one to starFrames
	for index := 0; index <= starFrame; index++ {
		err := screen.DrawImage(stars[index].Image, nil)
		if err != nil {
			return err
		}
	}
	starFrame = (starFrame + 1) % len(stars)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Star Fields"); err != nil {
		log.Fatal(err)
	}
}
