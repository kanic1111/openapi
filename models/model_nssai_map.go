/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NssaiMap struct {
	ServingSnssai *Snssai `json:"servingSnssai" yaml:"servingSnssai" bson:"servingSnssai" mapstructure:"ServingSnssai"`
	HomeSnssai    *Snssai `json:"homeSnssai" yaml:"homeSnssai" bson:"homeSnssai" mapstructure:"HomeSnssai"`
}
