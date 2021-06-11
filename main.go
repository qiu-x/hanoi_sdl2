// author: Jacky Boen

package main

import (
	"fmt"
	"os"
	"math"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Hanoi-SDL2"
var winWidth, winHeight int32 = 800, 600

var diff int32 = 50
var spacing int32 = 70
var height int32 = 50
var yoff int32 = winHeight - height - 50 
var block_count int32 = 4

func push(s []int, v int) []int {
	return append(s, v)
}

func pop(s []int) ([]int, int) {
	l := len(s)
	if l == 0 {
		panic("stack empty")
	}
	return  s[:l-1], s[l-1]
} 

func drawTowers(t [][]int, r *sdl.Renderer) {
	tower_count := int32(len(t))
	base_size := winWidth / tower_count - spacing
	
	for i, v := range t {
		i32 := int32(i)
		r.SetDrawColor(255, 255, 255, 255)
		block_start := spacing / 2 + (spacing + base_size) * i32
		rect := &sdl.Rect{block_start, yoff, base_size, height}
		r.DrawRect(rect)
		multiplier := base_size/block_count
		color_multiplier := 255/block_count
		for j, b := range v {
			j32 := int32(j)
			red := uint8(color_multiplier * int32(b))
			r.SetDrawColor(red, 255, 255, 255)
			reduction := ((block_count - int32(b)) * multiplier)/2
			rect := &sdl.Rect{block_start + reduction, yoff - height * (j32+1),
			base_size - reduction*2, height}
			r.FillRect(rect)
		}
	}
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	var towers [][]int
	tower1 := []int{4, 3, 2, 1}
	tower2 := []int{}
	tower3 := []int{}

	towers = append(towers, tower1)
	towers = append(towers, tower2)
	towers = append(towers, tower3)

	move_count := int32(math.Pow(2, float64(block_count)) - 1)
	move_counter := int32(0)
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		drawTowers(towers, renderer)
		if move_count >= move_counter {
			move_counter++
		}
		

		renderer.Present()
		sdl.Delay(16)
	}

	return 0
}

func main() {
	os.Exit(run())
}

// TODO:
// 1. Calculate the total number of moves required i.e. "pow(2, n) - 1" here n is number of disks.
// 2. If number of disks (i.e. n) is even then interchange destination pole and auxiliary pole.
// 3. for i = 1 to total number of moves:
// 4. if i%3 == 1: legal movement of top disk between source pole and destination pole
// 5. if i%3 == 2: legal movement top disk between source pole and auxiliary pole
// 6. if i%3 == 0: legal movement top disk between auxiliary pole and destination poles

