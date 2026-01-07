package test

import (
	"project/validator"
	"project/entity"
	"testing"

	. "github.com/onsi/gomega"
)

func testsalary(t *testing.T) {
	g := NewGomegaWithT(t)

	e := entity.Employees {
		Name: "tawan",
		Salary: 1000,
		EmployeeCode: "HR-5741",
	}

	err := validator.ValidateStruct(&e)

	g.Expect(err).To(BeNil())
	g.Expect(err.Error()).To(Equal("must be 15000-200000"))

}