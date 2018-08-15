// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/testfixtures.v2"
)

// func main() {
// 	cmd.Execute()
// }

func main() {

	// DB接続先
	mysqlConnect := "root:password@tcp(127.0.0.1:13306)/sakila"
	// yml出力先
	generatePath := "testdata"

	mysql, err := sql.Open("mysql", mysqlConnect)
	if err != nil {
		panic(err.Error())
	}
	defer mysql.Close()

	err = testfixtures.GenerateFixtures(mysql, &testfixtures.MySQL{}, generatePath)
	if err != nil {
		panic(err.Error())
	}

}
