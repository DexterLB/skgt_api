package backend

import (
	"testing"

	"github.com/DexterLB/skgt_api/common"
)

func TestBackend_Transports(t *testing.T) {
	backend := fillDatabase(t)
	defer closeBackend(t, backend)

	transports, err := backend.Transports()
	if err != nil {
		t.Fatal(err)
	}

	expected := []*common.Line{
		&common.Line{
			Vehicle: common.Bus,
			Number:  "94",
		},
		&common.Line{
			Vehicle: common.Tram,
			Number:  "10",
		},
	}

	assertEqualJSON(expected, transports, t)
}

func TestBackend_Routes(t *testing.T) {
	backend := fillDatabase(t)
	defer closeBackend(t, backend)

	routes, err := backend.Routes("94", common.Bus)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*common.Route{
		&common.Route{
			Direction: "A - B",
			Stops: []*common.Stop{
				&common.Stop{
					ID:          4,
					Name:        "qux",
					Description: "Qux",
					Latitude:    42,
					Longitude:  26,
				},
				&common.Stop{
					ID:          5,
					Name:        "quux",
					Description: "Quux",
					Latitude:    42,
					Longitude:  26,
				},
				&common.Stop{
					ID:          6,
					Name:        "corge",
					Description: "Corge",
					Latitude:    42,
					Longitude:  26,
				},
			},
		},
		&common.Route{
			Direction: "B - A",
			Stops: []*common.Stop{
				&common.Stop{
					ID:          7,
					Name:        "garply",
					Description: "Garply",
					Latitude:    42,
					Longitude:  26,
				},
				&common.Stop{
					ID:          8,
					Name:        "waldo",
					Description: "Waldo",
					Latitude:    42,
					Longitude:  26,
				},
				&common.Stop{
					ID:          9,
					Name:        "fred",
					Description: "Fred",
					Latitude:    42,
					Longitude:  26,
				},
			},
		},
	}

	assertEqualJSON(expected, routes, t)
}
