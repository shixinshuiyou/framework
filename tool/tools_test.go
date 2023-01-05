package tool

import "testing"

func TestGenModel(t *testing.T) {
	GenModel("advertiser", "")
}

func TestGenRepository(t *testing.T) {
	GenRepository("advertiser_ext")
}
