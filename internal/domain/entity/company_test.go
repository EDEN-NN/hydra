package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 13.114.403/0001-03
func TestCreateCompany(t *testing.T) {
	tests := []struct {
		name           string
		registryNumber string
		reputation     Reputation
		urlSite        string
		wantErr        bool
	}{
		{"Empresa de teste LTDA", "13114403000103", "https://github.com.br", "GOOD", false},
		{"Empr", "13114403000103", "http://ifood.com.br", "NOT_RECOMMENDED", true},
		{"Empresa de teste LTDA", "13103000103", "http://uber.com.br", "REGULAR", true},
	}

	for _, tt := range tests {
		_, err := CreateCompany(tt.name, tt.registryNumber, tt.urlSite, tt.reputation)
		if (err != nil) != tt.wantErr {
			t.Errorf("CreateCompany(%v), err: %v, wantErr: %v", tt, err, tt.wantErr)
		}
	}
}

func TestCompany_ChangeName(t *testing.T) {
	company, err := CreateCompany("Empresa de teste LTDA", "13114403000103", "https://github.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Empresa de Sociedade Anonima LTDA", false},
		{"Emp", true},
		{"", true},
	}

	for _, tt := range tests {
		companyOldName := company.Name
		err := company.ChangeName(tt.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("ChangeName(%v), err: %v, wantErr: %v", tt, err, tt.wantErr)
		}
		assert.NotEqual(t, companyOldName, company.Name)
	}

}

func TestCompany_ActiveProduct(t *testing.T) {
	company, _ := CreateCompany("Empresa de teste LTDA", "13114403000103", "https://github.com.br", "GOOD")

	oldProductStatus := company.HasProduct

	company.ActiveProduct()
	assert.False(t, oldProductStatus)
	assert.NotEqual(t, oldProductStatus, company.HasProduct)
	assert.True(t, company.HasProduct)
}

func TestCompany_ChangeRegistryNumber(t *testing.T) {
	company, err := CreateCompany("Empresa de teste LTDA", "13114403000103", "https://github.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	tests := []struct {
		registryNumber string
		wantErr        bool
	}{
		{"13114403000104", false},
		{"1311440304", true},
		{"131144030001042432321213213", true},
	}

	for _, tt := range tests {
		oldRegistryNumber := company.RegistryNumber
		err := company.ChangeRegistryNumber(tt.registryNumber)
		if (err != nil) != tt.wantErr {
			t.Errorf("ChangeRegistryNumber(%v), err: %v, wantErr: %v", tt, err, tt.wantErr)
		}
		assert.NotEqual(t, oldRegistryNumber, company.RegistryNumber)
	}
}

func TestCompany_UpdatedBiometryStatus(t *testing.T) {
	company, err := CreateCompany("Empresa de teste LTDA", "13114403000103", "https://github.com.br", "GOOD")

	assert.Nil(t, err)
	assert.NotNil(t, company)

	oldBiometryStatus := company.Biometry

	company.UpdatedBiometryStatus(ACCEPTED)

	assert.NotEqual(t, oldBiometryStatus, company.Biometry)
}
