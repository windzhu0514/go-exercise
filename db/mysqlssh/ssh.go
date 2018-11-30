package mysqlssh

// 通过ssh隧道连接数据库

import (
	"database/sql"
	"log"
	"net"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

type config struct {
	Robotcates map[string]struct {
		Server   string
		Username string
		Userpass string
	}
}

func TestDBssh() {

	var db *sql.DB

	log.SetFlags(log.Lshortfile)
	// flag.Parse()
	// if flag.NArg() < 1 {
	// 	log.Println("give me a robot id")
	// 	return
	// }
	//
	// var err error
	// robotID, err := strconv.Atoi(flag.Arg(0))
	// if err != nil {
	// 	log.Println("give me a robot id")
	// 	return
	// }
	robotID := "58959"

	var conf config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Println(err)
		return
	}

	database := conf.Robotcates["bus"]
	dns := database.Username + ":" + database.Userpass + "@tcp(" + database.Server + ":3306)/busticket?charset=utf8"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Println(err)
		return
	}

	dialFunc := func(addr string) (net.Conn, error) {
		// return net.Dial("tcp", addr)

		config := ssh.ClientConfig{
			Timeout:         time.Second * 30,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			User:            "username",
			Auth: []ssh.AuthMethod{
				ssh.Password("userpass"),
			},
		}

		client, err := ssh.Dial("tcp", "121.40.34.84:22", &config)
		if err != nil {
			log.Printf("Server dial error: %s\n", err)
			return nil, err
		}

		return client.Dial("tcp", "10.26.252.47:3306")
	}

	mysql.RegisterDial("tcp", dialFunc)

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}

	var robotURL string
	err = db.QueryRow("SELECT url FROM robots WHERE id=?", robotID).Scan(&robotURL)
	if err != nil {
		log.Println(err)
		return
	}

	robotURL = strings.TrimPrefix(robotURL, "http://")

	log.Println(robotURL)
}
