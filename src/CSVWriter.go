package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

func transformStruct(vessel VesselInfo) (vesselStr []string) {
	vesselStr = make([]string, 0)
	vesselStr = append(vesselStr, vessel.Mmsi)
	vesselStr = append(vesselStr, vessel.VesselName)
	vesselStr = append(vesselStr, strconv.FormatUint(uint64(vessel.VesselType), 16))
	vesselStr = append(vesselStr, strconv.FormatFloat(float64(vessel.Lat), 'f', -1, 32))
	vesselStr = append(vesselStr, strconv.FormatFloat(float64(vessel.Long), 'f', -1, 32))
	vesselStr = append(vesselStr, strconv.FormatFloat(float64(vessel.CourseEst), 'f', -1, 32))
	vesselStr = append(vesselStr, strconv.FormatUint(uint64(vessel.TimestampEst), 10))

	return vesselStr

}

func writeToCsv(vessels []VesselInfo, outputDirectory string) (resultCode int) {
	if len(vessels) == 0 {
		return handleError([2]byte{0x00, 0x04}, errors.New("no vessels found"))
	}

	writer, file, resultCode := createCSVWriter(outputDirectory)
	if resultCode != 0 {
		return resultCode
	}
	defer file.Close()

	badResultCode := 0
	writtenRecords := 0
	for _, vessel := range vessels {
		resultCode = writeCSVRecord(writer, transformStruct(vessel))
		if resultCode != 0 {
			badResultCode = resultCode
		} else {
			writtenRecords++
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return handleError([2]byte{0x00, 0x04}, err)
	}
	mainLog.Printf("Written %v records to csv file\n", writtenRecords)
	return badResultCode
}

func createCSVWriter(outputDirectory string) (writer *csv.Writer, file *os.File, resultCode int) {
	t := time.Now()
	fileName := outputDirectory + "/KaraOut_" + t.Format("20060102_15") + ".csv"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, handleError([2]byte{0x00, 0x04}, err)
	}

	writer = csv.NewWriter(file)
	writer.Comma = ';'
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, nil, handleError([2]byte{0x00, 0x04}, err)
	}
	if fileInfo.Size() == 0 {
		header := []string{"MMSI", "Name", "Type", "Lat", "Lon", "Course Est.", "Timestamp Est."}
		writeCSVRecord(writer, header)
	}
	return writer, file, 0
}

func writeCSVRecord(writer *csv.Writer, aircraft []string) (resultCode int) {
	err := writer.Write(aircraft)
	if err != nil {
		return handleError([2]byte{0x00, 0x04}, err)
	}
	return 0
}
