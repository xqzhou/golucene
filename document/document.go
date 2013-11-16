package document

type Document struct {
	Fields []*Field
	Count  int
}

func NewDocument() *Document {
	return &Document{
		Fields: make([]*Field, 10),
		Count:  0,
	}
}

func (d *Document) AddField(field *Field) {
	d.Fields[d.Count] = field
	d.Count = d.Count + 1
}
