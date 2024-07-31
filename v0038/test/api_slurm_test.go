/*
Slurm Rest API RO

Testing SlurmAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v0038

import (
	"context"
	"os"
	"testing"

	openapiclient "github.com/pcolladosoto/goslurm/v0038"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_oapigen_SlurmAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Host = "arc.ft.uam.es:6820"
	configuration.Scheme = "http"

	envUser, ok := os.LookupEnv("SLURM_USER")
	if !ok {
		t.Fatalf("couldn't find user on the SLURM_USER environment variable... is it defined?")
	}

	envToken, ok := os.LookupEnv("SLURM_JWT")
	if !ok {
		t.Fatalf("couldn't find token on the SLURM_JWT environment variable... is it defined?")
	}

	auth := context.WithValue(
		context.Background(),
		openapiclient.ContextAPIKeys,
		map[string]openapiclient.APIKey{
			"user":  {Key: envUser},
			"token": {Key: envToken},
		},
	)

	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlurmAPIService SlurmV0038Diag", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Diag(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetJob(auth, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetJobs", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetJobs(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetNode(auth, nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetNodes", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetNodes(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetPartition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetPartition(auth, partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetPartitions", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetPartitions(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetReservation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetReservation(auth, reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SlurmAPIService SlurmV0038GetReservations", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetReservations(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038Ping", func(t *testing.T) {

		// t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Ping(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038SlurmctldGetLicenses", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038SlurmctldGetLicenses(auth).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
