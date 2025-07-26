package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ChannelType string
type Result string

const (
	IA_DETECTOR        ChannelType = "IA_DETECTOR"
	WEBSITE_DETECTOR   ChannelType = "WEBSITE_DETECTOR"
	COMPLAINS_ANSWERED ChannelType = "COMPLAINS_ANSWERED"

	APPROVED    Result = "APPROVED"
	DISAPPROVED Result = "DISAPPROVED"
	IN_PROGRESS Result = "IN_PROGRESS"
)

type Channel struct {
	Attempts  []string    `json:"attempts" bson:"attempts"`
	Type      ChannelType `json:"type" bson:"type"`
	Result    Result      `json:"result" bson:"result"`
	CreatedAt time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type PreAudit struct {
	HasBiometry   bool `json:"hasBiometry" bson:"hasBiometry"`
	HasSite       bool `json:"hasSite" bson:"hasSite"`
	HasReputation bool `json:"hasReputation" bson:"hasReputation"`
}

type CompanyVerification struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	CorrelationId string             `json:"correlationId" bson:"correlationId"`
	Channels      []Channel          `json:"channels" bson:"channels"`
	PreAudit      PreAudit           `json:"preAudit" bson:"preAudit"`
	Result        Result             `json:"result" bson:"result"`
}

func CreateCompanyVerification(correlationId string) *CompanyVerification {

	preAudit := &PreAudit{
		HasBiometry:   true,
		HasSite:       true,
		HasReputation: true,
	}

	companyVerification := &CompanyVerification{
		ID:            primitive.NewObjectID(),
		CorrelationId: correlationId,
		Channels:      []Channel{},
		PreAudit:      *preAudit,
		Result:        IN_PROGRESS,
	}

	return companyVerification
}
