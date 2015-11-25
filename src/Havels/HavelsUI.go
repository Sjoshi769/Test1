package main

import (
        "fmt"
        "net/http"
		"net/url"
		"strconv"
		"tcpUtils"
		"dbUtils"
		"log"
		//"time"
)


var (
	LampController 	tcpUtils.TcpUtilsStruct
	dbController	dbUtils.DbUtilsStruct
	)


func login(w http.ResponseWriter, r *http.Request) {

	var pw string
	
	passed := false
        
	r.ParseForm()
    // logic part of log in
		
    username := r.Form["username"]
    password := r.Form["password"]

    //fmt.Println(username[0])
    //fmt.Println(password[0])


	if (dbController.DbConnected) {

        //row, err := db.Query("select user_email, password from login where user_email=?",1)
        rows, err := dbController.Db.Query("select password from login where user_email=?",username[0])

        if err != nil {
			fmt.Println(err)
		}


		for rows.Next() {

			err := rows.Scan(&pw)

            if err != nil {
				fmt.Println(err)
			} else if (password[0] == pw) {
			    passed = true
			}
		}

		rows.Close()
	}


	

	if (passed) {
		http.Redirect(w, r,  "HavelsHome.html", http.StatusFound)

	} else {

		fmt.Fprintf(w, "<h1>Invalid User Name Or Password</h1>")
		//http.Error(w, "Invalid User Name Or Password\n",http.StatusInternalServerError)
		//http.Redirect(w, r,  "index.html", http.StatusFound)

		//fmt.Println("Timer started")
		//timer1 := time.NewTimer(time.Second * 5)
		//<-timer1.C
		//fmt.Println("Timer stoped")
		//http.Redirect(w, r,  "index.html", http.StatusFound)

	}

}



func LampControl(w http.ResponseWriter, r *http.Request) {
    
    //fmt.Println(r.URL)    
    //fmt.Println("I am inside")
	u, _ := url.Parse(r.URL.String())
	m, _ := url.ParseQuery(u.RawQuery)
    //fmt.Println(m)
    fmt.Println(m["lampCounter1"][0])
    fmt.Println(m["i"][0])
	LampId, err := strconv.Atoi(m["lampCounter1"][0])
	if err != nil {
		fmt.Println("Invalid lamp id " + m["lampCounter1"][0] + " specified")
		return
	}
	LampVal, err := strconv.Atoi(m["i"][0])
	if err != nil {
		fmt.Println("Invalid lamp contral val  " + m["i"][0] + " specified")
		return
	}

	fmt.Printf("LampId = %d, LampVal = %d\n",LampId, LampVal)
	LampController.SendLightControl(LampId,LampVal) 


}




func main() {						


	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbController.DbUtilsInit()

	if	(dbController.DbConnected) {
		defer dbController.Db.Close()
	}


	go LampController.InitTcpUtilsStruct()

    http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", login)
	http.HandleFunc("/LampControl",LampControl)
    http.ListenAndServe(":3001", nil)
}