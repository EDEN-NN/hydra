package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 13.114.403/0001-03

func TestCreateCompany(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)
}

func TestCompany_ChangeName(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	companyOldName := company.Name

	company.ChangeName("Zalíssimo Gadders dos Santos LTDA")

	assert.NotEqual(t, companyOldName, company.Name)
}

func TestCompany_ChangeNameFailWhenHasLessThan5Characters(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	_, errs := company.ChangeName("Zalí")

	assert.NotNil(t, errs)
}

func TestCompany_ActiveProduct(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	oldProductStatus := company.HasRav

	company.ActiveProduct()

	assert.NotEqual(t, oldProductStatus, company.HasRav)
}

func TestCompany_ChangeRegistryNumber(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	oldRegistryNumber := company.RegistryNumber

	company.ChangeRegistryNumber("13114403000104")

	assert.NotEqual(t, oldRegistryNumber, company.RegistryNumber)
}

func TestCompany_ChangeRegistryNumberFailWhenRegistryNumberHasLessThan14Characters(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	_, errs := company.ChangeRegistryNumber("1170707070")

	assert.NotNil(t, errs)
}

func TestCompany_ChangeRegistryNumberFailWhenRegistryNumberHasMoreThan14Characters(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	_, errs := company.ChangeRegistryNumber("131144030001031032746343")

	assert.NotNil(t, errs)
}

func TestCompany_UpdatedBiometryStatus(t *testing.T) {
	company, err := CreateCompany("Zaludo Gadelhudo LTDA", "13114403000103", "http://reclameaqui.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	oldBiometryStatus := company.Biometry

	company.UpdatedBiometryStatus(ACCEPTED)

	assert.NotEqual(t, oldBiometryStatus, company.Biometry)
}
