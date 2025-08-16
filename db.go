package main

import (
	"database/sql"
	_ "github.com/glebarez/sqlite"
	"log"
)

var DB *sql.DB

func Init(dataPath string) {
	db, err := sql.Open("sqlite", dataPath)
	if err != nil {
		log.Println(err)
	}
	DB = db
	log.Println("Database initialized path:", dataPath)
}

type Result struct {
	Columns []string        `json:"columns"`
	Values  [][]interface{} `json:"values"`
}

func ExecQuery(param Param) ([]Result, error) {
	rows, err := DB.Query(param.Sql, param.Params...)
	if err != nil {
		log.Println(err)
		return []Result{}, err
	}
	if rows == nil {
		log.Println("无查询结构")
		return make([]Result, 0), nil
	}
	// 3. 解析查询结果
	var values [][]interface{}
	columns, _ := rows.Columns()
	var columnLength = len(columns)
	for rows.Next() {
		// 创建一个切片用于存放每列的值，每个值初始化为一个空接口
		valueList := make([]interface{}, columnLength)
		// 创建一个切片，每个元素是指向values中对应元素的指针
		pointers := make([]interface{}, columnLength)
		for i := range valueList {
			pointers[i] = &valueList[i]
		}
		if err := rows.Scan(pointers...); err != nil {
			log.Println(param.Sql)
			log.Println(err)
		}

		values = append(values, pointers)
	}
	var results []Result
	if len(values) == 0 {
		return make([]Result, 0), nil
	}
	results = append(results, Result{
		Columns: columns,
		Values:  values,
	})
	return results, nil
}

func ExecDel(param Param) {
	stmt, err := DB.Prepare(param.Sql)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(param.Params[0])
	if err != nil {
		log.Println(err)
	}
}

func ExecUpdate(param Param) error {
	stmt, err := DB.Prepare(param.Sql)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(param.Params...)
	if err != nil {
		log.Println(err)
	}
	return err
}
