/*
Payment Method Service v2_test

Testing PaymentMethodApiService

*/

// Code generated by OpenAPI Generator

package xendit

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	xendit "github.com/xendit/xendit-go/v3"
)

func Test_xendit_PaymentMethodApiService(t *testing.T) {

	apiKey := os.Getenv("XND_APIKEY")
	if apiKey == "" {
		t.Skip("XND_APIKEY not set")
	}
	
	apiClient := xendit.NewClient(apiKey)

	t.Run("Test PaymentMethodApiService CreatePaymentMethod", func(t *testing.T) {

		resp, httpRes, err := apiClient.PaymentMethodApi.CreatePaymentMethod(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService GetPaymentMethodByID", func(t *testing.T) {

		var paymentMethodId string

		resp, httpRes, err := apiClient.PaymentMethodApi.GetPaymentMethodByID(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService GetPaymentsByPaymentMethodId", func(t *testing.T) {

		var paymentMethodId string

		resp, httpRes, err := apiClient.PaymentMethodApi.GetPaymentsByPaymentMethodId(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService PatchPaymentMethod", func(t *testing.T) {

		var paymentMethodId string

		resp, httpRes, err := apiClient.PaymentMethodApi.PatchPaymentMethod(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService GetAllPaymentMethods", func(t *testing.T) {

		resp, httpRes, err := apiClient.PaymentMethodApi.GetAllPaymentMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService ExpirePaymentMethod", func(t *testing.T) {

		var paymentMethodId string

		resp, httpRes, err := apiClient.PaymentMethodApi.ExpirePaymentMethod(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService AuthPaymentMethod", func(t *testing.T) {

		var paymentMethodId string

		resp, httpRes, err := apiClient.PaymentMethodApi.AuthPaymentMethod(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodApiService SimulatePayment", func(t *testing.T) {

		var paymentMethodId string

		httpRes, err := apiClient.PaymentMethodApi.SimulatePayment(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
