package utils

import (
	"github.com/google/uuid"
	"hash/fnv"
)

type SerializableType int

const (
	JSON   SerializableType = iota
	String SerializableType = 1
	Avro   SerializableType = 2
)

func RemoveFromSlice(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func RemoveFromSliceInterface(slice []interface{}, s interface{}) []interface{} {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func NewUUIDAsString() string {
	return uuid.New().String()
}

func GetPartitionNumber(topicName string, noOfPartitions int) int {
	return int(hash(topicName) % uint32(noOfPartitions))
}

func GetPartitionNumberFromKey(key string, noOfPartitions int) int {
	return int(hash(key) % uint32(noOfPartitions))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
