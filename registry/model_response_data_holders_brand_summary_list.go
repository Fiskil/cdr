/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseDataHoldersBrandSummaryList struct {
	// Response data for the query
	Data []DataHolderBrandSummary `json:"data"`
	Links *Links `json:"links"`
	Meta *Meta `json:"meta"`
}
