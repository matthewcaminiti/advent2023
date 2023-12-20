package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func manhattanDist(a, b coord) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

var minHeat = 1000000000000000000

func minWarmthFrom(visited [][]bool, nums [][]int, from coord, at coord, currHeat int) {
	w := len(nums[0])
	h := len(nums)

	fmt.Printf("(%d, %d) -> (%d, %d): %d, %d\n", from.x, from.y, at.x, at.y, currHeat, minHeat)
	if at.x == 2 && at.y == 0 && from.x == 0 && from.y == 0 {
		fmt.Println(visited)
	}

	canRight := at.x == from.x
	canLeft := canRight
	canDown := at.y == from.y
	canUp := canDown

	if at.x == 0 && at.y == 0 && from.x == 0 && from.y == 0 {
		canRight = true
		canDown = true
	}

	v := nums[at.y][at.x]
	if at.x == w-1 && at.y == h-1 {
		minHeat = min(minHeat, v+currHeat)
		return
	}

	if v+currHeat+manhattanDist(at, coord{w - 1, h - 1}) >= minHeat {
		return
	}

	if visited[at.y][at.x] {
		return
	}

	// fmt.Printf("(%d, %d) -> (%d, %d)\n", from.x, from.y, at.x, at.y)

	jah := make([][]bool, len(visited))
	for y := range visited {
		jah[y] = make([]bool, len(visited[y]))
		for x := range visited[y] {
			jah[y][x] = visited[y][x]
		}
	}

	jah[at.y][at.x] = true

	dx := at.x - from.x
	if dx != 0 {
		for i := 0; i < abs(dx); i++ {
			if dx < 0 {
				jah[at.y][at.x-dx-i] = true
			} else {
				jah[at.y][at.x-dx+i] = true
			}
		}
	}

	dy := at.y - from.y
	if dy != 0 {
		for i := 0; i < abs(dy); i++ {
			if dy < 0 {
				jah[at.y-dy-i][at.x] = true
			} else {
				jah[at.y-dy+i][at.x] = true
			}
		}
	}

	// fmt.Println(from, at, minHeat, v+currHeat, manhattanDist(at, coord{w-1,h-1}))
	if canRight {
		if at.x+1 < w {
			minWarmthFrom(jah, nums, at, coord{at.x + 1, at.y}, currHeat+v)
		}
		if at.x+2 < w {
			v1 := nums[at.y][at.x+1]
			minWarmthFrom(jah, nums, at, coord{at.x + 2, at.y}, currHeat+v+v1)
		}
		if at.x+3 < w {
			v1 := nums[at.y][at.x+1]
			v2 := nums[at.y][at.x+2]
			minWarmthFrom(jah, nums, at, coord{at.x + 3, at.y}, currHeat+v+v1+v2)
		}
	}

	if canDown {
		if at.y+1 < h {
			minWarmthFrom(jah, nums, at, coord{at.x, at.y + 1}, currHeat+v)
		}
		if at.y+2 < h {
			v1 := nums[at.y+1][at.x]
			minWarmthFrom(jah, nums, at, coord{at.x, at.y + 2}, currHeat+v+v1)
		}
		if at.y+3 < h {
			v1 := nums[at.y+1][at.x]
			v2 := nums[at.y+2][at.x]
			minWarmthFrom(jah, nums, at, coord{at.x, at.y + 3}, currHeat+v+v1+v2)
		}
	}

	if canLeft {
		if at.x-1 >= 0 {
			minWarmthFrom(jah, nums, at, coord{at.x - 1, at.y}, currHeat+v)
		}
		if at.x-2 >= 0 {
			v1 := nums[at.y][at.x-1]
			minWarmthFrom(jah, nums, at, coord{at.x - 2, at.y}, currHeat+v+v1)
		}
		if at.x-3 >= 0 {
			v1 := nums[at.y][at.x-1]
			v2 := nums[at.y][at.x-2]
			minWarmthFrom(jah, nums, at, coord{at.x - 3, at.y}, currHeat+v+v1+v2)
		}
	}

	if canUp {
		if at.y-1 >= 0 {
			minWarmthFrom(jah, nums, at, coord{at.x, at.y - 1}, currHeat+v)
		}
		if at.y-2 >= 0 {
			v1 := nums[at.y-1][at.x]
			minWarmthFrom(jah, nums, at, coord{at.x, at.y - 2}, currHeat+v+v1)
		}
		if at.y-3 >= 0 {
			v1 := nums[at.y-1][at.x]
			v2 := nums[at.y-2][at.x]
			minWarmthFrom(jah, nums, at, coord{at.x, at.y - 3}, currHeat+v+v1+v2)
		}
	}
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	visited := make([][]bool, len(lines))
	for i, line := range lines {
		visited[i] = make([]bool, len(line))
	}

	nums := make([][]int, len(lines))
	for y := range lines {
		nums[y] = make([]int, len(lines[y]))
		for x := range lines[y] {
			nums[y][x], _ = strconv.Atoi(string(lines[y][x]))
		}
	}

	minWarmthFrom(visited, nums, coord{0, 0}, coord{0, 0}, 0)

	fmt.Println("Part 1:", minHeat)
}

