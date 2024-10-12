// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package planet

// Find planet with given id.
func Find(id string) *Planet {
	for _, v := range planets {
		if v.PlanetID == id {
			other := v
			return &other
		}
	}
	return nil
}

type Planet struct {
	PlanetID       string
	Name           string
	RotationPeriod uint64
	OrbitalPeriod  uint64
	Diameter       uint64
	Climate        string
	Gravity        string
	Terrain        string
	SurfaceWater   uint64
	Population     uint64
	// Relationships
	ResidentIDs []string
	FilmIDs     []string
}

var planets = []Planet{
	{
		PlanetID:       "1",
		Name:           "Tatooine",
		RotationPeriod: 23,
		OrbitalPeriod:  304,
		Diameter:       10465,
		Climate:        "arid",
		Gravity:        "1 standard",
		Terrain:        "desert",
		SurfaceWater:   1,
		Population:     200000,
		FilmIDs:        []string{"1", "3", "4", "5", "6"},
		ResidentIDs:    []string{"1", "2", "4", "6", "7", "8", "9", "11", "43", "62"},
	},
	{
		PlanetID:       "2",
		Name:           "Alderaan",
		RotationPeriod: 24,
		OrbitalPeriod:  364,
		Diameter:       12500,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "grasslands, mountains",
		SurfaceWater:   40,
		Population:     2000000000,
		FilmIDs:        []string{"1", "6"},
		ResidentIDs:    []string{"5", "68", "81"},
	},
	{
		PlanetID:       "3",
		Name:           "Yavin IV",
		RotationPeriod: 24,
		OrbitalPeriod:  4818,
		Diameter:       10200,
		Climate:        "temperate, tropical",
		Gravity:        "1 standard",
		Terrain:        "jungle, rainforests",
		SurfaceWater:   8,
		Population:     1000,
		FilmIDs:        []string{"1"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "4",
		Name:           "Hoth",
		RotationPeriod: 23,
		OrbitalPeriod:  549,
		Diameter:       7200,
		Climate:        "frozen",
		Gravity:        "1.1 standard",
		Terrain:        "tundra, ice caves, mountain ranges",
		SurfaceWater:   100,
		Population:     0,
		FilmIDs:        []string{"2"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "5",
		Name:           "Dagobah",
		RotationPeriod: 23,
		OrbitalPeriod:  341,
		Diameter:       8900,
		Climate:        "murky",
		Gravity:        "N/A",
		Terrain:        "swamp, jungles",
		SurfaceWater:   8,
		Population:     0,
		FilmIDs:        []string{"2", "3", "6"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "6",
		Name:           "Bespin",
		RotationPeriod: 12,
		OrbitalPeriod:  5110,
		Diameter:       118000,
		Climate:        "temperate",
		Gravity:        "1.5 (surface), 1 standard (Cloud City)",
		Terrain:        "gas giant",
		SurfaceWater:   0,
		Population:     6000000,
		FilmIDs:        []string{"2"},
		ResidentIDs:    []string{"26"},
	},
	{
		PlanetID:       "7",
		Name:           "Endor",
		RotationPeriod: 18,
		OrbitalPeriod:  402,
		Diameter:       4900,
		Climate:        "temperate",
		Gravity:        "0.85 standard",
		Terrain:        "forests, mountains, lakes",
		SurfaceWater:   8,
		Population:     30000000,
		FilmIDs:        []string{"3"},
		ResidentIDs:    []string{"30"},
	},
	{
		PlanetID:       "8",
		Name:           "Naboo",
		RotationPeriod: 26,
		OrbitalPeriod:  312,
		Diameter:       12120,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "grassy hills, swamps, forests, mountains",
		SurfaceWater:   12,
		Population:     4500000000,
		FilmIDs:        []string{"3", "4", "5", "6"},
		ResidentIDs:    []string{"3", "21", "35", "36", "37", "38", "39", "42", "60", "61", "66"},
	},
	{
		PlanetID:       "9",
		Name:           "Coruscant",
		RotationPeriod: 24,
		OrbitalPeriod:  368,
		Diameter:       12240,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "cityscape, mountains",
		SurfaceWater:   0,
		Population:     1000000000000,
		FilmIDs:        []string{"3", "4", "5", "6"},
		ResidentIDs:    []string{"34", "55", "74"},
	},
	{
		PlanetID:       "10",
		Name:           "Kamino",
		RotationPeriod: 27,
		OrbitalPeriod:  463,
		Diameter:       19720,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "ocean",
		SurfaceWater:   100,
		Population:     1000000000,
		FilmIDs:        []string{"5"},
		ResidentIDs:    []string{"22", "72", "73"},
	},
	{
		PlanetID:       "11",
		Name:           "Geonosis",
		RotationPeriod: 30,
		OrbitalPeriod:  256,
		Diameter:       11370,
		Climate:        "temperate, arid",
		Gravity:        "0.9 standard",
		Terrain:        "rock, desert, mountain, barren",
		SurfaceWater:   5,
		Population:     100000000000,
		FilmIDs:        []string{"5"},
		ResidentIDs:    []string{"63"},
	},
	{
		PlanetID:       "12",
		Name:           "Utapau",
		RotationPeriod: 27,
		OrbitalPeriod:  351,
		Diameter:       12900,
		Climate:        "temperate, arid, windy",
		Gravity:        "1 standard",
		Terrain:        "scrublands, savanna, canyons, sinkholes",
		SurfaceWater:   0,
		Population:     95000000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{"83"},
	},
	{
		PlanetID:       "13",
		Name:           "Mustafar",
		RotationPeriod: 36,
		OrbitalPeriod:  412,
		Diameter:       4200,
		Climate:        "hot",
		Gravity:        "1 standard",
		Terrain:        "volcanoes, lava rivers, mountains, caves",
		SurfaceWater:   0,
		Population:     20000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "14",
		Name:           "Kashyyyk",
		RotationPeriod: 26,
		OrbitalPeriod:  381,
		Diameter:       12765,
		Climate:        "tropical",
		Gravity:        "1 standard",
		Terrain:        "jungle, forests, lakes, rivers",
		SurfaceWater:   60,
		Population:     45000000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{"13", "80"},
	},
	{
		PlanetID:       "15",
		Name:           "Polis Massa",
		RotationPeriod: 24,
		OrbitalPeriod:  590,
		Diameter:       0,
		Climate:        "artificial temperate ",
		Gravity:        "0.56 standard",
		Terrain:        "airless asteroid",
		SurfaceWater:   0,
		Population:     1000000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "16",
		Name:           "Mygeeto",
		RotationPeriod: 12,
		OrbitalPeriod:  167,
		Diameter:       10088,
		Climate:        "frigid",
		Gravity:        "1 standard",
		Terrain:        "glaciers, mountains, ice canyons",
		SurfaceWater:   0,
		Population:     19000000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "17",
		Name:           "Felucia",
		RotationPeriod: 34,
		OrbitalPeriod:  231,
		Diameter:       9100,
		Climate:        "hot, humid",
		Gravity:        "0.75 standard",
		Terrain:        "fungus forests",
		SurfaceWater:   0,
		Population:     8500000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "18",
		Name:           "Cato Neimoidia",
		RotationPeriod: 25,
		OrbitalPeriod:  278,
		Diameter:       0,
		Climate:        "temperate, moist",
		Gravity:        "1 standard",
		Terrain:        "mountains, fields, forests, rock arches",
		SurfaceWater:   0,
		Population:     10000000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{"33"},
	},
	{
		PlanetID:       "19",
		Name:           "Saleucami",
		RotationPeriod: 26,
		OrbitalPeriod:  392,
		Diameter:       14920,
		Climate:        "hot",
		Gravity:        "unknown",
		Terrain:        "caves, desert, mountains, volcanoes",
		SurfaceWater:   0,
		Population:     1400000000,
		FilmIDs:        []string{"6"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "20",
		Name:           "Stewjon",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "grass",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"10"},
	},
	{
		PlanetID:       "21",
		Name:           "Eriadu",
		RotationPeriod: 24,
		OrbitalPeriod:  360,
		Diameter:       13490,
		Climate:        "polluted",
		Gravity:        "1 standard",
		Terrain:        "cityscape",
		SurfaceWater:   0,
		Population:     22000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"12"},
	},
	{
		PlanetID:       "22",
		Name:           "Corellia",
		RotationPeriod: 25,
		OrbitalPeriod:  329,
		Diameter:       11000,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "plains, urban, hills, forests",
		SurfaceWater:   70,
		Population:     3000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"14", "18"},
	},
	{
		PlanetID:       "23",
		Name:           "Rodia",
		RotationPeriod: 29,
		OrbitalPeriod:  305,
		Diameter:       7549,
		Climate:        "hot",
		Gravity:        "1 standard",
		Terrain:        "jungles, oceans, urban, swamps",
		SurfaceWater:   60,
		Population:     1300000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"15"},
	},
	{
		PlanetID:       "24",
		Name:           "Nal Hutta",
		RotationPeriod: 87,
		OrbitalPeriod:  413,
		Diameter:       12150,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "urban, oceans, swamps, bogs",
		SurfaceWater:   0,
		Population:     7000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"16"},
	},
	{
		PlanetID:       "25",
		Name:           "Dantooine",
		RotationPeriod: 25,
		OrbitalPeriod:  378,
		Diameter:       9830,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "oceans, savannas, mountains, grasslands",
		SurfaceWater:   0,
		Population:     1000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "26",
		Name:           "Bestine IV",
		RotationPeriod: 26,
		OrbitalPeriod:  680,
		Diameter:       6400,
		Climate:        "temperate",
		Gravity:        "unknown",
		Terrain:        "rocky islands, oceans",
		SurfaceWater:   98,
		Population:     62000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"19"},
	},
	{
		PlanetID:       "27",
		Name:           "Ord Mantell",
		RotationPeriod: 26,
		OrbitalPeriod:  334,
		Diameter:       14050,
		Climate:        "temperate",
		Gravity:        "1 standard",
		Terrain:        "plains, seas, mesas",
		SurfaceWater:   10,
		Population:     4000000000,
		FilmIDs:        []string{"2"},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "28",
		Name:           "unknown",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"20", "23", "29", "32", "75"},
	},
	{
		PlanetID:       "29",
		Name:           "Trandosha",
		RotationPeriod: 25,
		OrbitalPeriod:  371,
		Diameter:       0,
		Climate:        "arid",
		Gravity:        "0.62 standard",
		Terrain:        "mountains, seas, grasslands, deserts",
		SurfaceWater:   0,
		Population:     42000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"24"},
	},
	{
		PlanetID:       "30",
		Name:           "Socorro",
		RotationPeriod: 20,
		OrbitalPeriod:  326,
		Diameter:       0,
		Climate:        "arid",
		Gravity:        "1 standard",
		Terrain:        "deserts, mountains",
		SurfaceWater:   0,
		Population:     300000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"25"},
	},
	{
		PlanetID:       "31",
		Name:           "Mon Cala",
		RotationPeriod: 21,
		OrbitalPeriod:  398,
		Diameter:       11030,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "oceans, reefs, islands",
		SurfaceWater:   100,
		Population:     27000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"27"},
	},
	{
		PlanetID:       "32",
		Name:           "Chandrila",
		RotationPeriod: 20,
		OrbitalPeriod:  368,
		Diameter:       13500,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "plains, forests",
		SurfaceWater:   40,
		Population:     1200000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"28"},
	},
	{
		PlanetID:       "33",
		Name:           "Sullust",
		RotationPeriod: 20,
		OrbitalPeriod:  263,
		Diameter:       12780,
		Climate:        "superheated",
		Gravity:        "1",
		Terrain:        "mountains, volcanoes, rocky deserts",
		SurfaceWater:   5,
		Population:     18500000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"31"},
	},
	{
		PlanetID:       "34",
		Name:           "Toydaria",
		RotationPeriod: 21,
		OrbitalPeriod:  184,
		Diameter:       7900,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "swamps, lakes",
		SurfaceWater:   0,
		Population:     11000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"40"},
	},
	{
		PlanetID:       "35",
		Name:           "Malastare",
		RotationPeriod: 26,
		OrbitalPeriod:  201,
		Diameter:       18880,
		Climate:        "arid, temperate, tropical",
		Gravity:        "1.56",
		Terrain:        "swamps, deserts, jungles, mountains",
		SurfaceWater:   0,
		Population:     2000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"41"},
	},
	{
		PlanetID:       "36",
		Name:           "Dathomir",
		RotationPeriod: 24,
		OrbitalPeriod:  491,
		Diameter:       10480,
		Climate:        "temperate",
		Gravity:        "0.9",
		Terrain:        "forests, deserts, savannas",
		SurfaceWater:   0,
		Population:     5200,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"44"},
	},
	{
		PlanetID:       "37",
		Name:           "Ryloth",
		RotationPeriod: 30,
		OrbitalPeriod:  305,
		Diameter:       10600,
		Climate:        "temperate, arid, subartic",
		Gravity:        "1",
		Terrain:        "mountains, valleys, deserts, tundra",
		SurfaceWater:   5,
		Population:     1500000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"45", "46"},
	},
	{
		PlanetID:       "38",
		Name:           "Aleen Minor",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"47"},
	},
	{
		PlanetID:       "39",
		Name:           "Vulpter",
		RotationPeriod: 22,
		OrbitalPeriod:  391,
		Diameter:       14900,
		Climate:        "temperate, artic",
		Gravity:        "1",
		Terrain:        "urban, barren",
		SurfaceWater:   0,
		Population:     421000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"48"},
	},
	{
		PlanetID:       "40",
		Name:           "Troiken",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "desert, tundra, rainforests, mountains",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"49"},
	},
	{
		PlanetID:       "41",
		Name:           "Tund",
		RotationPeriod: 48,
		OrbitalPeriod:  1770,
		Diameter:       12190,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "barren, ash",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"50"},
	},
	{
		PlanetID:       "42",
		Name:           "Haruun Kal",
		RotationPeriod: 25,
		OrbitalPeriod:  383,
		Diameter:       10120,
		Climate:        "temperate",
		Gravity:        "0.98",
		Terrain:        "toxic cloudsea, plateaus, volcanoes",
		SurfaceWater:   0,
		Population:     705300,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"51"},
	},
	{
		PlanetID:       "43",
		Name:           "Cerea",
		RotationPeriod: 27,
		OrbitalPeriod:  386,
		Diameter:       0,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "verdant",
		SurfaceWater:   20,
		Population:     450000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"52"},
	},
	{
		PlanetID:       "44",
		Name:           "Glee Anselm",
		RotationPeriod: 33,
		OrbitalPeriod:  206,
		Diameter:       15600,
		Climate:        "tropical, temperate",
		Gravity:        "1",
		Terrain:        "lakes, islands, swamps, seas",
		SurfaceWater:   80,
		Population:     500000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"53"},
	},
	{
		PlanetID:       "45",
		Name:           "Iridonia",
		RotationPeriod: 29,
		OrbitalPeriod:  413,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "rocky canyons, acid pools",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"54"},
	},
	{
		PlanetID:       "46",
		Name:           "Tholoth",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{},
	},
	{
		PlanetID:       "47",
		Name:           "Iktotch",
		RotationPeriod: 22,
		OrbitalPeriod:  481,
		Diameter:       0,
		Climate:        "arid, rocky, windy",
		Gravity:        "1",
		Terrain:        "rocky",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"56"},
	},
	{
		PlanetID:       "48",
		Name:           "Quermia",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"57"},
	},
	{
		PlanetID:       "49",
		Name:           "Dorin",
		RotationPeriod: 22,
		OrbitalPeriod:  409,
		Diameter:       13400,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"58"},
	},
	{
		PlanetID:       "50",
		Name:           "Champala",
		RotationPeriod: 27,
		OrbitalPeriod:  318,
		Diameter:       0,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "oceans, rainforests, plateaus",
		SurfaceWater:   0,
		Population:     3500000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"59"},
	},
	{
		PlanetID:       "51",
		Name:           "Mirial",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "deserts",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"64", "65"},
	},
	{
		PlanetID:       "52",
		Name:           "Serenno",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "rainforests, rivers, mountains",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"67"},
	},
	{
		PlanetID:       "53",
		Name:           "Concord Dawn",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "jungles, forests, deserts",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"69"},
	},
	{
		PlanetID:       "54",
		Name:           "Zolan",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"70"},
	},
	{
		PlanetID:       "55",
		Name:           "Ojom",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "frigid",
		Gravity:        "unknown",
		Terrain:        "oceans, glaciers",
		SurfaceWater:   100,
		Population:     500000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"71"},
	},
	{
		PlanetID:       "56",
		Name:           "Skako",
		RotationPeriod: 27,
		OrbitalPeriod:  384,
		Diameter:       0,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "urban, vines",
		SurfaceWater:   0,
		Population:     500000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"76"},
	},
	{
		PlanetID:       "57",
		Name:           "Muunilinst",
		RotationPeriod: 28,
		OrbitalPeriod:  412,
		Diameter:       13800,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "plains, forests, hills, mountains",
		SurfaceWater:   25,
		Population:     5000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"77"},
	},
	{
		PlanetID:       "58",
		Name:           "Shili",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "temperate",
		Gravity:        "1",
		Terrain:        "cities, savannahs, seas, plains",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"78"},
	},
	{
		PlanetID:       "59",
		Name:           "Kalee",
		RotationPeriod: 23,
		OrbitalPeriod:  378,
		Diameter:       13850,
		Climate:        "arid, temperate, tropical",
		Gravity:        "1",
		Terrain:        "rainforests, cliffs, canyons, seas",
		SurfaceWater:   0,
		Population:     4000000000,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"79"},
	},
	{
		PlanetID:       "60",
		Name:           "Umbara",
		RotationPeriod: 0,
		OrbitalPeriod:  0,
		Diameter:       0,
		Climate:        "unknown",
		Gravity:        "unknown",
		Terrain:        "unknown",
		SurfaceWater:   0,
		Population:     0,
		FilmIDs:        []string{},
		ResidentIDs:    []string{"82"},
	},
}