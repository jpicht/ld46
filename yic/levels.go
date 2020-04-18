package yic

import (
	"strings"
)

type levels struct {
	byKey  map[string]level
	chosen string
}

type level struct {
	width  int
	height int

	headX int
	headY int

	fields [][]field

	chains []*chain
}

type field struct {
	typ int
}

const (
	fieldWay = iota
	fieldObstacle
	fieldBuildSpot
	fieldHead
)

func createLevels() *levels {
	return &levels{
		byKey: map[string]level{
			"1": parseLevel(levelOne),
		},
	}
}

func parseLevel(input string) level {
	fields := make([][]field, 0)
	lines := strings.Split(input, "\n")
	width := 0
	foundHead := false
	headX, headY := 0, 0
	waypoints := map[vector2D]*waypoint{}
	for i := range lines {

		// Skip empty lines
		if len(lines[i]) == 0 {
			continue
		}

		if len(lines[i]) > width {
			width = len(lines[i])
		}

		y := len(fields)
		currentRow := make([]field, len(lines[i]))
		for x := range lines[i] {
			position := vector2D{
				X: x,
				Y: y,
			}
			currentRow[x].typ = fieldBuildSpot
			switch lines[i][x] {
			case 'X':
				currentRow[x].typ = fieldObstacle
			case 'O':
				if !foundHead {
					headX = x
					headY = y
					foundHead = true
				}
			case 'v', '^', '<', '>':
				currentRow[x].typ = fieldWay
				waypoints[position] = &waypoint{
					position:  position,
					direction: waypointDirections[lines[i][x]],
					previous:  make([]vector2D, 0),
				}
			}
		}
		fields = append(fields, currentRow)
	}

	height := len(fields)

	// Head is too far to the left or the bottom, input is invalid.
	if headX+1 >= width || headY+1 >= height {
		panic("head out of bounds")
	}

	// Head is 2x2, all fields are head fields and waypoints.
	for dx := 0; dx <= 1; dx++ {
		for dy := 0; dy <= 1; dy++ {
			headpos := vector2D{
				X: headX + dx,
				Y: headY + dy,
			}
			waypoints[headpos] = &waypoint{
				position: headpos,
				isHead:   true,
			}
			fields[headpos.Y][headpos.X].typ = fieldHead
		}
	}

	// Some lines may be shorter, fill them with build spots.
	for i := range fields {
		for j := 0; j < width-len(fields[i]); j++ {
			fields[i] = append(fields[i], field{typ: fieldBuildSpot})
		}
	}

	// Waypoint post-processing. Connect waypoints.
	for v := range waypoints {
		// Head waypoints are always final and are handled by waypoints pointing to them.
		if waypoints[v].isHead {
			continue
		}
		currentPos := v
		for {
			currentPos = addVector2Ds(currentPos, waypoints[v].direction)

			// We left the level, input is invalid.
			if currentPos.X < 0 || currentPos.Y < 0 || currentPos.X >= width || currentPos.Y >= height {
				panic(waypoints)
			}

			// This waypoint pointed to another waypoint. Connect both.
			if w, ok := waypoints[currentPos]; ok {
				w.previous = append(w.previous, v)
				waypoints[v].next = w
				break
			}

			// Every field we visit (which is not already a way) will be converted into a way field.
			fields[currentPos.Y][currentPos.X].typ = fieldWay
		}
	}

	// Remove head waypoints which are not target of another waypoint.
	deletedHeadWayPoints := 0
	for dx := 0; dx <= 1; dx++ {
		for dy := 0; dy <= 1; dy++ {
			headpos := vector2D{
				X: headX + dx,
				Y: headY + dy,
			}
			if len(waypoints[headpos].previous) == 0 {
				delete(waypoints, headpos)
				deletedHeadWayPoints++
			}
		}
	}

	// If no head waypoint is target of another waypoint, the level is unlosable.
	if deletedHeadWayPoints == 4 {
		panic("head cannot be reached, input invalid")
	}

	startingWaypoints := map[vector2D]*waypoint{}

	for v := range waypoints {
		if len(waypoints[v].previous) == 0 {
			startingWaypoints[v] = waypoints[v]
		}
	}

	// Because we have at least one final (head) waypoint and there must be an
	// equal amount of next and previous waypoints, at least one waypoint is a
	// starting waypoint.

	chains := make([]*chain, 0, len(startingWaypoints))
	for v := range startingWaypoints {
		currentWP := startingWaypoints[v]
		ch := &chain{
			segments: make([]*segment, 0),
		}
		for {
			nextWP := currentWP.next
			if nextWP == nil {
				break
			}
			ch.segments = append(
				ch.segments,
				&segment{
					start: currentWP,
					end:   nextWP,
				},
			)
			currentWP = nextWP
		}
		chains = append(chains, ch)
	}

	return level{
		chains: chains,
		width:  width,
		height: height,
		headX:  headX,
		headY:  headY,
		fields: fields,
	}
}

type waypoint struct {
	position  vector2D
	direction vector2D
	next      *waypoint
	previous  []vector2D

	// isHead means this is a final waypoint. It may not be part of a chain. In this case it will be removed in the end.
	isHead bool
}

var waypointDirections = map[byte]vector2D{
	'^': directionUp,
	'v': directionDown,
	'<': directionLeft,
	'>': directionRight,
}

type chain struct {
	segments []*segment
}

func (ch chain) length() int {
	l := 0
	for i := range ch.segments {
		l += ch.segments[i].length()
	}
	return l
}

type segment struct {
	start *waypoint
	end   *waypoint
}

func (s segment) length() int {
	return addVector2Ds(s.end.position, s.start.position.scale(-1)).abs().sum()
}

var levelOne = `
.>.v..XXXX
.......OO.
...>...OO.
..........
XXX..v..<.
XXX.......
...XX....X
.........X
.^...<...X
X.......^.
`
