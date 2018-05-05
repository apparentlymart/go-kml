package kml

// feature can be embedded into other structs to make them automatically
// implement the Feature interface, with the exception of the methods
// inherited from Object.
type feature struct {
	name         string
	visible      bool
	address      string
	phoneNumber  string
	snippet      string
	description  string
	viewpoint    AbstractView
	time         TimePrimitive
	style        StyleSelector
	region       *Region
	extendedData *ExtendedData
}

func (f *feature) Name() string {
	return f.name
}

func (f *feature) Visible() bool {
	return f.visible
}

func (f *feature) Address() string {
	return f.address
}

func (f *feature) PhoneNumber() string {
	return f.phoneNumber
}

func (f *feature) Snippet() string {
	return f.snippet
}

func (f *feature) Description() string {
	return f.description
}

func (f *feature) Viewpoint() AbstractView {
	return f.viewpoint
}

func (f *feature) Time() TimePrimitive {
	return f.time
}

func (f *feature) Style() StyleSelector {
	return f.style
}

func (f *feature) Region() *Region {
	return f.region
}

func (f *feature) ExtendedData() *ExtendedData {
	return f.extendedData
}
