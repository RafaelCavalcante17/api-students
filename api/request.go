package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"Active"`
}

func errParamRequerid(param, typ string) error {
	return fmt.Errorf("param '%s' of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequerid("name", "string")
	}
	if s.Email == "" {
		return errParamRequerid("email", "string")
	}
	if s.CPF == 0 {
		return errParamRequerid("cpf", "string")
	}
	if s.Age == 0 {
		return errParamRequerid("age", "string")
	}
	if s.Active == nil {
		return errParamRequerid("active", "bool")
	}
	return nil

}
