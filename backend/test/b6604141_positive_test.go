package test

import (
	"project/validator"
	"project/entity"
	"testing"

	. "github.com/onsi/gomega"
)

func fullvalid(t *testing.T) {
	g := NewGomegaWithT(t)

	e := entity.Employees {
		Name: "tawan",
		Salary: 20000,
		EmployeeCode: "HR-5741",
	}

	err := validator.ValidateStruct(&e)

	g.Expect(err).NotTo(BeNil())

}