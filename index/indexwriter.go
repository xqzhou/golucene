package index

import (
	"github.com/xqzhou/golucene/document"
	"github.com/xqzhou/golucene/store"
)

type IndexWriter struct {
	directory store.Directory
}

func NewIndexWriter(directory store.Directory) *IndexWriter {
	return &IndexWriter{
		directory: directory,
	}
}

func (w *IndexWriter) AddDocument(d *document.Document) {

}

func (w *IndexWriter) Close() {

}
