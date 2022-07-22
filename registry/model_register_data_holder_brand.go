/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type RegisterDataHolderBrand struct {
	// Unique id of the Data Holder Brand issued by the CDR Register
	DataHolderBrandId string `json:"dataHolderBrandId"`
	// The name of Data Holder Brand
	BrandName string `json:"brandName"`
	// The industries the Data Holder Brand belongs to. Please note that the CDR Register entity model is constrained to one industry per brand which is planned to be relaxed in the future.
	Industries []string `json:"industries"`
	// Brand logo URI
	LogoUri string `json:"logoUri"`
	LegalEntity *LegalEntityDetail `json:"legalEntity"`
	Status string `json:"status"`
	EndpointDetail *RegisterDataHolderBrandServiceEndpoint `json:"endpointDetail"`
	AuthDetails []RegisterDataHolderAuth `json:"authDetails"`
	// The date/time that the Data Holder Brand data was last updated in the Register
	LastUpdated time.Time `json:"lastUpdated"`
}
