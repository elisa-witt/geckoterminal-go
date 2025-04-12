package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GeckoTerminalAPI struct {
	baseUrl string
}

func NewGeckoTerminalAPI() *GeckoTerminalAPI {
	return &GeckoTerminalAPI{
		baseUrl: "https://api.geckoterminal.com/api/v2/",
	}
}

// fetchRequest performs HTTP requests with retry logic and optional proxy support.
func (api *GeckoTerminalAPI) fetchRequest(path, method string, body []byte, retry int) (map[string]interface{}, error) {
	client := &http.Client{}
	if ZENROWS_API_KEY != "" {
		proxyURL, err := url.Parse("http://username:password@superproxy.zenrows.com:1337")
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse proxy URL")
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	urlStr := api.baseUrl + path
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	for i := 0; i <= retry; i++ {
		resp, err := client.Do(req)
		if err != nil {
			logrus.Errorf("Request failed: %v", err)
			if i < retry {
				logrus.Info("Retrying...")
				time.Sleep(5 * time.Second)
				continue
			}
			return nil, errors.Wrap(err, "max retries reached")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			return nil, fmt.Errorf("response status %d: %s", resp.StatusCode, string(bodyBytes))
		}

		var data map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, errors.Wrap(err, "failed to decode response")
		}
		return data, nil
	}
	return nil, fmt.Errorf("unexpected error")
}

// GetNetworks fetches a list of networks.
func (api *GeckoTerminalAPI) GetNetworks(page int) ([]Network, error) {
	path := fmt.Sprintf("networks?page=%d", page)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for networks")
	}
	var networks []Network
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var network Network
		if err := json.Unmarshal(itemJSON, &network); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal network")
		}
		networks = append(networks, network)
	}
	return networks, nil
}

// GetAllTrendingPools fetches trending pools across all networks.
func (api *GeckoTerminalAPI) GetAllTrendingPools(page int, duration string) ([]TrendingPool, error) {
	path := fmt.Sprintf("networks/trending_pools?page=%d&duration=%s", page, duration)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for trending pools")
	}
	var pools []TrendingPool
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var pool TrendingPool
		if err := json.Unmarshal(itemJSON, &pool); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal trending pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// GetTrendingPoolsFromNetwork fetches trending pools for a specific network.
func (api *GeckoTerminalAPI) GetTrendingPoolsFromNetwork(network string, page int, duration string) ([]TrendingPoolFromNetwork, error) {
	path := fmt.Sprintf("networks/%s/trending_pools?page=%d&duration=%s", network, page, duration)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for trending pools from network")
	}
	var pools []TrendingPoolFromNetwork
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var pool TrendingPoolFromNetwork
		if err := json.Unmarshal(itemJSON, &pool); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal trending pool from network")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// GetLatestPools fetches the latest pools across all networks.
func (api *GeckoTerminalAPI) GetLatestPools(page int, include string) ([]LatestPool, error) {
	params := url.Values{"page": {fmt.Sprintf("%d", page)}, "include": {include}}
	path := "networks/new_pools?" + params.Encode()
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for latest pools")
	}
	var pools []LatestPool
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var pool LatestPool
		if err := json.Unmarshal(itemJSON, &pool); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal latest pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// GetLatestPoolsFromNetwork fetches the latest pools for a specific network.
func (api *GeckoTerminalAPI) GetLatestPoolsFromNetwork(page int, network, include string) ([]Pool, error) {
	params := url.Values{"page": {fmt.Sprintf("%d", page)}, "include": {include}}
	path := fmt.Sprintf("networks/%s/new_pools?%s", network, params.Encode())
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for latest pools from network")
	}
	var pools []Pool
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var pool Pool
		if err := json.Unmarshal(itemJSON, &pool); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// GetRecentlyUpdatedTokens fetches recently updated tokens.
func (api *GeckoTerminalAPI) GetRecentlyUpdatedTokens(network string, limit int) ([]RecentlyUpdatedToken, error) {
	path := "tokens/info_recently_updated"
	if network != "" {
		path += "?network=" + network
	}
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for recently updated tokens")
	}
	var tokens []RecentlyUpdatedToken
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var token RecentlyUpdatedToken
		if err := json.Unmarshal(itemJSON, &token); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal token")
		}
		tokens = append(tokens, token)
	}
	if len(tokens) > limit {
		tokens = tokens[:limit]
	}
	return tokens, nil
}

// GetPoolInfo fetches detailed info for a specific pool.
func (api *GeckoTerminalAPI) GetPoolInfo(address, network string) (*TokenInfo, error) {
	path := fmt.Sprintf("networks/%s/pools/%s/info", network, address)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for pool info")
	}
	dataJSON, _ := json.Marshal(data)
	var info TokenInfo
	if err := json.Unmarshal(dataJSON, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal pool info")
	}
	return &info, nil
}

// GetPoolMarketInfo fetches market info for a specific pool.
func (api *GeckoTerminalAPI) GetPoolMarketInfo(address, network string) (*PoolMarketInfo, error) {
	path := fmt.Sprintf("networks/%s/pools/%s", network, address)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for pool market info")
	}
	dataJSON, _ := json.Marshal(data)
	var info PoolMarketInfo
	if err := json.Unmarshal(dataJSON, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal pool market info")
	}
	return &info, nil
}

// GetTokenInfo fetches detailed info for a specific token.
func (api *GeckoTerminalAPI) GetTokenInfo(address, network string) (*TokenInfo, error) {
	path := fmt.Sprintf("networks/%s/tokens/%s/info", network, address)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for token info")
	}
	dataJSON, _ := json.Marshal(data)
	var info TokenInfo
	if err := json.Unmarshal(dataJSON, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal token info")
	}
	return &info, nil
}

// GetTokenMarketInfo fetches market info for a specific token.
func (api *GeckoTerminalAPI) GetTokenMarketInfo(address, network string) (*TokenMarketInfo, error) {
	path := fmt.Sprintf("networks/%s/tokens/%s", network, address)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for token market info")
	}
	dataJSON, _ := json.Marshal(data)
	var info TokenMarketInfo
	if err := json.Unmarshal(dataJSON, &info); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal token market info")
	}
	return &info, nil
}

// GetTokenPoolInfo fetches pool info for a specific token.
func (api *GeckoTerminalAPI) GetTokenPoolInfo(address, network string) ([]TokenPoolInfo, error) {
	path := fmt.Sprintf("networks/%s/tokens/%s/pools", network, address)
	response, err := api.fetchRequest(path, "GET", nil, 3)
	if err != nil {
		return nil, err
	}
	data, ok := response["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format for token pool info")
	}
	var pools []TokenPoolInfo
	for _, item := range data {
		itemJSON, _ := json.Marshal(item)
		var pool TokenPoolInfo
		if err := json.Unmarshal(itemJSON, &pool); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal token pool info")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

