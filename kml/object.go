package kml

// object can be embedded into another struct to make it implement interface
// Object.
type object struct {
	isObject

	id       string
	targetID string
}

func (o object) ID() string {
	return o.id
}

func (o object) TargetID() string {
	return o.targetID
}
