package kml

// Schema represents a custom data schema, usually embedded in a Document.
type Schema struct {
	ID     string
	Fields []SchemaField
}

// SchemaField represents a single field in a Schema.
type SchemaField struct {
	Type        SchemaType
	Name        string
	DisplayName string
}

// SchemaType is an enumeration for the type of a SchemaField.
type SchemaType rune

const (
	SchemaString SchemaType = 'x'
	SchemaInt    SchemaType = 'i'
	SchemaUInt   SchemaType = 'I'
	SchemaShort  SchemaType = 's'
	SchemaUShort SchemaType = 'S'
	SchemaFloat  SchemaType = 'f'
	SchemaDouble SchemaType = 'F'
	SchemaBool   SchemaType = 'b'
)
