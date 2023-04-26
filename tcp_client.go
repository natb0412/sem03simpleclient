package main

import (
	"net"
	"log"
	"os"
	"github.com/natb0412/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:16")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])


	kryptertInput := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
 	_, err = conn.Write([]byte(string(kryptertInput)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	//response := string(buf[:n])
	dekrypterMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03) +5)
	//response := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03) -5)
	//log.Printf("reply from proxy: %s", dekrypterMelding)
	  log.Printf("Reply from proxy: %s", ([]byte(string(dekrypterMelding))))
}
