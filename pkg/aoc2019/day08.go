package aoc2019

/**
 * <a href="https://adventofcode.com/2019/day/8">Day 8</a>
 */

import (
	"strings"
)

type Day08 struct{}

func (d Day08) layerImages(input string) (int, string) {
	layers := d.toLayers(input)
	return d.zeroLayer(layers), d.drawImage(layers)
}

func (d Day08) zeroLayer(layers []string) int {
	zeroes := 150
	zeroLayer := ""
	for _, layer := range layers {
		zeroCount := strings.Count(layer, "0")
		if zeroCount < zeroes {
			zeroes = zeroCount
			zeroLayer = layer
		}
	}

	return strings.Count(zeroLayer, "1") * strings.Count(zeroLayer, "2")
}

func (d Day08) drawImage(layers []string) string {
	decoded := []rune(layers[0])
	for i := 1; i < len(layers); i++ {
		layer := layers[i]
		for p := 0; p < len(layer); p++ {
			if decoded[p] == '2' {
				decoded[p] = rune(layer[p])
			}
		}
	}

	return string(decoded)
}

func (d Day08) toLayers(input string) []string {
	count := len(input) / 150 // The image you received is 25 pixels wide and 6 pixels tall.
	layers := make([]string, count)
	for i := 0; i < count; i++ {
		next := i * 150
		layers[i] = input[next:(next + 150)]
	}
	return layers
}
