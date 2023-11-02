/*
## Задача № 1
Написать API для указанных маршрутов(endpoints)
"/info"   // Информация об API
"/first"  // Случайное число
"/second" // Случайное число
"/add"    // Сумма двух случайных чисел
"/sub"    // Разность
"/mul"    // Произведение
"/div"    // Деление

*результат вернуть в виде JSON

"math/rand"
number := rand.Intn(100)
! не забудьте про Seed()


GET http://127.0.0.1:1234/first

GET http://127.0.0.1:1234/second

GET http://127.0.0.1:1234/add
GET http://127.0.0.1:1234/sub
GET http://127.0.0.1:1234/mul
GET http://127.0.0.1:1234/div
GET http://127.0.0.1:1234/info
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	 "math/rand"
	 "time"
)


type Calc struct {
	FirstNum int `json:"firstNum"`
	SecondNum int `json:"secondNum"`
	Result int `json:"result"`
}

var calc Calc

func main() {
	 http.HandleFunc("/info", GetInfo) 			// Информация об API
	 http.HandleFunc("/first", SetFirstNum) 	// Случайное число
	 http.HandleFunc("/second", SetSecondNum) 	// Случайное число
	 http.HandleFunc("/add", GetSumm) 			// Сумма двух случайных чисел
	 http.HandleFunc("/sub", GetDiff) 			// Разность
	 http.HandleFunc("/mul", GetCompos) 		// Произведение
	 http.HandleFunc("/div", GetDiv) 			// Деление

	log.Fatal(http.ListenAndServe("0.0.0.0:1234", nil))
}


//задаем первое число
func SetFirstNum(w http.ResponseWriter, r *http.Request) {
	calc.FirstNum = GenerateNum(100)
	calc.Result = 0
	fmt.Fprintln(w, GetJson())
}

//задаем второе число
func SetSecondNum(w http.ResponseWriter, r *http.Request) {
	calc.SecondNum = GenerateNum(100)
	calc.Result = 0
	fmt.Fprintln(w, GetJson())
}

//Сумма двух чисел
func GetSumm(w http.ResponseWriter, r *http.Request){
	calc.Result = calc.FirstNum + calc.SecondNum
	fmt.Fprintln(w, GetJson())
}

//Разность
func GetDiff(w http.ResponseWriter, r *http.Request){
	calc.Result = calc.FirstNum - calc.SecondNum
	fmt.Fprintln(w, GetJson())
}

//Произведение
func GetCompos(w http.ResponseWriter, r *http.Request){
	calc.Result = calc.FirstNum * calc.SecondNum
	fmt.Fprintln(w, GetJson())
}

//Деление
func GetDiv(w http.ResponseWriter, r *http.Request){
	calc.Result = calc.FirstNum / calc.SecondNum
	fmt.Fprintln(w, GetJson())
}

//берем pretty-print json
func GetJson() string {
	json, _ := json.MarshalIndent(calc, "", "  ")
   	return string(json)
}

//генератор случайных чисел
func GenerateNum(max int) int {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    return rand.Intn(max)
}

// информация
func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "## Задача № 1")
	fmt.Fprintln(w, "Написать API для указанных маршрутов(endpoints)")
	fmt.Fprintln(w, "\"/info\"   // Информация об API")
	fmt.Fprintln(w, "\"/first\"  // Случайное число")
	fmt.Fprintln(w, "\"/second\" // Случайное число")
	fmt.Fprintln(w, "\"/add\"    // Сумма двух случайных чисел")
	fmt.Fprintln(w, "\"/sub\"    // Разность")
	fmt.Fprintln(w, "\"/mul\"    // Произведение")
	fmt.Fprintln(w, "\"/div\"    // Деление")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "*результат вернуть в виде JSON")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "\"math/rand\"")
	fmt.Fprintln(w, "number := rand.Intn(100)")
	fmt.Fprintln(w, "! не забудьте про Seed()")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/first")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/second")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/add")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/sub")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/mul")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/div")
	fmt.Fprintln(w, "GET http://127.0.0.1:1234/info")
}