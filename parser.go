package sherlock

import (
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/theapemachine/wrkspc/spd"
	"github.com/wrk-grp/errnie"
)

type Parser struct {
	datagram *spd.Datagram
	stmt     sqlparser.Statement
	err      error
}

func NewParser(datagram *spd.Datagram) *Parser {
	buf, err := datagram.ReadAt(0)
	errnie.Handles(err)

	errnie.Informs("SQL ->", string(buf))

	stmt, err := sqlparser.Parse(string(buf))
	errnie.Handles(err)

	return &Parser{datagram, stmt, nil}
}

func (parser *Parser) ToPrefix() {
	parser.stmt.WalkSubtree(
		func(node sqlparser.SQLNode) (kontinue bool, err error) {
			if node == nil {
				return false, nil
			}

			switch v := node.(type) {
			case nil:
				return false, nil
			case sqlparser.TableName:
				var identity []byte
				name := v.Name

				if name.IsEmpty() {
					return true, nil
				}

				identity, parser.err = parser.datagram.Identity()

				prefix := strings.Join([]string{
					string(identity),
					strings.ReplaceAll(name.String(), "_", "/"),
				}, "/")

				errnie.Informs("mapping table name to prefix:", prefix)
			default:
			}

			return true, nil
		},
	)
}
