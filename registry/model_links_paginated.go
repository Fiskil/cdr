/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LinksPaginated struct {
	// URI to the first page of this set. Mandatory if this response is not the first page
	First string `json:"first,omitempty"`
	// URI to the last page of this set. Mandatory if this response is not the last page
	Last string `json:"last,omitempty"`
	// URI to the next page of this set. Mandatory if this response is not the last page
	Next string `json:"next,omitempty"`
	// URI to the previous page of this set. Mandatory if this response is not the first page
	Prev string `json:"prev,omitempty"`
	// Fully qualified link to this API call
	Self string `json:"self"`
}