func minWarmthFromMem(
	visited [][]bool,
	nums [][]int,
	hist [][]int,
	from coord,
	at coord,
) int {
	w := len(nums[0])
	h := len(nums)

	// fmt.Printf("(%d, %d) -> (%d, %d): %d, %d\n", from.x, from.y, at.x, at.y, currHeat, minHeat)
	canRight := at.x == from.x
	canLeft := canRight
	canDown := at.y == from.y
	canUp := canDown

	if at.x == from.x && at.y == from.y {
		canRight = true
		canDown = true
		canUp = true
		canLeft = true
	}

	v := nums[at.y][at.x]
	if at.x == w-1 && at.y == h-1 {
		hist[at.y][at.x] = v
		return v
	}

	if visited[at.y][at.x] {
		return 1000000000000000000
	}

	// fmt.Printf("(%d, %d) -> (%d, %d)\n", from.x, from.y, at.x, at.y)

	jah := make([][]bool, len(visited))
	for y := range visited {
		jah[y] = make([]bool, len(visited[y]))
		for x := range visited[y] {
			jah[y][x] = visited[y][x]
		}
	}

	jah[at.y][at.x] = true

	dx := at.x - from.x
	if dx != 0 {
		for i := 0; i < abs(dx); i++ {
			if dx < 0 {
				jah[at.y][at.x-dx-i] = true
			} else {
				jah[at.y][at.x-dx+i] = true
			}
		}
	}

	dy := at.y - from.y
	if dy != 0 {
		for i := 0; i < abs(dy); i++ {
			if dy < 0 {
				jah[at.y-dy-i][at.x] = true
			} else {
				jah[at.y-dy+i][at.x] = true
			}
		}
	}

	// fmt.Println(from, at, minHeat, v+currHeat, manhattanDist(at, coord{w-1,h-1}))
	res := 1000000000000000000

	if canRight {
		if at.x+1 < w {
			if hist[at.y][at.x+1] != -1 {
				res = min(res, hist[at.y][at.x+1]+v)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x + 1, at.y})
				res = min(res, x+v)
			}
		}
		if at.x+2 < w {
			v1 := nums[at.y][at.x+1]
			if hist[at.y][at.x+2] != -1 {
				res = min(res, hist[at.y][at.x+2]+v+v1)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x + 2, at.y})
				res = min(res, x+v+v1)
			}
		}
		if at.x+3 < w {
			v1 := nums[at.y][at.x+1]
			v2 := nums[at.y][at.x+2]
			if hist[at.y][at.x+3] != -1 {
				res = min(res, hist[at.y][at.x+3]+v+v1+v2)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x + 3, at.y})
				res = min(res, x+v+v1+v2)
			}
		}
	}

	if canDown {
		if at.y+1 < h {
			if hist[at.y+1][at.x] != -1 {
				res = min(res, hist[at.y+1][at.x]+v)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x, at.y + 1})
				res = min(res, x+v)
			}
		}
		if at.y+2 < h {
			v1 := nums[at.y+1][at.x]
			if hist[at.y+2][at.x] != -1 {
				res = min(res, hist[at.y+2][at.x]+v+v1)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x, at.y + 2})
				res = min(res, x+v+v1)
			}
		}
		if at.y+3 < h {
			v1 := nums[at.y+1][at.x]
			v2 := nums[at.y+2][at.x]
			if hist[at.y+3][at.x] != -1 {
				res = min(res, hist[at.y+3][at.x]+v+v1+v2)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x, at.y + 3})
				res = min(res, x+v+v1+v2)
			}
		}
	}

	if canLeft {
		if at.x-1 >= 0 {
			if hist[at.y][at.x-1] != -1 {
				res = min(res, hist[at.y][at.x-1]+v)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x - 1, at.y})
				res = min(res, x+v)
			}
		}
		if at.x-2 >= 0 {
			v1 := nums[at.y][at.x-1]
			if hist[at.y][at.x-2] != -1 {
				res = min(res, hist[at.y][at.x-2]+v+v1)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x - 2, at.y})
				res = min(res, x+v+v1)
			}
		}
		if at.x-3 >= 0 {
			v1 := nums[at.y][at.x-1]
			v2 := nums[at.y][at.x-2]
			if hist[at.y][at.x-3] != -1 {
				res = min(res, hist[at.y][at.x-3]+v+v1+v2)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x - 3, at.y})
				res = min(res, x+v+v1+v2)
			}
		}
	}

	if canUp {
		if at.y-1 >= 0 {
			if hist[at.y-1][at.x] != -1 {
				res = min(res, hist[at.y-1][at.x]+v)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x, at.y - 1})
				res = min(res, x+v)
			}
		}
		if at.y-2 >= 0 {
			v1 := nums[at.y-1][at.x]
			if hist[at.y-2][at.x] != -1 {
				res = min(res, hist[at.y-2][at.x]+v+v1)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x, at.y - 2})
				res = min(res, x+v+v1)
			}
		}
		if at.y-3 >= 0 {
			v1 := nums[at.y-1][at.x]
			v2 := nums[at.y-2][at.x]
			if hist[at.y-3][at.x] != -1 {
				res = min(res, hist[at.y-3][at.x]+v+v1+v2)
			} else {
				x := minWarmthFromMem(jah, nums, hist, at, coord{at.x, at.y - 3})
				res = min(res, x+v+v1+v2)
			}
		}
	}

	hist[at.y][at.x] = res
	return res
}

