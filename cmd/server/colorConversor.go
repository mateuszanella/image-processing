package main

import (
	"math"
	"strconv"
)

func HexToRGB(hex string) (int, int, int, error) {
	red, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	green, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	blue, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	return int(red), int(green), int(blue), nil
}

func HexToHSL(hex string) (float64, float64, float64, error) {
	r, g, b, err := HexToRGB(hex)
	if err != nil {
		return 0, 0, 0, err
	}

	h, s, l := RGBToHSL(r, g, b)

	return h, s, l, nil
}

func RGBToHex(red, green, blue int) string {
	return "#" + strconv.FormatInt(int64(red), 16) + strconv.FormatInt(int64(green), 16) + strconv.FormatInt(int64(blue), 16)
}

func RGBToHSL(red, green, blue int) (float64, float64, float64) {
	rNorm := float64(red) / 255
	gNorm := float64(green) / 255
	bNorm := float64(blue) / 255

	min := math.Min(math.Min(rNorm, gNorm), bNorm)
	max := math.Max(math.Max(rNorm, gNorm), bNorm)
	diff := max - min

	var h, s, l float64
	l = (max + min) / 2

	if diff == 0 {
		h, s = 0, 0
	} else {
		if l > 0.5 {
			s = diff / (2 - max - min)
		} else {
			s = diff / (max + min)
		}

		switch max {
		case rNorm:
			h = (gNorm - bNorm) / diff
			if gNorm < bNorm {
				h += 6
			}
		case gNorm:
			h = (bNorm-rNorm)/diff + 2
		case bNorm:
			h = (rNorm-gNorm)/diff + 4
		}
		h /= 6
	}

	return h, s, l
}

func HSLToHex(h, s, l float64) string {
	r, g, b := HSLToRGB(h, s, l)
	return RGBToHex(r, g, b)
}

func HSLToRGB(h, s, l float64) (int, int, int) {
	var r, g, b float64

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod((h/60), 2)-1))
	m := l - (c / 2)

	if h < 60 {
		r = c
		g = x
		b = 0
	} else if h < 120 {
		r = x
		g = c
		b = 0
	} else if h < 180 {
		r = 0
		g = c
		b = x
	} else if h < 240 {
		r = 0
		g = x
		b = c
	} else if h < 300 {
		r = x
		g = 0
		b = c
	} else {
		r = c
		g = 0
		b = x
	}

	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return int(math.Floor(r)), int(math.Floor(g)), int(math.Floor(b))
}
