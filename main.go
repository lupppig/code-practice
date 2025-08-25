package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	inputVal, err := readLine(reader)
	checkOnErr(err)

	intCount, err := strconv.ParseInt(inputVal, 10, 32)
	checkOnErr(err)

	llist := &SinglyLinkedList{}
	for i := 0; i < int(intCount); i++ {
		inputVal, err := readLine(reader)
		checkOnErr(err)
		intCount, err := strconv.ParseInt(inputVal, 10, 64)
		checkOnErr(err)
		llist.insertNodeAtHead(int32(intCount))
	}
	printLinkedList(llist.head)

}

func readLine(reader *bufio.Reader) (string, error) {
	str, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(str), "\r\n"), nil
}

func checkOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
