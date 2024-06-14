/**
 * Copyright 2023-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prime

import (
	"context"
	"fmt"
	"github.com/coinbase-samples/core-go"
)

type CreateWalletRequest struct {
	PortfolioId string `json:"portfolio_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Type        string `json:"wallet_type"`
}

type CreateWalletResponse struct {
	ActivityId string               `json:"activity_id"`
	Name       string               `json:"name"`
	Symbol     string               `json:"symbol"`
	Type       string               `json:"wallet_type"`
	Request    *CreateWalletRequest `json:"request"`
}

func (c *Client) CreateWallet(ctx context.Context, request *CreateWalletRequest) (*CreateWalletResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/wallets", request.PortfolioId)

	response := &CreateWalletResponse{Request: request}

	if err := core.Post(ctx, c, path, core.EmptyQueryParams, request, response, addPrimeHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