func ReversePart1(input string) {
	lines := strings.Split(input, "\n")

	// start from bottom right corner 2-by-2
	// get every tile's least heat loss
	// memoize each tile's
	nums := make([][]int, len(lines))
	hist := make([][]int, len(lines))
	visited := make([][]bool, len(lines))
	for y := range lines {
		nums[y] = make([]int, len(lines[y]))
		hist[y] = make([]int, len(lines[y]))
		visited[y] = make([]bool, len(lines[y]))
		for x := range lines[y] {
			nums[y][x], _ = strconv.Atoi(string(lines[y][x]))
			hist[y][x] = -1
		}
	}

	h := len(lines)
	w := len(lines[0])

	for dim := 5; dim <= w; dim++ {
		// take subsets of space
		_visited := make([][]bool, dim)
		_nums := make([][]int, dim)
		_hist := make([][]int, dim)
		for y := 0; y < dim; y++ {
			_visited[y] = make([]bool, dim)
			_nums[y] = make([]int, dim)
			_hist[y] = make([]int, dim)
			for x := 0; x < dim; x++ {
				_visited[y][x] = visited[h-dim+y][w-dim+x]
				_nums[y][x] = nums[h-dim+y][w-dim+x]
				// if dim > 3 && (x == 1 || y == 1) {
				//     _hist[y][x] = -1
				// } else {
				_hist[y][x] = hist[h-dim+y][w-dim+x]
				// }
			}

			// fmt.Printf("%v | %v | %v\n", _visited[y], _nums[y], _hist[y])
		}

		for _, line := range _hist {
			fmt.Println(line)
		}
		fmt.Println("vvvvv")

		for y := 0; y < dim; y++ {
			for x := 0; x < dim; x++ {
				if (x == 2 || y == 2) && x >= 1 && y >= 1 {
					minWarmthFromMem(_visited, _nums, _hist, coord{x, y}, coord{x, y})
				}
			}
		}

		for _, line := range _hist {
			fmt.Println(line)
		}
		fmt.Println("=======")
		// debrief hist
		for y := 0; y < dim; y++ {
			for x := 0; x < dim; x++ {
				hist[h-dim+y][w-dim+x] = _hist[y][x]
			}
		}
		// break
	}

	// for _, line := range hist {
	//     fmt.Println(line)
	// }

	fmt.Println("Reverse Memoize: ", hist[0][0])
}

