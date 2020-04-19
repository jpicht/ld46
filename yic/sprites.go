package yic

import (
	"github.com/GodsBoss/ld46/pkg/ui"
)

func Sprites() map[string]ui.Sprite {
	return map[string]ui.Sprite{
		"field_way": ui.Sprite{
			X:      0,
			Y:      36,
			W:      fieldSize.X,
			H:      fieldSize.Y,
			Frames: 1,
		},
		"field_buildspot": ui.Sprite{
			X:      0,
			Y:      0,
			W:      fieldSize.X,
			H:      fieldSize.Y,
			Frames: 1,
		},
		"field_obstacle": ui.Sprite{
			X:      0,
			Y:      18,
			W:      fieldSize.X,
			H:      fieldSize.Y,
			Frames: 1,
		},
		"head_toddler": ui.Sprite{
			X:      72,
			Y:      0,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
		"head_child": ui.Sprite{
			X:      72,
			Y:      36,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
		"head_teen": ui.Sprite{
			X:      72,
			Y:      72,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
		"head_adult": ui.Sprite{
			X:      216,
			Y:      0,
			W:      fieldSize.X * 2,
			H:      fieldSize.Y * 2,
			Frames: 4,
		},
		"bg_title": ui.Sprite{
			X:      400,
			Y:      0,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_playing": ui.Sprite{
			X:      400,
			Y:      300,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_level_select": ui.Sprite{
			X:      400,
			Y:      600,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_game_over": ui.Sprite{
			X:      400,
			Y:      900,
			W:      400,
			H:      300,
			Frames: 1,
		},
		"bg_hiscore": ui.Sprite{
			X:      400,
			Y:      1200,
			W:      400,
			H:      300,
			Frames: 1,
		},
		responsibilityType1: ui.Sprite{
			X:      0,
			Y:      60,
			W:      16,
			H:      16,
			Frames: 4,
		},
		responsibilityType2: ui.Sprite{
			X:      0,
			Y:      76,
			W:      16,
			H:      16,
			Frames: 4,
		},
		responsibilityType3: ui.Sprite{
			X:      0,
			Y:      92,
			W:      16,
			H:      16,
			Frames: 4,
		},
		responsibilityType4: ui.Sprite{
			X:      0,
			Y:      108,
			W:      16,
			H:      16,
			Frames: 4,
		},
		"grid_cursor": ui.Sprite{
			X:      18,
			Y:      0,
			W:      18,
			H:      18,
			Frames: 1,
		},
		"grid_cursor_yes": ui.Sprite{
			X:      18,
			Y:      18,
			W:      18,
			H:      18,
			Frames: 1,
		},
		"grid_cursor_no": ui.Sprite{
			X:      18,
			Y:      36,
			W:      18,
			H:      18,
			Frames: 1,
		},
		"building_income": ui.Sprite{
			X:      72,
			Y:      124,
			W:      18,
			H:      18,
			Frames: 2,
		},
		"building_flame": ui.Sprite{
			X:      72,
			Y:      142,
			W:      18,
			H:      18,
			Frames: 2,
		},
		"building_gun": ui.Sprite{
			X:      72,
			Y:      160,
			W:      18,
			H:      18,
			Frames: 2,
		},
		gunShot: ui.Sprite{
			X:      108,
			Y:      108,
			W:      8,
			H:      8,
			Frames: 8,
		},
		gunHit: ui.Sprite{
			X:      108,
			Y:      116,
			W:      8,
			H:      8,
			Frames: 8,
		},
		"char_a": ui.Sprite{
			X:      0,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_b": ui.Sprite{
			X:      5,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_c": ui.Sprite{
			X:      10,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_d": ui.Sprite{
			X:      15,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_e": ui.Sprite{
			X:      20,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_f": ui.Sprite{
			X:      25,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_g": ui.Sprite{
			X:      30,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_h": ui.Sprite{
			X:      35,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_i": ui.Sprite{
			X:      40,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_j": ui.Sprite{
			X:      45,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_k": ui.Sprite{
			X:      50,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_l": ui.Sprite{
			X:      55,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_m": ui.Sprite{
			X:      60,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_n": ui.Sprite{
			X:      65,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_o": ui.Sprite{
			X:      70,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_p": ui.Sprite{
			X:      75,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_q": ui.Sprite{
			X:      80,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_r": ui.Sprite{
			X:      85,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_s": ui.Sprite{
			X:      90,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_t": ui.Sprite{
			X:      95,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_u": ui.Sprite{
			X:      100,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_v": ui.Sprite{
			X:      105,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_w": ui.Sprite{
			X:      110,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_x": ui.Sprite{
			X:      115,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_y": ui.Sprite{
			X:      120,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_z": ui.Sprite{
			X:      125,
			Y:      204,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_0": ui.Sprite{
			X:      0,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_1": ui.Sprite{
			X:      5,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_2": ui.Sprite{
			X:      10,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_3": ui.Sprite{
			X:      15,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_4": ui.Sprite{
			X:      20,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_5": ui.Sprite{
			X:      25,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_6": ui.Sprite{
			X:      30,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_7": ui.Sprite{
			X:      35,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_8": ui.Sprite{
			X:      40,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_9": ui.Sprite{
			X:      45,
			Y:      209,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_.": ui.Sprite{
			X:      0,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_,": ui.Sprite{
			X:      5,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_:": ui.Sprite{
			X:      10,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_;": ui.Sprite{
			X:      15,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_!": ui.Sprite{
			X:      20,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_?": ui.Sprite{
			X:      25,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_(": ui.Sprite{
			X:      30,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_)": ui.Sprite{
			X:      35,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_-": ui.Sprite{
			X:      40,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char__": ui.Sprite{
			X:      45,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_<": ui.Sprite{
			X:      50,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_>": ui.Sprite{
			X:      55,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
		"char_ ": ui.Sprite{
			X:      60,
			Y:      214,
			W:      5,
			H:      5,
			Frames: 1,
		},
	}
}
