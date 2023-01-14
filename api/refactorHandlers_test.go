package api

import (
	"fmt"
	"github.com/klasrak/go-meli-test-dojo/clients/swapi"
	"github.com/klasrak/go-meli-test-dojo/errors"
	"github.com/klasrak/go-meli-test-dojo/mockeable"
	"github.com/klasrak/go-meli-test-dojo/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetStarshipHandler(t *testing.T) {

	type args struct {
		id                          interface{}
		expectedResponseBody        string
		expectedStatusCode          int
		expectedMockSuccessResponse models.Starship
		expectedMockErrorResponse   error
		expectedMockCountCalls      int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Bad Request",
			args: args{
				id:                          "invalid_id",
				expectedResponseBody:        `{"type": "BAD_REQUEST", "message": "Bad request. Reason: invalid id"}`,
				expectedStatusCode:          http.StatusBadRequest,
				expectedMockSuccessResponse: models.Starship{},
				expectedMockErrorResponse:   errors.NewBadRequest("invalid id"),
				expectedMockCountCalls:      0,
			},
		},
		{
			name: "Not Found",
			args: args{
				id:                          9,
				expectedResponseBody:        `{"type": "NOT_FOUND", "message": "resource: starship with id: 9 not found"}`,
				expectedStatusCode:          http.StatusNotFound,
				expectedMockSuccessResponse: models.Starship{},
				expectedMockErrorResponse:   errors.NewNotFound("starship", "9"),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				id:                          1,
				expectedResponseBody:        `{"type": "INTERNAL_SERVER_ERROR", "message": "Internal server error."}`,
				expectedStatusCode:          http.StatusInternalServerError,
				expectedMockSuccessResponse: models.Starship{},
				expectedMockErrorResponse:   errors.NewInternal(),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Success",
			args: args{
				id:                   2,
				expectedResponseBody: `{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}`,
				expectedStatusCode:   http.StatusOK,
				expectedMockSuccessResponse: models.Starship{
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
				},
				expectedMockCountCalls: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			swapiMock := swapi.MockClient{

				GetStarshipFunc: func(id int) (models.Starship, error) {
					assert.Equal(t, testCase.args.id, id)
					return testCase.args.expectedMockSuccessResponse, testCase.args.expectedMockErrorResponse
				},

				GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: testCase.args.expectedMockCountCalls},
			}

			swapiMock.Use()
			defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

			handlerURL := fmt.Sprintf("/api/v1/starships/%v", testCase.args.id)

			response := DoRequest(http.MethodGet, handlerURL, nil, "")

			assert.Equal(t, testCase.args.expectedStatusCode, response.StatusCode)
			assert.JSONEq(t, testCase.args.expectedResponseBody, response.StringBody())

		})
	}
}

func TestGetStarshipsHandler(t *testing.T) {

	type args struct {
		expectedResponseBody        string
		expectedStatusCode          int
		expectedMockSuccessResponse models.Starships
		expectedMockErrorResponse   error
		expectedMockCountCalls      int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Not Found",
			args: args{
				expectedResponseBody:        `{"type": "NOT_FOUND", "message": "resource: starships not found"}`,
				expectedStatusCode:          http.StatusNotFound,
				expectedMockSuccessResponse: models.Starships{},
				expectedMockErrorResponse:   errors.NewNotFound("starships", ""),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				expectedResponseBody:        `{"type": "INTERNAL_SERVER_ERROR", "message": "Internal server error."}`,
				expectedStatusCode:          http.StatusInternalServerError,
				expectedMockSuccessResponse: models.Starships{},
				expectedMockErrorResponse:   errors.NewInternal(),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Success",
			args: args{
				expectedResponseBody: `{"count":2,"results":[{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null},{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}]}`,
				expectedStatusCode:   http.StatusOK,
				expectedMockSuccessResponse: models.Starships{
					Count: 2,
					Results: []models.Starship{
						{
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
						},
						{
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
						},
					},
				},

				expectedMockCountCalls: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			swapiMock := swapi.MockClient{

				GetStarshipsFunc: func() (models.Starships, error) {
					return testCase.args.expectedMockSuccessResponse, testCase.args.expectedMockErrorResponse
				},

				GetStarshipsFuncControl: mockeable.CallsFuncControl{ExpectedCalls: testCase.args.expectedMockCountCalls},
			}

			swapiMock.Use()
			defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

			handlerURL := "/api/v1/starships"

			response := DoRequest(http.MethodGet, handlerURL, nil, "")

			assert.Equal(t, testCase.args.expectedStatusCode, response.StatusCode)
			assert.JSONEq(t, testCase.args.expectedResponseBody, response.StringBody())

		})
	}
}

