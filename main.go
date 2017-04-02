package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"flashCoder/flashdb"
)

func main() {
	flashdb:=new(flashdb.FMyDB)
	flashdb.Init("root:huidai021@/flashCoder")
	sql:="SELECT behavior FROM flash_behavior WHERE name = ?"
	condition:=[]interface{}{13}
	var squareNum int
	res:=[]interface{}{&squareNum}
	err:=flashdb.SelectOne(sql,condition,res)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(squareNum)
	flashdb.Close()
	//
	//db, err := sql.Open("mysql", "root:huidai021@/flashCoder")
	//if err != nil {
	//	panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	//}
	//defer db.Close()
	//
	//// Prepare statement for inserting data
	//stmtIns, err := db.Prepare("INSERT INTO flash_behavior VALUES(1,?,?)") // ? = placeholder
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	//
	//// Prepare statement for reading data
	//stmtOut, err := db.Prepare("SELECT behavior FROM flash_behavior WHERE name = ?")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//defer stmtOut.Close()
	//
	//// Insert square numbers for 0-24 in the database
	//for i := 0; i < 25; i++ {
	//	_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
	//	if err != nil {
	//		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
	//}
	//
	//var squareNum int // we "scan" the result in here
	//
	//// Query the square-number of 13
	//row:=stmtOut.QueryRow(13)
	//fmt.Println(row)
	//err = row.Scan(&squareNum) // WHERE number = 13
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Printf("The square number of 13 is: %d", squareNum)
	//
	//// Query another number.. 1 maybe?
	//err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Printf("The square number of 1 is: %d", squareNum)
      //
	//rows,err := db.Query("SELECT * FROM flash_behavior")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//// Get column names
	//columns, err := rows.Columns()
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
      //fmt.Println(columns)
	//// Make a slice for the values
	//values := make([]sql.RawBytes, len(columns))
      //
	//// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	//// references into such a slice
	//// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	//scanArgs := make([]interface{}, len(values))
	//for i := range values {
	//	scanArgs[i] = &values[i]
	//}
	//// Fetch rows
	//for rows.Next() {
	//	// get RawBytes from data
	//	err = rows.Scan(scanArgs...)
	//	if err != nil {
	//		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
      //
	//	// Now do something with the data.
	//	 //Here we just print each column as a string.
	//	var value string
	//	for i, col := range values {
	//		// Here we can check if the value is nil (NULL value)
	//		if col == nil {
	//			value = "NULL"
	//		} else {
	//			value = string(col)
	//		}
	//		fmt.Println(columns[i], ": ", value)
	//	}
	//	fmt.Println("-----------------------------------")
	//}
	//if err = rows.Err(); err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
}