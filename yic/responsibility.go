package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

type responsibilities struct {
	p *playing

	byChain map[int][]*responsibility
}

func (resps *responsibilities) Init() {
	resps.byChain = make(map[int][]*responsibility)
	for chainIndex := range resps.p.levels.ChosenLevel().chains {
		resps.byChain[chainIndex] = make([]*responsibility, 0)
	}
	resps.byChain[0] = append(
		resps.byChain[0],
		&responsibility{
			typ:   responsibilityType1,
			life:  1500,
			speed: 2.5,
		},
	)
}

func (resps *responsibilities) Tick(ms int) *engine.Transition {
	factor := float64(ms) / 1000.0

	// Check for resps without health and remove them.
	for chainIndex := range resps.byChain {
		respsWithoutHealth := make(map[int]struct{})
		for i := range resps.byChain[chainIndex] {
			if resps.byChain[chainIndex][i].life <= 0 {
				respsWithoutHealth[i] = struct{}{}
			}
		}
		if len(respsWithoutHealth) > 0 {
			remaining := make([]*responsibility, 0)
			for i := range resps.byChain[chainIndex] {
				if _, ok := respsWithoutHealth[i]; !ok {
					remaining = append(remaining, resps.byChain[chainIndex][i])
				}
			}
			resps.byChain[chainIndex] = remaining
		}
	}

	// Move all resps, then check wether they reached head.
	for chainIndex := range resps.byChain {
		respsAtHead := make(map[int]struct{})
		for i := range resps.byChain[chainIndex] {
			var headReached bool
			resp := resps.byChain[chainIndex][i]
			resp.position += resp.speed * factor
			resp.x, resp.y, headReached = resps.p.levels.ChosenLevel().responsibilityPosition(chainIndex, resp.position)
			if headReached {
				respsAtHead[i] = struct{}{}
			}
		}
		if len(respsAtHead) > 0 {
			remaining := make([]*responsibility, 0, len(resps.byChain)-len(respsAtHead))
			for i := range resps.byChain[chainIndex] {
				resp := resps.byChain[chainIndex][i]
				if _, okRemove := respsAtHead[i]; okRemove {
					resps.p.head.receiveDamage(resp.life)
				} else {
					remaining = append(remaining, resp)
				}
			}
			resps.byChain[chainIndex] = remaining
		}
	}
	return nil
}

func (resps *responsibilities) Objects() []engine.Object {
	objects := make([]engine.Object, 0)
	for chainIndex := range resps.byChain {
		for i := range resps.byChain[chainIndex] {
			objects = append(
				objects,
				engine.Object{
					Key: resps.byChain[chainIndex][i].typ,
					X:   int(resps.byChain[chainIndex][i].x),
					Y:   int(resps.byChain[chainIndex][i].y),
				},
			)
		}
	}
	return objects
}

type responsibility struct {
	typ   string
	life  float64
	speed float64

	// position is the position of the responsibility on its chain.
	position float64

	// x and y are calculated via position.
	x float64
	y float64
}

func (r *responsibility) receiveDamage(dmg float64) {
	r.life -= dmg
}

const (
	responsibilityType1 = "responsibility_1"
	responsibilityType2 = "responsibility_2"
	responsibilityType3 = "responsibility_3"
)
