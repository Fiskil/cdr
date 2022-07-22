/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The data that is common to all organisations, regardless of the type (e.g. company, trust, partnership, government)
type LegalEntityDetail struct {
	// Unique id of the organisation issued by the CDR Register
	LegalEntityId string `json:"legalEntityId"`
	// Unique legal name of the organisation
	LegalEntityName string `json:"legalEntityName"`
	// Legal Entity logo URI
	LogoUri string `json:"logoUri"`
	// Unique registration number (if the company is registered outside Australia)
	RegistrationNumber string `json:"registrationNumber,omitempty"`
	// Date of registration (if the company is registered outside Australia)
	RegistrationDate string `json:"registrationDate,omitempty"`
	// Country of registeration (if the company is registered outside Australia)
	RegisteredCountry string `json:"registeredCountry,omitempty"`
	// Australian Business Number for the organisation
	Abn string `json:"abn,omitempty"`
	// Australian Company Number for the organisation
	Acn string `json:"acn,omitempty"`
	// Australian Registered Body Number.  ARBNs are issued to registrable Australian bodies and foreign companies
	Arbn string `json:"arbn,omitempty"`
	// ANZSIC division of the organisation. **[[ANZSIC-2006]](#iref-ANZSIC-2006)**
	AnzsicDivision string `json:"anzsicDivision,omitempty"`
	// Legal organisation type
	OrganisationType string `json:"organisationType,omitempty"`
	Status string `json:"status"`
}