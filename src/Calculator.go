package main

import "strconv"

func calculateFloatCoordinate(coordinate string, zoom int) (outputFloat float32, resultCode int) {
	inputInt, err := strconv.ParseInt(coordinate, 10, 64)
	if err != nil {
		return 0, handleError([2]byte{0x01, 0x03}, err)
	}
	outputFloat = float32(inputInt) * 3 / (100000 * float32(zoom))
	return outputFloat, resultCode
}

func calculateIntCoordinate(coordinate string, zoom int) (result string, resultCode int) {
	inputFloat, err := strconv.ParseFloat(coordinate, 32)
	if err != nil {
		return "", handleError([2]byte{0x01, 0x03}, err)
	}
	outputInt := int32(inputFloat * 100000 * float64(zoom) / 3)
	result = strconv.FormatInt(int64(outputInt), 10)
	return result, resultCode
}
