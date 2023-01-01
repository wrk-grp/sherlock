package sherlock

import (
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/wrk-grp/errnie"
)

func (parser *Parser) table(v sqlparser.TableName) (bool, error) {
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
	return true, nil
}
