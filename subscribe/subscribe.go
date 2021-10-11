package subscribe

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"post/publish"
	"strconv"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

const (
	host = "192.168.15.128"
	// host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "123"
	dbname   = "mydb"
)

func Get(c echo.Context) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	map1 := make(map[string]int)
	rand.Seed(time.Now().UTC().UnixNano())

	sqlStatement := `
	INSERT INTO user03 (value, time)
	VALUES ($1, $2)
	RETURNING devideid`
	for i := 1; i <= 5; i++ {
		value := randInt(-40, 120)
		now := time.Now().Format("01/02/2006 15:04:05")
		id := 0
		map1[now] = value
		err = db.QueryRow(sqlStatement, value, now).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("New record devideid is:", id)
		devideid := "Devide" + strconv.Itoa(i) + "/temperature."
		publish.Publish(devideid, strconv.Itoa(value))
		time.Sleep(10 * time.Second)
	}
	return c.JSON(http.StatusOK, map1)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
