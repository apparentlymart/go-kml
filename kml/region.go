package kml

type Region struct {
	Bounds        LatLonAltBox
	LevelOfDetail LevelOfDetail
}

type LatLonAltBox struct {
	AltitudeMode AltitudeMode
	MinAltitude  float64
	MaxAltitude  float64
	North        float64
	South        float64
	East         float64
	West         float64
}

type LevelOfDetail struct {
	MinPixels     int
	MaxPixels     int
	MinFadeExtent int
	MaxFadeExtent int
}

// AltitudeMode is an enumeration used with LatLonAltBox
type AltitudeMode rune

const (
	AltitudeUnset              AltitudeMode = 0
	AltitudeAbsolute           AltitudeMode = 'A'
	AltitudeRelativeToGround   AltitudeMode = 'R'
	AltitudeClampToGround      AltitudeMode = 'C'
	AltitudeRelativeToSeaFloor AltitudeMode = 'r'
	AltitudeClampToSeaFloor    AltitudeMode = 'c'
)
