package index

import (
	"github.com/xqzhou/golucene/document"
	"github.com/xqzhou/golucene/store"
)

type DocumentWriter struct {
	fieldInfos FieldInfos
}

func NewDocumentWriter(directory store.Directory) *DocumentWriter {
	return &DocumentWriter{
		FieldInfos: NewFieldInfos(),
	}
}

func (w *DocumentWriter) AddDocument(d *document.Document) {
	w.fieldInfos.AddDocument(d)
}

func (w *DocumentWriter) Flush() {

}