var inf = math.MaxInt64

type edge struct {
	toId int
	dist int
}

type node struct {
	id    int
	edges []edge
}

func nodeId(w, x, y int) int {
	return w*y + x
}

func coordFromId(w, id int) coord {
	y := int(id / w)
	x := id % w

	return coord{x, y}
}

func dijk(g []node, source node, canMove func(from, at, to int) bool) ([]int, []int) {
	dist := make([]int, len(g))
	prev := make([]int, len(g))
	q := make([]node, len(g))

	for i, node := range g {
		dist[node.id] = inf
		prev[node.id] = -1
		q[i] = node
	}
	dist[source.id] = 0

	for len(q) > 0 {
		minIdx, minVal := -1, inf
		for i, node := range q {
			if dist[node.id] < minVal {
				minIdx = i
				minVal = dist[node.id]
			}
		}

		u := q[minIdx]

		q[minIdx] = q[len(q)-1]
		q = q[:len(q)-1]

		for _, v := range u.edges {
			for _, n := range q {
				if v.toId == n.id {
					canDoMove := true
					if prev[u.id] != -1 {
						// determine if, given previous node, can move to this node
						canDoMove = canMove(prev[u.id], u.id, v.toId)
					}

					if !canDoMove {
						continue
					}

					alt := dist[u.id] + v.dist
					if alt < dist[v.toId] {
						dist[v.toId] = alt
						prev[v.toId] = u.id
					}
				}
			}
		}
	}

	return dist, prev
}

func Part1Dijk(input string) {
	lines := strings.Split(input, "\n")

	h := len(lines)
	w := len(lines[0])

	nums := make([][]int, len(lines))
	for y := range lines {
		nums[y] = make([]int, len(lines[y]))
		for x := range lines[y] {
			nums[y][x], _ = strconv.Atoi(string(lines[y][x]))
		}
	}

	g := []node{}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			id := nodeId(w, x, y)

			edges := []edge{}
			dirs := []coord{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
			for _, dir := range dirs {
				if x == w-1 && y == h-1 {
					break
				}
				for i := 1; i <= 3; i++ {
					xi := x + dir.x*i
					yi := y + dir.y*i
					if xi < w && xi >= 0 && yi < h && yi >= 0 {
						dist := 0
						for j := 1; j <= i; j++ {
							dist += nums[y+dir.y*j][x+dir.x*j]
						}
						edges = append(edges, edge{toId: nodeId(w, xi, yi), dist: dist})
					}
				}
			}

			n := node{id, edges}
			// fmt.Printf("(%d, %d): %+v\n", x, y, n)
			g = append(g, n)
		}
	}

	source := g[0]

	canMove := func(from, at, to int) bool {
		fromCoord := coordFromId(w, from)
		atCoord := coordFromId(w, at)
		toCoord := coordFromId(w, to)

		dx := atCoord.x - fromCoord.x
		ndx := toCoord.x - atCoord.x
		if dx != 0 && ndx != 0 {
			if (dx < 0 && ndx > 0) || (dx > 0 && ndx < 0) {
				return false
			}

			return abs(dx)+abs(ndx) <= 3
		}

		dy := atCoord.y - fromCoord.y
		ndy := toCoord.y - atCoord.y
		if dy != 0 && ndy != 0 {
			if (dy < 0 && ndy > 0) || (dy > 0 && ndy < 0) {
				return false
			}

			return abs(dy)+abs(ndy) <= 3
		}

		return true
	}

	dist, _ := dijk(g, source, canMove)
	fmt.Println(dist[nodeId(w, w-1, h-1)])
	for y := 0; y < h; y++ {
		for _, n := range dist[y*h : (y+1)*h] {
			fmt.Printf("%3d ", n)
		}
		fmt.Println()
	}
}
