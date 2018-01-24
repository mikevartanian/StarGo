/* TODO:
    Work on error codes
    Add in ping, stats, and reset stats functionality
    Add in sytax checking for invalid characters during compression
*/



package main

import (
     "fmt"   
	"net"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"bytes"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)


type Header struct{
	magic uint32
	payloadLength uint16
	code uint16
}

type Stats struct{
	bytesSent uint32
	bytesReceived uint32
	ratio byte
}
	

var totalBytesReceived int = 0
var totalBytesSent int = 0

func main(){
	l, _ := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	defer l.Close()

	for {
		conn, _:= l.Accept()
	
		go handleRequest(conn)
	}
}


func decodeHeader(buf []byte) (Header) { 
  var h Header
  s := hex.EncodeToString(buf)
  magicHex := s[0:8]
  payloadLengthHex := s[8:12]
  codeHex := s[12:]

  magicDecode , _ := hex.DecodeString(magicHex)
  payloadLengthDecode , _  := hex.DecodeString(payloadLengthHex)
  codeDecode , _ := hex.DecodeString(codeHex)

  h.magic = binary.BigEndian.Uint32(magicDecode)
  h.payloadLength = binary.BigEndian.Uint16(payloadLengthDecode)
  h.code = binary.BigEndian.Uint16(codeDecode)

  fmt.Println("Magic:" ,h.magic, " Payload Length:", h.payloadLength, " Code:", h.code)
  return h 
}




func handleRequest(conn net.Conn){
  buf := make([]byte, 8)
  // Read the incoming connection into the buffer.
  conn.Read(buf)
  h := decodeHeader(buf)

  //requestCode := h.code


  totalBytesReceived += len(buf)

  fmt.Println(totalBytesReceived)

  var code int
  var messageRet []byte
  payload := make([]byte , h.payloadLength)
if h.code == 1 && h.payloadLength == 0{
		code , payload = ping()
	} else if h.code == 2 && h.payloadLength == 0{
		code , payload = getStats()
	} else if h.code == 3 && h.payloadLength == 0{
		code , payload = resetStats()
	} else if h.code == 4{
		if h.payloadLength >= 4096{
			code, payload = 2 , []byte("Message Too Large")
		}else{
			conn.Read(payload)
			totalBytesReceived += int(h.payloadLength)
		code , payload = compress(string(payload), int(h.payloadLength))	
		}
	} else {
		code , messageRet = error()
	}
	fmt.Println(messageRet)
  h.code = uint16(code)
  var bin_buf bytes.Buffer
  binary.Write(&bin_buf , binary.BigEndian , h)
  conn.Write(bin_buf.Bytes())
  totalBytesSent += len(bin_buf.Bytes())
  conn.Write(payload)
  totalBytesSent += len(payload)
  fmt.Println(totalBytesSent)
  conn.Close()
}

func getStats() (int, []byte){
    var ratioFloat float64
	if totalBytesSent == 0 { //To avoid divide by zero error
		ratioFloat = 0.0
	} else{
		ratioFloat = float64(totalBytesSent) / float64(totalBytesReceived)
	}

	ratioFloat *= 100
	ratio := int(ratioFloat)
	s := Stats{uint32(totalBytesSent), uint32(totalBytesReceived), byte(ratio)}
	var stats_buf bytes.Buffer
	binary.Write(&stats_buf , binary.BigEndian , s)
	return 0 , stats_buf.Bytes()
}

func resetStats() (int, []byte){
	totalBytesReceived = 0
	totalBytesSent = 0

	return 0 , []byte("Stats have been reset")
}

func ping() (int, []byte){
	return 0 , []byte("Service is operating")
}

func checkForSyntaxError(s string) (bool , bool, bool){
	alphaError := false
	numError := false
	unknownError := false

	for _ , val := range s{
		if val >= 65 && val <=90{
			alphaError = true
			break
		} else if val >=48 && val <= 57{
			numError = true
			break
		} else if val < 97 || val > 123{
			unknownError = true
			break
		}
	}

		return alphaError , numError , unknownError

}
func compress(payload string , payloadLength int) (int, []byte){
	 letterError , numError , unknownError	:= checkForSyntaxError(payload)
	//alphaError, numError := checkForSyntaxError(payload)
	if letterError	{
		return 4 , []byte("<invalid: contains uppercase characters>")
	}

	if numError{
		return 5 , []byte("<invalid: contains numbers>")
	}

	if unknownError	{
		return 6 , []byte("An unknown character was encountered")
	}
	count := 1
	s := ""

	if payloadLength == 1 || payloadLength == 2{
		return 0 , []byte(payload)
	}
	for i := 0; i < len(payload) - 1 ; i++{

		if payload[i] == payload[i + 1] {
			count++
		} else{
			if count == 1{
				s += string(payload[i])
			} else if count == 2{
				s += string(payload[i])
				s += string(payload[i])
			} else{
				s += strconv.Itoa(count)
				s += string(payload[i])
			}
			count = 1
		}
	}


    lastIndex := len(payload) - 1


	if count == 1{
				s += string(payload[lastIndex])
			} else if count == 2{
				s += string(payload[lastIndex])
				s += string(payload[lastIndex])
			} else{
				s += strconv.Itoa(count)
				s += string(payload[lastIndex])
			}
	return 0,[]byte(s)
}

func error() (int, []byte){
	return 1 , []byte("Error")
}