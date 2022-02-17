package executor

import (
	"log"
	"testing"
)

func TestHashName1(t *testing.T) {
	s1 := "SELECT 1"
	r1 := hash(s1)
	expected := "2813671660"
	if r1 != expected {
		log.Fatalf("Hash is %v, expected %v", r1, expected)
	}
}

func TestHashName2(t *testing.T) {
	s2 := "SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type"
	r2 := hash(s2)
	expected := "2720750463"
	if r2 != expected {
		log.Fatalf("Hash is %v, expected %v", r2, expected)
	}
}
