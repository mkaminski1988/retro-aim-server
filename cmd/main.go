package main

import (
	"fmt"
	"github.com/mkaminski/goaim/oscar"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const testFile string = "/Users/mike/dev/goaim/aim.db"

func main() {

	sm := oscar.NewSessionManager()
	fm, err := oscar.NewFeedbagStore(testFile)
	if err != nil {
		log.Fatal(err)
	}

	go listenBOS(sm, fm)
	go listenStats()
	go listenAlert()
	go listenOdir()

	//todo implement CHATNAV and ALERT

	// Listen on TCP port 5190
	listener, err := net.Listen("tcp", ":5190")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 5190")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Handle connection in a separate goroutine
		go handleAuthConnection(sm, conn)
	}
}

func webServer(ch chan string) {
	http.HandleFunc("/send-im", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		ch <- string(body)
	})

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err.Error())
	}
}

func listenBOS(sm *oscar.SessionManager, fm *oscar.FeedbagStore) {
	// Listen on TCP port 5190
	listener, err := net.Listen("tcp", ":5191")
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string, 1)
	go webServer(ch)

	defer listener.Close()

	fmt.Println("Server is listening on port 5191")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		seq := uint32(100)
		go handleBOSConnection(sm, fm, conn, &seq)
		go sendIM(conn, ch, &seq)
	}
}

func sendIM(conn net.Conn, ch chan string, seq *uint32) {
	for msg := range ch {
		fmt.Printf("sending im... %s\n", msg)
		vals := strings.Split(msg, ":")

		if err := oscar.SendIM(conn, seq, vals[0], vals[1]); err != nil {
			//panic(err.Error())
		}
		if err := oscar.SetBuddyArrived(conn, seq, vals[0]); err != nil {
			//panic(err.Error())
		}
	}
}

func listenStats() {
	// Listen on TCP port 5190
	listener, err := net.Listen("tcp", ":5192")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 5192")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("got a connection on listenStats")
		seq := uint32(100)
		if err := oscar.ReadBos(nil, nil, conn, &seq); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}
}

func listenAlert() {
	// Listen on TCP port 5190
	listener, err := net.Listen("tcp", ":5193")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 5193")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("got a connection on listenAlert")
		seq := uint32(100)
		if err := oscar.ReadBos(nil, nil, conn, &seq); err != nil && err != io.EOF {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}
}

func listenOdir() {
	// Listen on TCP port 5190
	listener, err := net.Listen("tcp", ":5194")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 5194")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("got a connection on listenOdir")
		seq := uint32(100)
		if err := oscar.ReadBos(nil, nil, conn, &seq); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}
}

func handleAuthConnection(sm *oscar.SessionManager, conn net.Conn) {
	defer conn.Close()
	seq := uint32(100)
	_, err := oscar.SendAndReceiveSignonFrame(conn, &seq)
	if err != nil {
		log.Println(err)
		return
	}

	sess, err := sm.NewSession()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = oscar.ReceiveAndSendAuthChallenge(sess, conn, conn, &seq)
	if err != nil {
		log.Println(err)
		return
	}

	err = oscar.ReceiveAndSendBUCPLoginRequest(sess, conn, conn, &seq)
	if err != nil {
		log.Println(err)
		return
	}
}

func handleBOSConnection(sm *oscar.SessionManager, fm *oscar.FeedbagStore, conn net.Conn, seq *uint32) {
	fmt.Println("VerifyLogin...")
	sess, err := oscar.VerifyLogin(sm, conn, seq)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("writeOServiceHostOnline...")
	if err := oscar.WriteOServiceHostOnline(conn, seq); err != nil {
		if err == io.EOF {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if err := oscar.ReadBos(sess, fm, conn, seq); err != nil && err != io.EOF {
		if err != io.EOF {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
