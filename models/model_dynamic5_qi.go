/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Dynamic5Qi struct {
	ResourceType      QosResourceType `json:"resourceType"`
	PriorityLevel     int32           `json:"priorityLevel"`
	PacketDelayBudget int32           `json:"packetDelayBudget"`
	PacketErrRate     string          `json:"packetErrRate"`
	AverWindow        int32           `json:"averWindow,omitempty"`
	MaxDataBurstVol   int32           `json:"maxDataBurstVol,omitempty"`
}