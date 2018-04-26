package main

import (
	"strings"
	"testing"

	pb "github.com/0990/golearn/protobuf/pb"
	"github.com/golang/protobuf/proto"
)

func TestPromptForAddressReturnsAddress(t *testing.T) {
	in := `123456 Example Name name@example.com123-456 home222-222 mobile777 unknown`
	got, err := promptForAddress(strings.NewReader(in))

	if err != nil {
		t.Fatalf("promptForAddress(%q) had unexpected error:%s", in, err.Error())
	}

	if got.Id != 12345 {
		t.Errorf("promptForAddress(%q) got %d ,wantID:%d", in, got.Id, 12345)
	}

	if got.Name != "Example Name" {
		t.Errorf("promptForAddress(%q) => want name %q, got %q", in, "Example Name", got.Name)
	}
	if got.Email != "name@example.com" {
		t.Errorf("promptForAddress(%q) => want email %q, got %q", in, "name@example.com", got.Email)
	}

	want := []*pb.Person_PhoneNumber{
		{Number: "123-456", Type: pb.Person_HOME},
		{Number: "222-222", Type: pb.Person_MOBILE},
		{Number: "777", Type: pb.Person_MOBILE},
	}
	if len(got.Phones) != len(want) {
		t.Errorf("want %d phone numbers, got %d", len(want), len(got.Phones))
	}
	phones := len(got.Phones)
	if phones > len(want) {
		phones = len(want)
	}
	for i := 0; i < phones; i++ {
		if !proto.Equal(got.Phones[i], want[i]) {
			t.Errorf("want phone %q, got %q", *want[i], *got.Phones[i])
		}

	}
}