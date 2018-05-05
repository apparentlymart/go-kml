package kml

// Object is an interface implemented by all of the model types in this package.
// It is not particularly useful on its own, but is included to align with
// the KML specification.
//
// Object is a closed interface, implementable only by types within this package.
type Object interface {
	ID() string
	TargetID() string

	objectSigil() isObject
}

// Feature is an interface implemented by all of the concrete feature types.
type Feature interface {
	Object

	Name() string
	Visible() bool
	Address() string
	PhoneNumber() string
	Snippet() string
	Description() string
	Viewpoint() AbstractView
	Time() TimePrimitive
	Style() StyleSelector
	Region() *Region
	ExtendedData() *ExtendedData
}

// Container is an interface implemented by the container types (Folder and Document)
type Container interface {
	Feature

	Open() bool
	Features() []Feature

	// SubContainers is similar to Features but it filters the contained features
	// to include only those that are themselves containers.
	SubContainers() []Container
}

// Overlay is an interface implemented by the overlay types (PhotoOverlay,
// ScreenOverlay and GroundOverlay).
type Overlay interface {
	Feature

	Color() Color
	DrawOrder() int
	Icon() Icon
}

// Geometry is an interface implemented by all of the concrete geometry types.
type Geometry interface {
	Object

	// The Geometry type has no public-facing functionality of its own, so
	// we use a sigil method here to distinguish it from a plain Object.
	geometrySigil() isGeometry
}

// StyleSelector is an interface implemented by the style definition types
// (Style and StyleMap).
type StyleSelector interface {
	Object

	// The StyleSelector type has no public-facing functionality of its own, so
	// we use a sigil method here to distinguish it from a plain Object.
	styleSelectorSigil()
}

// TimePrimitive is an interface implemented by the time primitive types
// (TimeSpan and TimeStamp).
type TimePrimitive interface {
	Object

	// The TimePrimitive type has no public-facing functionality of its own, so
	// we use a sigil method here to distinguish it from a plain Object.
	timePrimitiveSigil()
}

// AbstractView is an interface implemented by the view types (Camera and LookAt)
type AbstractView interface {
	Object

	// The AbstractView type has no public-facing functionality of its own, so
	// we use a sigil method here to distinguish it from a plain Object.
	abstractViewSigil()
}
