package golucene

import (
	"github.com/xqzhou/golucene/document"
	"github.com/xqzhou/golucene/index"
	"github.com/xqzhou/golucene/store"
	"testing"
)

func TestIndex(t *testing.T) {
	directory := store.NewRAMDirectory()
	indexWriter := index.NewIndexWriter(directory)
	d := document.NewDocument()
	d.AddField(document.NewField("id", "1"))
	d.AddField(document.NewField("country", "2"))
	indexWriter.AddDocument(d)
	indexWriter.Close()
}
