package testdata

import "errors"

type (
	PersonExtended struct {
		Name         string   `validate:"required,min=3,max=100"`
		Age          int      `validate:"gte=0,lte=130"`
		NickNameList []string `validate:"required,min=2"`

		Address         Address   `validate:"required"`
		MainCard        *Card     `validate:"required"`
		HobbyList       []Hobby   `validate:"required,min=2,dive"`
		HobbyListOfList [][]Hobby `validate:"required,min=2,dive"`
	}

	Address struct {
		Country string `validate:"required,min=3"`
		City    string `validate:"required,min=3"`
	}

	Card struct {
		Name   string `validate:"required,min=3,max=100"`
		Number string `validate:"required,min=3,max=100"`
	}

	Hobby struct {
		Name  string `validate:"required,min=3,max=100"`
		Score int    `validate:"gte=0,lte=5"`
	}
)

func (m *Address) Validate() (err error) {
	if m.Country == "Unknown" {
		err = errors.New("country cannot be 'Unknown'")
		return
	}
	return
}

func (m *Hobby) Validate() (err error) {
	if m.Name == "Everything" {
		err = errors.New("hobby name cannot be 'Everything'")
		return
	}
	return
}

func (m *Card) Validate() (err error) {
	if m.Name == "Unknown" {
		err = errors.New("card name cannot be 'Unknown'")
		return
	}
	return
}

func NewDefaultPersonExtended() PersonExtended {
	return PersonExtended{
		Name: "John Doe",
		Age:  30,

		NickNameList: []string{
			"John",
			"Doe",
		},

		Address: Address{
			Country: "USA",
			City:    "New York",
		},

		MainCard: &Card{
			Name:   "Visa",
			Number: "123456789",
		},

		HobbyList: []Hobby{
			{
				Name:  "Dancing",
				Score: 2,
			},
			{
				Name:  "Traveling",
				Score: 5,
			},
		},

		HobbyListOfList: [][]Hobby{
			{
				{
					Name:  "Dancing",
					Score: 2,
				},
				{
					Name:  "Traveling",
					Score: 5,
				},
			},
			{
				{
					Name:  "Cooking",
					Score: 3,
				},
				{
					Name:  "Reading",
					Score: 4,
				},
			},
		},
	}
}
