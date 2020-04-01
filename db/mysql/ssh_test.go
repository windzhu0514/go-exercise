package main

// 通过ssh隧道连接数据库

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"testing"
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

func TestDBssh(t *testing.T) {

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
	robotID := "14001"

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

	key, err := ioutil.ReadFile("D:\\MyConfiguration\\ljc43026\\.ssh\\id_rsa_84")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	dialFunc := func(addr string) (net.Conn, error) {
		// return net.Dial("tcp", addr)

		config := ssh.ClientConfig{
			Timeout:         time.Second * 30,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			User:            "lijinchang",
			Auth: []ssh.AuthMethod{
				//ssh.Password("userpass"),
				ssh.PublicKeys(signer),
			},
		}

		client, err := ssh.Dial("tcp", "47.110.127.250:22", &config)
		if err != nil {
			log.Printf("Server dial error: %s\n", err)
			return nil, err
		}

		return client.Dial("tcp", "10.111.21.25:3306")
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
