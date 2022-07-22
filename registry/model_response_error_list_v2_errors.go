/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseErrorListV2Errors struct {
	// The code of the error encountered. Where the error is specific to the respondent, an application-specific error code, expressed as a string value. If the error is application-specific, the URN code that the specific error extends must be provided in the meta object. Otherwise, the value is the error code URN.
	Code string `json:"code"`
	// A short, human-readable summary of the problem that MUST NOT change from occurrence to occurrence of the problem represented by the error code.
	Title string `json:"title"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`
	Meta *MetaError `json:"meta,omitempty"`
}