/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DataRecipientStatus struct {
	// Unique id of the Data Recipient Legal Entity issued by the CDR Register
	LegalEntityId string `json:"legalEntityId"`
	// Data Recipient status in the CDR Register
	Status string `json:"status"`
}
