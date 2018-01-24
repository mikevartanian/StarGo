package main

import "fmt"
import "net"
import "encoding/binary"
//import "encoding/hex"

func main(){
	
/*	var magic uint32 = 1398035033
	var payloadLength uint16 = 2
	var code uint16 = 2

	header := 13980350330201


*/

	
    //headerCompress := 6004514745397542916

    //headerPing := 6004514745397280769


   headerGetStats := 6004514745397280770

    //headerResetStats := 6004514745397280771
	conn, _ := net.Dial("tcp","127.0.0.1:4000")

	/*arr = append(arr, magicArr)
	arr = append(arr , payloadArr)
	arr = append(arr, codeArr)
*/

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf,uint64(headerGetStats))


	conn.Write(buf)


	buf2 := []byte("aaaa")


	conn.Write(buf2)
	ret := make([]byte, 8)
    conn.Read(ret)

    payload := make([]byte , 9)

    conn.Read(payload)

    fmt.Println(payload)
	fmt.Println(ret)
}
	