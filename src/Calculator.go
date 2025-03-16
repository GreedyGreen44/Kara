package main

import "strconv"

func calculateCoordinate(coordinate string, zoom int) (result string, resultCode int) {
	inputInt, err := strconv.ParseInt(coordinate, 10, 64)
	if err != nil {
		return "", handleError([2]byte{0x01, 0x03}, err)
	}
	outputFloat := float32(inputInt * 3 / (100000 * int64(zoom)))
	result = strconv.FormatFloat(float64(outputFloat), 'f', -1, 32)
	return result, resultCode
}