func TestGetPeopleHandler(t *testing.T) {

	type args struct {
		id                          interface{}
		expectedResponseBody        string
		expectedStatusCode          int
		expectedMockSuccessResponse models.People
		expectedMockErrorResponse   error
		expectedMockCountCalls      int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Bad Request",
			args: args{
				id:                          "invalid_id",
				expectedResponseBody:        `{"type": "BAD_REQUEST", "message": "Bad request. Reason: invalid id"}`,
				expectedStatusCode:          http.StatusBadRequest,
				expectedMockSuccessResponse: models.People{},
				expectedMockErrorResponse:   errors.NewBadRequest("invalid id"),
				expectedMockCountCalls:      0,
			},
		},
		{
			name: "Not Found",
			args: args{
				id:                          9,
				expectedResponseBody:        `{"type": "NOT_FOUND", "message": "resource: people with id: 9 not found"}`,
				expectedStatusCode:          http.StatusNotFound,
				expectedMockSuccessResponse: models.People{},
				expectedMockErrorResponse:   errors.NewNotFound("people", "9"),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				id:                          1,
				expectedResponseBody:        `{"type": "INTERNAL_SERVER_ERROR", "message": "Internal server error."}`,
				expectedStatusCode:          http.StatusInternalServerError,
				expectedMockSuccessResponse: models.People{},
				expectedMockErrorResponse:   errors.NewInternal(),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Success",
			args: args{
				id:                   2,
				expectedResponseBody: `{"name":"Luiza","birth_year":"01/03/1962","eye_color":"Marrom","gender":"Feminino","hair_color":"Castanho escuro","height":"1,67","mass":"57Kg","skin_color":"Marrom","homeworld":"1000000000000","films":["https://swapi.dev/api/films/1/"],"species":["https://swapi.dev/api/films/1/"],"starships":["https://swapi.dev/api/films/1/"]}`,
				expectedStatusCode:   http.StatusOK,
				expectedMockSuccessResponse: models.People{
					Name:      "Luiza",
					BirthYear: "01/03/1962",
					EyeColor:  "Marrom",
					Gender:    "Feminino",
					HairColor: "Castanho escuro",
					Height:    "1,67",
					Mass:      "57Kg",
					SkinColor: "Marrom",
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
				},
				expectedMockCountCalls: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			swapiMock := swapi.MockClient{

				GetPeopleFunc: func(id int) (models.People, error) {
					assert.Equal(t, testCase.args.id, id)
					return testCase.args.expectedMockSuccessResponse, testCase.args.expectedMockErrorResponse
				},

				GetPeopleFuncControl: mockeable.CallsFuncControl{ExpectedCalls: testCase.args.expectedMockCountCalls},
			}

			swapiMock.Use()
			defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

			handlerURL := fmt.Sprintf("/api/v1/people/%v", testCase.args.id)

			response := DoRequest(http.MethodGet, handlerURL, nil, "")

			assert.Equal(t, testCase.args.expectedStatusCode, response.StatusCode)
			assert.JSONEq(t, testCase.args.expectedResponseBody, response.StringBody())

		})
	}
}

func TestGetPeopleListHandler(t *testing.T) {

	type args struct {
		expectedResponseBody        string
		expectedStatusCode          int
		expectedMockSuccessResponse models.PeopleList
		expectedMockErrorResponse   error
		expectedMockCountCalls      int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Not Found",
			args: args{
				expectedResponseBody:        `{"type": "NOT_FOUND", "message": "resource: peoples not found"}`,
				expectedStatusCode:          http.StatusNotFound,
				expectedMockSuccessResponse: models.PeopleList{},
				expectedMockErrorResponse:   errors.NewNotFound("peoples", ""),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				expectedResponseBody:        `{"type": "INTERNAL_SERVER_ERROR", "message": "Internal server error."}`,
				expectedStatusCode:          http.StatusInternalServerError,
				expectedMockSuccessResponse: models.PeopleList{},
				expectedMockErrorResponse:   errors.NewInternal(),
				expectedMockCountCalls:      1,
			},
		},
		{
			name: "Success",
			args: args{
				expectedResponseBody: `{"count":2,"results":[{"name":"Luiza","birth_year":"01/03/1962","eye_color":"Marrom","gender":"Feminino","hair_color":"Castanho escuro","height":"1,67","mass":"57Kg","skin_color":"Marrom","homeworld":"1000000000000","films":["https://swapi.dev/api/films/1/"],"species":["https://swapi.dev/api/films/1/"],"starships":["https://swapi.dev/api/films/1/"]},{"name":"Luiza","birth_year":"01/03/1962","eye_color":"Marrom","gender":"Feminino","hair_color":"Castanho escuro","height":"1,67","mass":"57Kg","skin_color":"Marrom","homeworld":"1000000000000","films":["https://swapi.dev/api/films/1/"],"species":["https://swapi.dev/api/films/1/"],"starships":["https://swapi.dev/api/films/1/"]}]}`,
				expectedStatusCode:   http.StatusOK,
				expectedMockSuccessResponse: models.PeopleList{
					Count: 2,
					Results: []models.People{
						{
							Name:      "Luiza",
							BirthYear: "01/03/1962",
							EyeColor:  "Marrom",
							Gender:    "Feminino",
							HairColor: "Castanho escuro",
							Height:    "1,67",
							Mass:      "57Kg",
							SkinColor: "Marrom",
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
						},
						{
							Name:      "Luiza",
							BirthYear: "01/03/1962",
							EyeColor:  "Marrom",
							Gender:    "Feminino",
							HairColor: "Castanho escuro",
							Height:    "1,67",
							Mass:      "57Kg",
							SkinColor: "Marrom",
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
						},
					},
				},

				expectedMockCountCalls: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			swapiMock := swapi.MockClient{

				GetPeopleListFunc: func() (models.PeopleList, error) {
					return testCase.args.expectedMockSuccessResponse, testCase.args.expectedMockErrorResponse
				},

				GetPeopleListFuncControl: mockeable.CallsFuncControl{ExpectedCalls: testCase.args.expectedMockCountCalls},
			}

			swapiMock.Use()
			defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

			handlerURL := "/api/v1/people"

			response := DoRequest(http.MethodGet, handlerURL, nil, "")

			assert.Equal(t, testCase.args.expectedStatusCode, response.StatusCode)
			assert.JSONEq(t, testCase.args.expectedResponseBody, response.StringBody())

		})
	}
}
