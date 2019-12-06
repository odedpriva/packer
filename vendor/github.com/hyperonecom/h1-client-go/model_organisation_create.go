/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// OrganisationCreate struct for OrganisationCreate
type OrganisationCreate struct {
	Name         string                           `json:"name,omitempty"`
	Service      string                           `json:"service,omitempty"`
	Billing      OrganisationCreateBilling        `json:"billing,omitempty"`
	AccessRights []OrganisationCreateAccessRights `json:"accessRights,omitempty"`
	Tag          map[string]string                `json:"tag,omitempty"`
}
