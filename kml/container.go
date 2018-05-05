package kml

// container can be embedded in another struct to make it implement
// the methods unique to interface Container.
//
// All containers are also required to be features, so a type embedding this
// should also embed feature.
type container struct {
	open     bool
	features []Feature
}

func (c container) Open() bool {
	return c.open
}

func (c container) Features() []Feature {
	return c.features
}

func (c container) SubContainers() []Container {
	if len(c.features) == 0 {
		return nil
	}
	ret := make([]Container, 0, len(c.features))
	for _, f := range c.features {
		c, ok := f.(Container)
		if !ok {
			continue
		}
		ret = append(ret, c)
	}
	if len(ret) == 0 {
		return nil
	}
	return ret
}

type Folder struct {
	object
	feature
	container
}

type Document struct {
	object
	feature
	container
	schemas []*Schema
	styles  []StyleSelector
}

// Schemas returns all of the schemas associated with this document.
//
// The caller must not mutate the backing array of the returned slice.
func (d *Document) Schemas() []*Schema {
	return d.schemas
}

// Schema returns the first schema in the schema list with the given id, or
// nil if no such schema is found.
func (d *Document) Schema(id string) *Schema {
	for _, schema := range d.schemas {
		if schema.ID == id {
			return schema
		}
	}
	return nil
}
