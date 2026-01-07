package test

import (
	"project/validator"
	"project/entity"
	"testing"

	. "github.com/onsi/gomega"
)

func testemployeecode(t *testing.T) {
	g := NewGomegaWithT(t)

	e := entity.Employees {
		Name: "tawan",
		Salary: 20000,
		EmployeeCode: "tv-5741",
	}

	err := validator.ValidateStruct(&e)

	g.Expect(err).To(BeNil())
	g.Expect(err.Error()).To(Equal("not matches(^[HR]-{4}$)"))

}