package index

import (
	"github.com/xqzhou/golucene/document"
	"github.com/xqzhou/golucene/store"
)

type IndexWriter struct {
	directory      store.Directory
	documentWriter *DocumentWriter
}

func NewIndexWriter(directory store.Directory) *IndexWriter {
	return &IndexWriter{
		directory:      directory,
		documentWriter: NewDocumentWriter(directory),
	}
}

func (w *IndexWriter) AddDocument(d *document.Document) {
	w.documentWriter.AddDocument(d)
}

func (w *IndexWriter) Close() {
	w.documentWriter.Flush()
}
