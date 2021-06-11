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
var yoff int32 = winHeight - height - 30 

// Here you can change the number of blocks
var block_count int32 = 10

// This value is the delay between the moves
var delay uint32 = 11

func push(s []int, v int) []int {
	return append(s, v)
}

func pop(s []int) ([]int, int) {
	l := len(s)
	if l == 0 {
		fmt.Println("stack empty")
		return []int{}, 0
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

func moveBlock(counter int32, towers [][]int) [][]int {
	legalMove := func(src, dst int, t [][]int) [][]int {
		var s int
		if len(t[src]) == 0 || len(t[dst]) == 0 {
			if len(t[src]) > len(t[dst]) {
				t[src], s = pop(t[src])
				if s == 0 {return t}
				t[dst] = push(t[dst], s)
			} else {
				t[dst], s = pop(t[dst])
				if s == 0 {return t}
				t[src] = push(t[src], s)
			}
			return t
		}
		if towers[src][len(towers[src])-1] < towers[dst][len(towers[dst])-1] {
			towers[src], s = pop(towers[src])
			towers[dst] = push(towers[dst], s)
		} else {
			towers[dst], s = pop(towers[dst])
			towers[src] = push(towers[src], s)
		}
		return t
	}
	if counter % 3 == 1 {
		towers = legalMove(0, 2, towers)
	}
	if counter % 3 == 2 {
		towers = legalMove(0, 1, towers)
	}
	if counter % 3 == 0 {
		towers = legalMove(1, 2, towers)
	}
	return towers
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
	tower := []int{}
	for i := int(block_count); i > 0; i-- {
		tower = push(tower, i)
	}

	towers = append(towers, tower)
	towers = append(towers, []int{})
	towers = append(towers, []int{})

	move_count := int32(math.Pow(2, float64(block_count)))
	move_counter := int32(1)
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

		if move_count >= move_counter {
			towers = moveBlock(move_counter, towers)
			move_counter++
		} 
		drawTowers(towers, renderer)

		renderer.Present()
		// sdl.Delay(16)
		sdl.Delay(delay)
	}

	return 0
}

func main() {
	os.Exit(run())
}
