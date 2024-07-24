/*
Slurm Rest API

Testing SlurmAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package oapigen

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pcolladosoto/goslurm/oapigen"
)

func Test_oapigen_SlurmAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlurmAPIService SlurmV0038CancelJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		httpRes, err := apiClient.SlurmAPI.SlurmV0038CancelJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038Diag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetPartition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetPartition(context.Background(), partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetPartitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetReservation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetReservation(context.Background(), reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetReservations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038Ping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Ping(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038SlurmctldGetLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038SlurmctldGetLicenses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038SubmitJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038SubmitJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038UpdateJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		httpRes, err := apiClient.SlurmAPI.SlurmV0038UpdateJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
