package main

type Network struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                     string  `json:"name"`
		CoingeckoAssetPlatformID *string `json:"coingecko_asset_platform_id,omitempty"`
	} `json:"attributes"`
}

type TrendingPool struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		BaseTokenPriceUSD             string  `json:"base_token_price_usd"`
		BaseTokenPriceNativeCurrency  string  `json:"base_token_price_native_currency"`
		QuoteTokenPriceUSD            string  `json:"quote_token_price_usd"`
		QuoteTokenPriceNativeCurrency string  `json:"quote_token_price_native_currency"`
		BaseTokenPriceQuoteToken      string  `json:"base_token_price_quote_token"`
		QuoteTokenPriceBaseToken      string  `json:"quote_token_price_base_token"`
		Address                       string  `json:"address"`
		Name                          string  `json:"name"`
		PoolCreatedAt                 string  `json:"pool_created_at"`
		FDVUSD                        string  `json:"fdv_usd"`
		MarketCapUSD                  *string `json:"market_cap_usd,omitempty"`
		PriceChangePercentage         struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"price_change_percentage"`
		Transactions struct {
			M5 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m5"`
			M15 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m15"`
			M30 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m30"`
			H1 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"h1"`
			H24 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"h24"`
		} `json:"transactions"`
		VolumeUSD struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"volume_usd"`
		ReserveInUSD string `json:"reserve_in_usd"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
		QuoteToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"quote_token"`
		Network struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"network"`
		Dex struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dex"`
	} `json:"relationships"`
}

type TrendingPoolFromNetwork struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                          string  `json:"name"`
		Address                       string  `json:"address"`
		BaseTokenPriceUSD             *string `json:"base_token_price_usd,omitempty"`
		BaseTokenPriceNativeCurrency  *string `json:"base_token_price_native_currency,omitempty"`
		QuoteTokenPriceUSD            *string `json:"quote_token_price_usd,omitempty"`
		QuoteTokenPriceNativeCurrency *string `json:"quote_token_price_native_currency,omitempty"`
		BaseTokenPriceQuoteToken      *string `json:"base_token_price_quote_token,omitempty"`
		QuoteTokenPriceBaseToken      *string `json:"quote_token_price_base_token,omitempty"`
		PoolCreatedAt                 string  `json:"pool_created_at"`
		FDVUSD                        *string `json:"fdv_usd,omitempty"`
		MarketCapUSD                  *string `json:"market_cap_usd,omitempty"`
		PriceChangePercentage         struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"price_change_percentage"`
		Transactions *struct {
			M5 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m5,omitempty"`
			M15 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m15,omitempty"`
			M30 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m30,omitempty"`
			H1 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"h1,omitempty"`
			H24 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"h24,omitempty"`
		} `json:"transactions,omitempty"`
		VolumeUSD struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"volume_usd"`
		ReserveInUSD string `json:"reserve_in_usd"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
		QuoteToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"quote_token"`
		Dex struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dex"`
	} `json:"relationships"`
}

type LatestPool struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                          string  `json:"name"`
		Address                       string  `json:"address"`
		BaseTokenPriceUSD             *string `json:"base_token_price_usd,omitempty"`
		BaseTokenPriceNativeCurrency  *string `json:"base_token_price_native_currency,omitempty"`
		QuoteTokenPriceUSD            *string `json:"quote_token_price_usd,omitempty"`
		QuoteTokenPriceNativeCurrency *string `json:"quote_token_price_native_currency,omitempty"`
		BaseTokenPriceQuoteToken      *string `json:"base_token_price_quote_token,omitempty"`
		QuoteTokenPriceBaseToken      *string `json:"quote_token_price_base_token,omitempty"`
		PoolCreatedAt                 string  `json:"pool_created_at"`
		FDVUSD                        string  `json:"fdv_usd"`
		MarketCapUSD                  *string `json:"market_cap_usd,omitempty"`
		PriceChangePercentage         struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"price_change_percentage"`
		Transactions *struct {
			M5 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m5,omitempty"`
			M15 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m15,omitempty"`
			M30 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m30,omitempty"`
			H1 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"h1,omitempty"`
			H24 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"h24,omitempty"`
		} `json:"transactions,omitempty"`
		VolumeUSD struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"volume_usd"`
		ReserveInUSD string `json:"reserve_in_usd"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
		QuoteToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"quote_token"`
		Network struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"network"`
		Dex struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dex"`
	} `json:"relationships"`
}

