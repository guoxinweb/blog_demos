/*
 * Strimzi Kafka Bridge API Reference
 *
 * The Strimzi Kafka Bridge provides a REST API for integrating HTTP based client applications with a Kafka cluster. You can use the API to create and manage consumers and send and receive records over HTTP rather than the native Kafka protocol. 
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProducerRecordToPartitionList struct {
	Records []ProducerRecordToPartition `json:"records,omitempty"`
}
