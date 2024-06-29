package database

import (
	"strings"

	"github.com/paas-ok/service/materials"
)

type MaterialFilter struct {
	Color []string
	Name  []string
}

func (store *Store) FindAllMaterials(filter MaterialFilter) ([]*materials.Material, error) {
	query := "SELECT id,type,price,density,color FROM materials WHERE 1=1"
	args := make([]interface{}, 0)

	if count := len(filter.Color); count > 0 {
		query += " AND color IN (" + strings.Repeat("?,", count-1) + "?)"
		for _, color := range filter.Color {
			args = append(args, color)
		}
	}

	if count := len(filter.Name); count > 0 {
		query += " AND name IN (" + strings.Repeat("?,", count-1) + "?)"
		for _, name := range filter.Name {
			args = append(args, name)
		}
	}

	query += " ORDER BY price ASC;"
	rows, err := store.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	mats := make([]*materials.Material, 0)
	for rows.Next() {
		material := &materials.Material{}
		err := rows.Scan(&material.Id, &material.Type, &material.Price, &material.Density, &material.Color)
		if err != nil {
			return nil, err
		}

		mats = append(mats, material)
	}

	return mats, nil
}
