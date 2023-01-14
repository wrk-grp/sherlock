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
	JOINQUERY       = "SELECT * FROM tests LEFT JOIN joins ON tests.join_id = joins.id"
)

var (
	DATAGRAM = func(query string) *spd.Datagram {
		dg := spd.New(spd.APPTXT, spd.QUESTION, spd.DATALAKE)
		buf := bytes.NewBufferString(query)
		io.Copy(dg, buf)
		return dg
	}
)

func TestNewParser(t *testing.T) {
	Convey(GIVENANINSTANCE, t, func() {
		parser := NewParser(DATAGRAM(SIMPLEQUERY))

		Convey("It should be correctly initialized", func() {
			So(parser.err, ShouldBeNil)
		})

		Convey("When joining another scope", func() {
			parser = NewParser(DATAGRAM(JOINQUERY))

			Convey("It should return the joined results", func() {

			})
		})
	})
}

func TestToPrefix(t *testing.T) {
	Convey(GIVENANINSTANCE, t, func() {
		parser := NewParser(DATAGRAM(SIMPLEQUERY))

		Convey("It should walk the subtree", func() {
			parser.ToPrefix()
		})

		Convey("When joining another scope", func() {
			parser = NewParser(DATAGRAM(JOINQUERY))

			Convey("It should return a dynamic prefix", func() {
				parser.ToPrefix()
			})
		})
	})
}

func BenchmarkNewParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewParser(DATAGRAM(SIMPLEQUERY))
	}
}
