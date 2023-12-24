package day22

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// return z axis of collision
func willCollide(top, bot brick) int {
	topMinX := min(top.v1.x, top.v2.x)
	topWidth := abs(top.v1.x-top.v2.x) + 1
	topMinY := min(top.v1.y, top.v2.y)
	topHeight := abs(top.v1.y-top.v2.y) + 1

	bMinX := min(bot.v1.x, bot.v2.x)
	bWidth := abs(bot.v1.x-bot.v2.x) + 1
	bMinY := min(bot.v1.y, bot.v2.y)
	bHeight := abs(bot.v1.y-bot.v2.y) + 1

	if topMinX < bMinX+bWidth &&
		topMinX+topWidth > bMinX &&
		topMinY < bMinY+bHeight &&
		topMinY+topHeight > bMinY {
		return max(bot.v1.z, bot.v2.z) + 1
	}

	return -1
}

type coord struct {
	x, y, z int
}

type brick struct {
	id         int
	v1, v2     coord
	supporting []int
}

func Part1(input string) {
	lines := strings.Split(input, "\n")

	// sort all bricks from bottom to top
	// move each brick as far down as it can go, until it obstructs any other bricks
	bricks := []brick{}
	for i, line := range lines {
		rawCoords := strings.Split(line, "~")

		b := brick{id: i, supporting: []int{}}
		for i, rc := range rawCoords {
			c := coord{}
			for j, _v := range strings.Split(rc, ",") {
				v, _ := strconv.Atoi(_v)
				if j == 0 {
					c.x = v
				} else if j == 1 {
					c.y = v
				} else {
					c.z = v
				}
			}
			if i == 0 {
				b.v1 = c
			} else {
				b.v2 = c
			}
		}

		bricks = append(bricks, b)
	}

	sort.Slice(bricks, func(i, j int) bool {
		iMinZ := min(bricks[i].v1.z, bricks[i].v2.z)
		jMinZ := min(bricks[j].v1.z, bricks[j].v2.z)

		return iMinZ > jMinZ
	})

	supportedBy := map[int][]int{}
	for i := len(bricks) - 1; i >= 0; i-- {
		b := bricks[i]

		collided := false
		maxZ := 1
		colliders := []struct{ idx, z int }{}
		for j, lowerBrick := range bricks[i+1:] {
			v := willCollide(b, lowerBrick)
			if v == -1 {
				continue
			}

			collided = true

			colliders = append(colliders, struct {
				idx int
				z   int
			}{i + j + 1, v})
			if v > maxZ {
				maxZ = v
			}
		}

		if collided {
			if bricks[i].v1.z < bricks[i].v2.z {
				d := bricks[i].v2.z - bricks[i].v1.z
				bricks[i].v1.z = maxZ
				bricks[i].v2.z = maxZ + d
			} else {
				d := bricks[i].v1.z - bricks[i].v2.z
				bricks[i].v2.z = maxZ
				bricks[i].v1.z = maxZ + d
			}

			for _, collider := range colliders {
				cb := bricks[collider.idx]
				if max(cb.v1.z, cb.v2.z)+1 == maxZ {
					// this collided brick is supporting the `brick`
					if _, exists := supportedBy[b.id]; !exists {
						supportedBy[b.id] = []int{}
					}
					supportedBy[b.id] = append(supportedBy[b.id], cb.id)
					bricks[collider.idx].supporting = append(bricks[collider.idx].supporting, b.id)

				}
			}
		}

		if !collided {
			v := 1
			if bricks[i].v1.z < bricks[i].v2.z {
				d := bricks[i].v2.z - bricks[i].v1.z
				bricks[i].v1.z = v
				bricks[i].v2.z = v + d
			} else {
				d := bricks[i].v1.z - bricks[i].v2.z
				bricks[i].v1.z = v + d
				bricks[i].v2.z = v
			}
		}
	}

	tot := 0
	for i := 0; i < len(bricks); i++ {
		b := bricks[i]

		if len(b.supporting) == 0 {
			tot++
			continue
		}

		allShared := true
		for _, supportedId := range b.supporting {
			if len(supportedBy[supportedId]) == 1 {
				allShared = false
				break
			}
		}

		if allShared {
			tot++
		}
	}

	fmt.Println("Part 1:", tot)
}

