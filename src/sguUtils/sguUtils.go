package sguUtils



import (
    "fmt"
	"scuUtils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)
	   


type sguUtilsStruct struct {

	SGUID						uint64		//PRIMARY KEY
	//dbTransactionHandler		*sql.Tx
	SCUList						[]scuUtils
	Lattitude					double
	Longitude					double
	NumOfSCUs					int

}




									

