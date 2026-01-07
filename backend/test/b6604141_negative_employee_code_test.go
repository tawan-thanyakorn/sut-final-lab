package test

import (
	"project/entity"
	"project/validator"
	"testing"

	. "github.com/onsi/gomega"
)

func testemployeecode(t *testing.T) {
	g := NewGomegaWithT(t)

	e := entity.Employees{
		Name:         "tawan",
		Salary:       20000,
		EmployeeCode: "tv-5741",
	}

	err := validator.ValidateStruct(&e)

	g.Expect(err).To(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))

}
