package main

type VesselInfo struct {
	TimestampEst int64
	Mmsi         string
	Lat          float32
	Long         float32
	NameLength   int
	VesselName   string
	CourseEst    float32
	VesselType   uint8
}