type Pool struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                          string  `json:"name"`
		Address                       string  `json:"address"`
		BaseTokenPriceUSD             *string `json:"base_token_price_usd,omitempty"`
		BaseTokenPriceNativeCurrency  *string `json:"base_token_price_native_currency,omitempty"`
		QuoteTokenPriceUSD            *string `json:"quote_token_price_usd,omitempty"`
		QuoteTokenPriceNativeCurrency *string `json:"quote_token_price_native_currency,omitempty"`
		BaseTokenPriceQuoteToken      *string `json:"base_token_price_quote_token,omitempty"`
		QuoteTokenPriceBaseToken      *string `json:"quote_token_price_base_token,omitempty"`
		PoolCreatedAt                 string  `json:"pool_created_at"`
		FDVUSD                        *string `json:"fdv_usd,omitempty"`
		MarketCapUSD                  *string `json:"market_cap_usd,omitempty"`
		PriceChangePercentage         struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"price_change_percentage"`
		Transactions *struct {
			M5 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m5,omitempty"`
		} `json:"transactions,omitempty"`
		VolumeUSD struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"volume_usd"`
		ReserveInUSD string `json:"reserve_in_usd"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
		QuoteToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"quote_token"`
		Dex struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dex"`
	} `json:"relationships"`
}

type TokenInfo struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address         string   `json:"address"`
		Name            string   `json:"name"`
		Symbol          string   `json:"symbol"`
		Decimals        int      `json:"decimals"`
		ImageURL        *string  `json:"image_url,omitempty"`
		CoingeckoCoinID *string  `json:"coingecko_coin_id,omitempty"`
		Websites        []string `json:"websites,omitempty"`
		Description     *string  `json:"description,omitempty"`
		DiscordURL      *string  `json:"discord_url,omitempty"`
		TelegramHandle  *string  `json:"telegram_handle,omitempty"`
		TwitterHandle   *string  `json:"twitter_handle,omitempty"`
		Categories      []string `json:"categories"`
		GTCategoryIDs   []string `json:"gt_category_ids"`
		GTScore         *float64 `json:"gt_score,omitempty"`
		GTScoreDetails  struct {
			Pool        *float64 `json:"pool,omitempty"`
			Transaction *float64 `json:"transaction,omitempty"`
			Creation    *float64 `json:"creation,omitempty"`
			Info        *float64 `json:"info,omitempty"`
			Holders     *float64 `json:"holders,omitempty"`
		} `json:"gt_score_details"`
		Holders *struct {
			Count                  int `json:"count"`
			DistributionPercentage *struct {
				Top10         *string `json:"top_10,omitempty"`
				ElevenTo20    *string `json:"11_20,omitempty"`
				TwentyOneTo40 *string `json:"21_40,omitempty"`
				Rest          *string `json:"rest,omitempty"`
			} `json:"distribution_percentage,omitempty"`
			LastUpdated *string `json:"last_updated,omitempty"`
		} `json:"holders,omitempty"`
		MetadataUpdatedAt *string `json:"metadata_updated_at,omitempty"`
		MintAuthority     *string `json:"mint_authority,omitempty"`
		FreezeAuthority   *string `json:"freeze_authority,omitempty"`
	} `json:"attributes"`
}

type TokenPoolInfo struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                          string  `json:"name"`
		Address                       string  `json:"address"`
		BaseTokenPriceUSD             string  `json:"base_token_price_usd"`
		QuoteTokenPriceUSD            string  `json:"quote_token_price_usd"`
		BaseTokenPriceNativeCurrency  string  `json:"base_token_price_native_currency"`
		QuoteTokenPriceNativeCurrency string  `json:"quote_token_price_native_currency"`
		BaseTokenPriceQuoteToken      string  `json:"base_token_price_quote_token"`
		QuoteTokenPriceBaseToken      string  `json:"quote_token_price_base_token"`
		PoolCreatedAt                 string  `json:"pool_created_at"`
		ReserveInUSD                  string  `json:"reserve_in_usd"`
		FDVUSD                        string  `json:"fdv_usd"`
		MarketCapUSD                  *string `json:"market_cap_usd,omitempty"`
	} `json:"attributes"`
}

