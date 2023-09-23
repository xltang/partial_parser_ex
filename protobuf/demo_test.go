package protobuf

import (
	"fmt"
	"testing"

	. "github.com/easierway/partialparse/protobuf_def"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PrepareTestData() ([]byte, error) {
	now := timestamppb.Now()
	p := Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*Person_PhoneNumber{
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
			{
				Number: "555-4321",
				Type:   Person_PHONE_TYPE_HOME,
			},
		},
		T1: "t1",
		T2: "t2",
		T3: "t3",
		T4: "t3",
		T5: "t3",
		T6: "t3",
		T7: "t3",
		T8: "t3",
		T9: "t3",
		T10: "t3",
		T11: "t3",
		T12: "t3",
		T13: "t3",
		T14: "t3",
		T15: "t3",
		T16: "t3",
		T17: "t3",
		T18: "t3",
		T19: "t3",
		T20: "t3",
		T21: "t3",
		T22: "t3",
		T23: "t3",
		T24: "t3",
		T25: "t3",
		T26: "t3",
		T27: "t3",
		T28: "t3",
		T29: "t3",
		T30: "t3",
		T31: "t3",
		LastUpdated: now,
	}
	return proto.Marshal(&p)
}

func TestFullyParse(t *testing.T) {
	data, err := PrepareTestData()
	if err != nil {
		t.Fatal(err)
	}
	p1 := &Person{}
	proto.Unmarshal(data, p1)
	fmt.Println(p1)
}

func TestPartialParse(t *testing.T) {
	data, err := PrepareTestData()
	if err != nil {
		t.Fatal(err)
	}
	p2 := &PartialPerson{}
	proto.Unmarshal(data, p2)
	fmt.Println(p2)
}

func BenchmarkFullyParse(b *testing.B) {
	data, err := PrepareTestData()
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		p1 := &Person{}
		proto.Unmarshal(data, p1)
	}
}

func BenchmarkPartialParse(b *testing.B) {
	data, err := PrepareTestData()

	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		p2 := &PartialPerson{}
		proto.Unmarshal(data, p2)
	}
}
