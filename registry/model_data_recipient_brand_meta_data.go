/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Metadata related to Data Recipient Brand
type DataRecipientBrandMetaData struct {
	// Unique id of the Data Recipient brand issued by the CDR Register
	DataRecipientBrandId string `json:"dataRecipientBrandId"`
	// Data Recipient Brand name
	BrandName string `json:"brandName"`
	// Data Recipient Brand logo URI
	LogoUri string `json:"logoUri"`
	SoftwareProducts []SoftwareProductMetaData `json:"softwareProducts,omitempty"`
	// Data Recipient Brand status in the CDR Register
	Status string `json:"status"`
}
