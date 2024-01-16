package command

type ModelsTemplateCollectionRetrievalResponseNP struct {
	Id                     int64                                                                   `json:"Id"`
	CommonName             string                                                                  `json:"CommonName"`
	TemplateName           string                                                                  `json:"TemplateName"`
	Oid                    string                                                                  `json:"Oid"`
	KeySize                string                                                                  `json:"KeySize"`
	KeyType                string                                                                  `json:"KeyType"`
	ForestRoot             *string                                                                 `json:"ForestRoot,omitempty"`
	ConfigurationTenant    string                                                                  `json:"ConfigurationTenant"`
	FriendlyName           *string                                                                 `json:"FriendlyName,omitempty"`
	KeyRetention           int64                                                                   `json:"KeyRetention"`
	KeyRetentionDays       *int64                                                                  `json:"KeyRetentionDays,omitempty"`
	KeyArchival            bool                                                                    `json:"KeyArchival"`
	EnrollmentFields       []ModelsTemplateCollectionRetrievalResponseTemplateEnrollmentFieldModel `json:"EnrollmentFields,omitempty"`
	AllowedEnrollmentTypes int64                                                                   `json:"AllowedEnrollmentTypes"`
	TemplateRegexes        []ModelsTemplateCollectionRetrievalResponseTemplateRegexModel           `json:"TemplateRegexes,omitempty"`
	UseAllowedRequesters   bool                                                                    `json:"UseAllowedRequesters"`
	AllowedRequesters      []string                                                                `json:"AllowedRequesters,omitempty"`
	DisplayName            string                                                                  `json:"DisplayName"`
	RequiresApproval       bool                                                                    `json:"RequiresApproval"`
	KeyUsage               int64                                                                   `json:"KeyUsage"`
	ExtendedKeyUsages      []ModelsTemplateCollectionRetrievalResponseExtendedKeyUsageModel        `json:"ExtendedKeyUsages,omitempty"`
	AdditionalProperties   map[string]interface{}
}
