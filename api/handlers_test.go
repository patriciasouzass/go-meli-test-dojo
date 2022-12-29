package api

import (
	"github.com/klasrak/go-meli-test-dojo/clients/swapi"
	"github.com/klasrak/go-meli-test-dojo/errors"
	"github.com/klasrak/go-meli-test-dojo/mockeable"
	"github.com/klasrak/go-meli-test-dojo/models"
	"net/http"
	"testing"
)

func TestGetStarshipHandlerBadRequest(t *testing.T) {

	url := "/api/v1/starships/:id"
	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 400

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}
}

func TestGetStarshipHandlerSuccess(t *testing.T) {

	url := "/api/v1/starships/9"

	mock := swapi.MockClient{
		GetStarshipFunc: func(id int) (models.Starship, error) {
			if id != 9 {
				t.Errorf("Assertion error. Expected: %d, Got: %d", 9, id)
			}
			return models.Starship{
				Name:                 "Death Star",
				Model:                "DS-1 Orbital Battle Station",
				Manufacturer:         "Imperial Department of Military Research, Sienar Fleet Systems",
				CostInCredits:        "1000000000000",
				Length:               "120000",
				MaxAtmospheringSpeed: "n/a",
				Crew:                 "342953",
				Passengers:           "843342",
				CargoCapacity:        "1000000000000",
				Consumables:          "3 years",
				HyperdriveRating:     "4.0",
				MGLT:                 "10",
				Class:                "Deep Space Mobile Battlestation",
				Films: []string{
					"https://swapi.dev/api/films/1/",
				},
			}, nil
		},

		GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 200
	expectedBody := `{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}`

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}

	if response.StringBody() != expectedBody {
		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
	}
}

func TestGetStarshipHandlerNotFound(t *testing.T) {

	url := "/api/v1/starships/9"

	mock := swapi.MockClient{
		GetStarshipFunc: func(id int) (models.Starship, error) {
			if id != 9 {
				t.Errorf("Assertion error. Expected: %d, Got: %d", 9, id)
			}
			return models.Starship{}, errors.NewNotFound("Not Found", "Starship not found")
		},

		GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 404

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}
}

func TestGetPeopleHandlerBadRequest(t *testing.T) {

	url := "/api/v1/people/:id"
	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 400

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}
}

func TestGetPeopleHandlerSuccess(t *testing.T) {

	url := "/api/v1/people/9"

	mock := swapi.MockClient{
		GetPeopleFunc: func(id int) (models.People, error) {
			if id != 9 {
				t.Errorf("Assertion error. Expected: %d, Got: %d", 9, id)
			}
			return models.People{
				Name:      "Death Star",
				BirthYear: "DS-1 Orbital Battle Station",
				EyeColor:  "Imperial Department of Military Research, Sienar Fleet Systems",
				Gender:    "1000000000000",
				HairColor: "120000",
				Height:    "n/a",
				Mass:      "342953",
				SkinColor: "843342",
				Homeworld: "1000000000000",
				Films: []string{
					"https://swapi.dev/api/films/1/",
				},
				Species: []string{
					"https://swapi.dev/api/films/1/",
				},
				Starships: []string{
					"https://swapi.dev/api/films/1/",
				},
			}, nil
		},

		GetPeopleFuncControl: mockeable.CallsFuncControl{ExpectedCalls: 1},
	}

	mock.Use()
	defer mockeable.CleanUpAndAssertControls(t, &mock)

	response := DoRequest(http.MethodGet, url, nil, "")
	statusCodeExpected := 200
	expectedBody := `{"name":"Death Star","birth_year":"DS-1 Orbital Battle Station","eye_color":"Imperial Department of Military Research, Sienar Fleet Systems","gender":"1000000000000","hair_color":"120000","height":"n/a","mass":"342953","skin_color":"843342","homeworld":"1000000000000","films":["https://swapi.dev/api/films/1/"],"species":["https://swapi.dev/api/films/1/"],"starships":["https://swapi.dev/api/films/1/"]}`

	if response.StatusCode != statusCodeExpected {
		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
	}

	if response.StringBody() != expectedBody {
		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
	}
}

//func TestGetPeopleListHandlerSuccess(t *testing.T) {
//
//	url := "/api/v1/people/1"
//
//	mock := swapi.MockClient{
//		GetPeopleListFunc: func() (models.PeopleList, error) {
//			return 1,
//			[]models.PeopleList{
//				Name:                 "Death Star",
//				BirthYear:                "DS-1 Orbital Battle Station",
//				EyeColor:         "Imperial Department of Military Research, Sienar Fleet Systems",
//				Gender:        "1000000000000",
//				HairColor:               "120000",
//				Height: "n/a",
//				Mass:                 "342953",
//				SkinColor:           "843342",
//				Homeworld:        "1000000000000",
//				Films: []string{
//					"https://swapi.dev/api/films/1/",
//				},
//				Species: []string{
//					"https://swapi.dev/api/films/1/",
//				},
//				Starships: []string{
//					"https://swapi.dev/api/films/1/",
//				},
//		},
//
//				GetPeopleListFuncControl: mockeable.CallsFuncControl{ExpectedCalls: },
//	}
//
//	mock.Use()
//	defer mockeable.CleanUpAndAssertControls(t, &mock)
//
//	response := DoRequest(http.MethodGet, url, nil, "")
//	statusCodeExpected := 200
//	expectedBody := `{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}`
//
//	if response.StatusCode != statusCodeExpected {
//		t.Errorf("Assertion error. Expected: %d, Got: %d", statusCodeExpected, response.StatusCode)
//	}
//
//	if response.StringBody() != expectedBody {
//		t.Errorf("Assertion error. Expected: %s, Got: %s", expectedBody, response.StringBody())
//	}
//}
