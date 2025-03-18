package main

import (
	"encoding/binary"
	"errors"
	"strconv"
	"time"
)

func processData(data []byte, zoom int) (vessels []VesselInfo, resultCode int) {
	const headerSize = 12
	dataSize := len(data)

	if dataSize < headerSize {
		return nil, handleError([2]byte{0x00, 0x03}, errors.New("invalid size of response data"))
	}

	unixTimestamp := time.Now().Unix()
	var offset int
	offset = headerSize

	vessels = make([]VesselInfo, 0)

	for offset < dataSize {
		addHeaderLength := 0
		if offset+10 >= dataSize {
			break
		}

		if (data[offset] & 0x80) == 0x00 {
			addHeaderLength = 10
		}

		latitude, resultCode := calculateFloatCoordinate(strconv.FormatInt(int64(int32(binary.BigEndian.Uint32(data[offset+addHeaderLength+6:offset+addHeaderLength+10]))), 10), zoom)
		if resultCode != 0 {
			continue
		}

		longitude, resultCode := calculateFloatCoordinate(strconv.FormatInt(int64(int32(binary.BigEndian.Uint32(data[offset+addHeaderLength+10:offset+addHeaderLength+14]))), 10), zoom)
		if resultCode != 0 {
			continue
		}

		nameLength := int(data[offset+addHeaderLength+15])
		frameLength := 16 + nameLength + addHeaderLength

		seen := data[offset+addHeaderLength+14]
		var timeSt int64
		if seen&0x80 == 0 {
			timeSt = unixTimestamp - int64(seen&0x3F)
		} else {
			timeSt = unixTimestamp - int64(seen&0x3F)*3600
		}

		var estimatedCourse float32

		if data[offset+addHeaderLength]&0x20 == 0 {
			estimatedCourse = float32(data[offset+addHeaderLength]&0x1F) * 11.25
		} else {
			estimatedCourse = -1
		}

		vessels = append(vessels, VesselInfo{
			TimestampEst: timeSt,
			Mmsi:         strconv.FormatInt(int64(binary.BigEndian.Uint32(data[offset+addHeaderLength+2:offset+addHeaderLength+6])), 10),
			Lat:          latitude,
			Long:         longitude,
			NameLength:   nameLength,
			VesselName:   string(data[offset+addHeaderLength+16 : offset+frameLength]),
			CourseEst:    estimatedCourse,
			VesselType:   data[offset+addHeaderLength+1] & 0xF0,
		})

		offset += frameLength
	}
	return vessels, 0
}
