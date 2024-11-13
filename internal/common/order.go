package common

import (
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
)

type OrderBy struct {
	Columns []OrderByColumn
}

type OrderByColumn struct {
	Name string
	Desc bool
}

func (o OrderBy) ToClauseOrderBy() clause.OrderBy {
	return clause.OrderBy{
		Columns: lo.Map(o.Columns, func(item OrderByColumn, index int) clause.OrderByColumn {
			return clause.OrderByColumn{
				Column: clause.Column{
					Name: item.Name,
				},
				Desc: item.Desc,
			}
		}),
	}
}
