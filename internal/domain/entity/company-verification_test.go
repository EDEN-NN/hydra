package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompanyVerification_TestCreateCompanyVerification(t *testing.T) {
	company, _ := CreateCompany("Empresa de teste LTDA", "13114403000103", "https://github.com.br", "GOOD")

	companyVerification := CreateCompanyVerification(company.ID.String())

	assert.NotNil(t, companyVerification)
}
