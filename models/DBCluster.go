package models

import (
	"database/sql"
	"fmt"
)

type DBCluster struct {
	SqlDB *sql.DB
}

func (dbCluster *DBCluster) QueryDataForMap(strSQL string) ([]map[string]interface{}, error) {
	rowData := make([]map[string]interface{}, 0)
	rows, err := dbCluster.SqlDB.Query(strSQL)
	if err != nil {
		return rowData, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return rowData, err
	}
	count := len(columns)

	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]

			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}

			entry[col] = v
		}
		rowData = append(rowData, entry)
	}
	return rowData, nil
}
func (dbCluster *DBCluster) QueryDataForPaged(pageModel *PageModel, strSQL string) (map[string]interface{}, error) {
	rowData := make([]map[string]interface{}, 0)

	err := dbCluster.SqlDB.QueryRow(fmt.Sprintf("SELECT COUNT(1) totalCount FROM (%s) tmppage1 ", strSQL)).Scan(&pageModel.RecordCount)
	if err != nil {
		return pageModel.Paginator(), err
	}
	rows, err := dbCluster.SqlDB.Query(fmt.Sprintf(" %s LIMIT %d OFFSET %d ", strSQL, pageModel.PageSize, (pageModel.CurrentPage-1)*pageModel.PageSize))
	if err != nil {
		return pageModel.Paginator(), err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return pageModel.Paginator(), err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]

			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}

			entry[col] = v
		}
		rowData = append(rowData, entry)
	}
	pageModel.Records = rowData
	return pageModel.Paginator(), nil
}
