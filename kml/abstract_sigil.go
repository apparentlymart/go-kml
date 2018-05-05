package kml

// isObject should be embedded into all structs that implement Object, to
// ensure that only implementations from this package can be used.
type isObject struct {
}

func (o isObject) objectSigil() isObject {
	return isObject{}
}

// isGeometry should be embedded into all structs that implement Geometry, to
// ensure that only implementations from this package can be used.
type isGeometry struct {
	isObject
}

func (g isGeometry) objectSigil() isGeometry {
	return g
}

// isStyleSelector should be embedded into all structs that implement
// StyleSelector, to ensure that only implementations from this package can be
// used.
type isStyleSelector struct {
	isObject
}

func (ss isStyleSelector) styleSelectorSigil() isStyleSelector {
	return ss
}

// isTimePrimitive should be embedded into all structs that implement
// TimePrimitive, to ensure that only implementations from this package can be
// used.
type isTimePrimitive struct {
	isObject
}

func (tp isTimePrimitive) timePrimitiveSigil() isTimePrimitive {
	return tp
}

// isAbstractView should be embedded into all structs that implement
// AbstractView, to ensure that only implementations from this package can be
// used.
type isAbstractView struct {
	isObject
}

func (av isAbstractView) abstractViewSigil() isAbstractView {
	return av
}
