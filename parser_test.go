package sherlock

import (
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/theapemachine/wrkspc/spd"
)

func TestNewParser(t *testing.T) {
	Convey("Given a new instance", t, func() {
		dg := spd.New(spd.APPTXT, spd.QUESTION, spd.DATALAKE)
		buf := bytes.NewBufferString("SELECT * FROM tests")
		io.Copy(dg, buf)

		parser := NewParser(dg)

		Convey("It should be correctly initialized", func() {
			So(parser.err, ShouldBeNil)
		})
	})
}

func TestToPrefix(t *testing.T) {
	Convey("Given a new instance", t, func() {
		dg := spd.New(spd.APPTXT, spd.QUESTION, spd.DATALAKE)
		buf := bytes.NewBufferString("SELECT * FROM tests")
		io.Copy(dg, buf)

		parser := NewParser(dg)

		Convey("It should walk the subtree", func() {
			parser.ToPrefix()
		})
	})
}

func BenchmarkNewParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dg := spd.New(spd.APPTXT, spd.QUESTION, spd.DATALAKE)
		buf := bytes.NewBufferString("SELECT * FROM tests")
		io.Copy(dg, buf)

		_ = NewParser(dg)
	}
}