func Part2(input string) {
	lines := strings.Split(input, "\n")

	bricks := []brick{}
	for i, line := range lines {
		rawCoords := strings.Split(line, "~")

		b := brick{id: i, supporting: []int{}}
		for i, rc := range rawCoords {
			c := coord{}
			for j, _v := range strings.Split(rc, ",") {
				v, _ := strconv.Atoi(_v)
				if j == 0 {
					c.x = v
				} else if j == 1 {
					c.y = v
				} else {
					c.z = v
				}
			}
			if i == 0 {
				b.v1 = c
			} else {
				b.v2 = c
			}
		}

		bricks = append(bricks, b)
	}

	sort.Slice(bricks, func(i, j int) bool {
		iMinZ := min(bricks[i].v1.z, bricks[i].v2.z)
		jMinZ := min(bricks[j].v1.z, bricks[j].v2.z)

		return iMinZ > jMinZ
	})

	supportedBy := map[int][]int{}
	for i := len(bricks) - 1; i >= 0; i-- {
		b := bricks[i]

		collided := false
		maxZ := 1
		colliders := []struct{ idx, z int }{}
		for j, lowerBrick := range bricks[i+1:] {
			v := willCollide(b, lowerBrick)
			if v == -1 {
				continue
			}

			collided = true

			colliders = append(colliders, struct {
				idx int
				z   int
			}{i + j + 1, v})
			if v > maxZ {
				maxZ = v
			}
		}

		if collided {
			if bricks[i].v1.z < bricks[i].v2.z {
				d := bricks[i].v2.z - bricks[i].v1.z
				bricks[i].v1.z = maxZ
				bricks[i].v2.z = maxZ + d
			} else {
				d := bricks[i].v1.z - bricks[i].v2.z
				bricks[i].v2.z = maxZ
				bricks[i].v1.z = maxZ + d
			}

			for _, collider := range colliders {
				cb := bricks[collider.idx]
				if max(cb.v1.z, cb.v2.z)+1 == maxZ {
					// this collided brick is supporting the `brick`
					if _, exists := supportedBy[b.id]; !exists {
						supportedBy[b.id] = []int{}
					}
					supportedBy[b.id] = append(supportedBy[b.id], cb.id)
					bricks[collider.idx].supporting = append(bricks[collider.idx].supporting, b.id)

				}
			}
		}

		if !collided {
			v := 1
			if bricks[i].v1.z < bricks[i].v2.z {
				d := bricks[i].v2.z - bricks[i].v1.z
				bricks[i].v1.z = v
				bricks[i].v2.z = v + d
			} else {
				d := bricks[i].v1.z - bricks[i].v2.z
				bricks[i].v1.z = v + d
				bricks[i].v2.z = v
			}
		}
	}

	var getNDest func(id int, siblings []int, destroyed map[int]bool) int

	isDestroyable := func(id int, siblings []int, destroyed map[int]bool) bool {
		// have all of the siblings broken by previous id
		// check if all bricks supporting `supp` are included in siblings, if so, increment
		for _, supporter := range supportedBy[id] {
			if destroyed[supporter] {
				continue
			}

			found := false
			for _, sib := range siblings {
				if sib == supporter {
					found = true
					break
				}
			}

			if !found {
				return false
			}
		}

		return true
	}

	getNDest = func(id int, siblings []int, destroyed map[int]bool) int {
		n := 0
		b := brick{}
		for _, _b := range bricks {
			if _b.id != id {
				continue
			}
			b = _b
		}

		if len(b.supporting) == 0 {
			return 0
		}

		for _, supp := range b.supporting {
			if destroyed[supp] {
				continue
			}

			if isDestroyable(supp, siblings, destroyed) {
				destroyed[supp] = true
				destroyableSiblings := []int{}
				for _, sib := range b.supporting {
					if isDestroyable(sib, siblings, destroyed) {
						destroyableSiblings = append(destroyableSiblings, sib)
					}
				}

				v := getNDest(supp, destroyableSiblings, destroyed)
				n += v + 1
			}
		}

		return n
	}

	tot := 0
	for i := 0; i < len(bricks); i++ {
		v := getNDest(bricks[i].id, []int{bricks[i].id}, map[int]bool{})
		tot += v
	}

	fmt.Println("Part 2:", tot)
}
