package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompanyVerification_TestCreateCompanyVerification(t *testing.T) {
	company, _ := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	companyVerification := CreateCompanyVerification(company.ID.String())

	assert.NotNil(t, companyVerification)
}
