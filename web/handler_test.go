package web

import (
	m "github.com/scotthelm/gominion/models"
	"log"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	x := IndexHandler(reflect.TypeOf(m.Campaign{}))
	log.Println(x)
}