type PoolMarketInfo struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                          string  `json:"name"`
		Address                       string  `json:"address"`
		BaseTokenPriceUSD             *string `json:"base_token_price_usd,omitempty"`
		BaseTokenPriceNativeCurrency  *string `json:"base_token_price_native_currency,omitempty"`
		QuoteTokenPriceUSD            *string `json:"quote_token_price_usd,omitempty"`
		QuoteTokenPriceNativeCurrency *string `json:"quote_token_price_native_currency,omitempty"`
		BaseTokenPriceQuoteToken      *string `json:"base_token_price_quote_token,omitempty"`
		QuoteTokenPriceBaseToken      *string `json:"quote_token_price_base_token,omitempty"`
		PoolCreatedAt                 string  `json:"pool_created_at"`
		FDVUSD                        string  `json:"fdv_usd"`
		MarketCapUSD                  *string `json:"market_cap_usd,omitempty"`
		PriceChangePercentage         struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"price_change_percentage"`
		Transactions struct {
			M5 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m5"`
			M15 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m15"`
			M30 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"m30"`
			H1 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"h1"`
			H24 struct {
				Buys    int `json:"buys"`
				Sells   int `json:"sells"`
				Buyers  int `json:"buyers"`
				Sellers int `json:"sellers"`
			} `json:"h24"`
		} `json:"transactions"`
		VolumeUSD struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"volume_usd"`
		ReserveInUSD              string  `json:"reserve_in_usd"`
		LockedLiquidityPercentage *string `json:"locked_liquidity_percentage,omitempty"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
		QuoteToken struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"quote_token"`
		Dex struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dex"`
	} `json:"relationships"`
}

type TokenMarketInfo struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address               string  `json:"address"`
		Name                  string  `json:"name"`
		Symbol                string  `json:"symbol"`
		Decimals              int     `json:"decimals"`
		ImageURL              string  `json:"image_url"`
		CoingeckoCoinID       *string `json:"coingecko_coin_id,omitempty"`
		TotalSupply           string  `json:"total_supply"`
		PriceUSD              string  `json:"price_usd"`
		FDVUSD                string  `json:"fdv_usd"`
		TotalReserveInUSD     *string `json:"total_reserve_in_usd,omitempty"`
		MarketCapUSD          *string `json:"market_cap_usd,omitempty"`
		PriceChangePercentage *struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 string  `json:"h24"`
		} `json:"price_change_percentage,omitempty"`
		Transactions *struct {
			M5 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m5,omitempty"`
			M15 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m15,omitempty"`
			M30 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"m30,omitempty"`
			H1 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"h1,omitempty"`
			H24 *struct {
				Buys    *int `json:"buys,omitempty"`
				Sells   *int `json:"sells,omitempty"`
				Buyers  *int `json:"buyers,omitempty"`
				Sellers *int `json:"sellers,omitempty"`
			} `json:"h24,omitempty"`
		} `json:"transactions,omitempty"`
		VolumeUSD *struct {
			M5  *string `json:"m5,omitempty"`
			H1  *string `json:"h1,omitempty"`
			H6  *string `json:"h6,omitempty"`
			H24 *string `json:"h24,omitempty"`
		} `json:"volume_usd,omitempty"`
	} `json:"attributes"`
}

type RecentlyUpdatedToken struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address           string   `json:"address"`
		Name              string   `json:"name"`
		Symbol            string   `json:"symbol"`
		Decimals          int      `json:"decimals"`
		ImageURL          *string  `json:"image_url,omitempty"`
		CoingeckoCoinID   *string  `json:"coingecko_coin_id,omitempty"`
		Websites          []string `json:"websites,omitempty"`
		Description       *string  `json:"description,omitempty"`
		DiscordURL        *string  `json:"discord_url,omitempty"`
		TelegramHandle    *string  `json:"telegram_handle,omitempty"`
		TwitterHandle     *string  `json:"twitter_handle,omitempty"`
		Categories        []string `json:"categories,omitempty"`
		GTCategoryIDs     []string `json:"gt_category_ids,omitempty"`
		GTScore           *float64 `json:"gt_score,omitempty"`
		MetadataUpdatedAt *string  `json:"metadata_updated_at,omitempty"`
	} `json:"attributes"`
	Relationships *struct {
		Network struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"network"`
	} `json:"relationships,omitempty"`
}
