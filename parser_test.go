package sherlock

import (
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/theapemachine/wrkspc/spd"
)

const (
	GIVENANINSTANCE = "Given an instance"
	SIMPLEQUERY     = "SELECT * FROM tests"
)

func TestNewParser(t *testing.T) {
	Convey(GIVENANINSTANCE, t, func() {
		dg := spd.New(spd.APPTXT, spd.QUESTION, spd.DATALAKE)
		buf := bytes.NewBufferString(SIMPLEQUERY)
		io.Copy(dg, buf)

		parser := NewParser(dg)

		Convey("It should be correctly initialized", func() {
			So(parser.err, ShouldBeNil)
		})
	})
}

func TestToPrefix(t *testing.T) {
	Convey(GIVENANINSTANCE, t, func() {
		dg := spd.New(spd.APPTXT, spd.QUESTION, spd.DATALAKE)
		buf := bytes.NewBufferString(SIMPLEQUERY)
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
		buf := bytes.NewBufferString(SIMPLEQUERY)
		io.Copy(dg, buf)

		_ = NewParser(dg)
	}
}
