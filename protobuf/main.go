package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	protob "github.com/0990/golearn/protobuf/pb"
	"github.com/golang/protobuf/proto"
)

func promptForAddress(r io.Reader) (*protob.Person, error) {
	p := &protob.Person{}

	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID number:")
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}

	p.Name = strings.TrimSpace(name)
	fmt.Print("Enter email addrss(black for none):")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}

	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number(or leave blank to finish):")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		pn := &protob.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile,hone,or work phone?")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		ptype = strings.TrimSpace(ptype)

		switch ptype {
		case "mobile":
			pn.Type = protob.Person_MOBILE
		case "home":
			pn.Type = protob.Person_HOME
		case "work":
			pn.Type = protob.Person_WORK
		default:
			fmt.Printf("Unkonw type:%q,USing default.\n", ptype)
		}
		p.Phones = append(p.Phones, pn)
	}
	return p, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage:%s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s:File not found.Createing new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	book := &protob.AddressBook{}

	fmt.Println("AddressBook protoMessageName:",proto.MessageName(book))

	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(book)

	addr, err := promptForAddress(os.Stdin)

	if err != nil {
		log.Fatalln("error with addrsss:", err)
	}
	book.People = append(book.People, addr)

	out, err := proto.Marshal(book)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
