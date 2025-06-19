package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEstimateRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request EstimateRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request",
			request: EstimateRequest{
				Pool:      "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
				Src:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Dst:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcAmount: "10000000",
			},
			wantErr: false,
		},
		{
			name: "valid request without 0x prefix",
			request: EstimateRequest{
				Pool:      "0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
				Src:       "dAC17F958D2ee523a2206206994597C13D831ec7",
				Dst:       "c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcAmount: "10000000",
			},
			wantErr: false,
		},
		{
			name: "empty pool address",
			request: EstimateRequest{
				Pool:      "",
				Src:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Dst:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcAmount: "10000000",
			},
			wantErr: true,
			errMsg:  "invalid pool address",
		},
		{
			name: "invalid pool address length",
			request: EstimateRequest{
				Pool:      "0x123",
				Src:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Dst:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcAmount: "10000000",
			},
			wantErr: true,
			errMsg:  "invalid pool address",
		},
		{
			name: "invalid pool address non-hex",
			request: EstimateRequest{
				Pool:      "0xZd4a11d5eeaac28ec3f61d100daf4d40471f1852",
				Src:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Dst:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcAmount: "10000000",
			},
			wantErr: true,
			errMsg:  "invalid pool address",
		},
		{
			name: "empty src amount",
			request: EstimateRequest{
				Pool:      "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
				Src:       "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Dst:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcAmount: "",
			},
			wantErr: true,
			errMsg:  "invalid amount format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid address with 0x prefix",
			address: "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
			wantErr: false,
		},
		{
			name:    "valid address without 0x prefix",
			address: "0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
			wantErr: false,
		},
		{
			name:    "valid address uppercase",
			address: "0x0D4A11D5EEAAC28EC3F61D100DAF4D40471F1852",
			wantErr: false,
		},
		{
			name:    "empty address",
			address: "",
			wantErr: true,
			errMsg:  "address cannot be empty",
		},
		{
			name:    "short address",
			address: "0x123",
			wantErr: true,
			errMsg:  "address must be 40 hex characters",
		},
		{
			name:    "long address",
			address: "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852ff",
			wantErr: true,
			errMsg:  "address must be 40 hex characters",
		},
		{
			name:    "non-hex characters",
			address: "0xZd4a11d5eeaac28ec3f61d100daf4d40471f1852",
			wantErr: true,
			errMsg:  "address must contain only hex characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAddress(tt.address)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
