/*
 * OFAC API
 *
 * OFAC (Office of Foreign Assets Control) API is designed to facilitate the enforcement of US government economic sanctions programs required by federal law. This project implements a modern REST HTTP API for companies and organizations to obey federal law and use OFAC data in their applications.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Specially designated national from OFAC list
type Sdn struct {
	EntityID string `json:"EntityID,omitempty"`
	SDNName  string `json:"SDNName,omitempty"`
	SDNType  string `json:"SDNType,omitempty"`
	Program  string `json:"Program,omitempty"`
	Title    string `json:"Title,omitempty"`
	Remarks  string `json:"Remarks,omitempty"`
}