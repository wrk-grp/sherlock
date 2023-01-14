package sherlock

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func (parser *Parser) table(v sqlparser.TableName) string {
	return v.Name.String()
}
