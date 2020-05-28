package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
)

//TODO zabbix添加zabbix错误处理
func ZabbixGet(server string, port int, key string, timeout int) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	dataLen := make([]byte, 8)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(key)))
	buffer := append([]byte("ZBXD\x01"), dataLen...)
	buffer = append(buffer, []byte(key)...)
	fmt.Println(string(buffer))
	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}
	return result[13:], err
}

func main() {
	result, err := ZabbixGet("127.0.0.1", 10050, "agent.ping", 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(string(result))
}
