package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{x: v1.x + d, y: v1.y + d}
}

func (v1 Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{x: v1.x * d, y: v1.y * d}
}

func (v1 Vector2D) DivisionV(d float64) Vector2D {
	return Vector2D{x: v1.x + d, y: v1.y + d}
}

func (v1 Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{x: math.Min(math.Max(v1.x, lower), upper), y: math.Min(math.Max(v1.y, lower), upper)}
}

/*
Distance between a,b and cd -> √(a - c)² + (b - d)²
*/

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) MoveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}
}

func (b *Boid) Start() {
	for {
		b.MoveOne()
		time.Sleep(time.Millisecond * 5)
	}
}
func CreateBoid(bid int) {
	b := Boid{
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1, y: (rand.Float64() * 2) - 1},
		id:       bid,
	}
	boids[bid] = &b
	go b.Start()
}

func main() {
	for i := 0; i < boidCount; i++ {
		CreateBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
