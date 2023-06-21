<h1 align="center">Uniswap Tracker</h1>

<h3 align="center"> High Performance Go-Fiber Microservice to Track UniswapV3 Liquidity Pools</h3>

<!-- TABLE OF CONTENTS -->
<details open>
  <summary>Table of Contents</summary>
  <ul>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#tech-stack">Tech Stack</a></li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#how-to-use">How to use?</a></li>
  </ul>
</details>

## About The Project

Monitoring service for Uniswap V3 pools that continuously tracks and logs essential data points, stores them in a persistent datastore, and provides access to the data through a REST endpoint.

## Tech Stack

[![](https://img.shields.io/badge/Built_with-Go-green?style=for-the-badge&logo=Go)](https://go.dev/)

## Prerequisites

Download and install [Golang 1.20](https://go.dev/doc/install) (or higher).

## How To Use?

1. Navigate to `uniswap-tracker/`:
   ``` 
   cd /path/to/folder/uniswap-tracker/
   ```
2. Open `config.json` file and fill in the `eth_client_url` field. This is useful to connect to the Ethereum Node. 
Also fill in the `pool_address` of the Uniswap V3 Pool you want to track. Rest of the fields can be left to default.
Here are some details about the fields in the config file:
   1. `pause_duration`: Pause duration between consecutive RPC Calls. Default is 6 seconds. (Default=6000)
   2. `log_frequency`: After this many number of blocks the pool data be stored in the db. (Default=12)
3. Get dependencies:
   ``` 
   go mod tidy
   ```
4. Run the app:
   ``` 
   go run . 
   # use "--verbose" flag to get additional logs
   go run . --verbose 
   ```
5. Get latest data with pool_id being 1:
    ```
    curl -X GET "http://127.0.0.1:3000/v1/api/pool/1?latest"
    ```
6. Get block 420 data with pool_id being 1:
    ```
    curl -X GET "http://127.0.0.1:3000/v1/api/pool/1?block=420"
    ```
7. Get historic data:
    ```
    curl -X GET "http://127.0.0.1:3000/v1/api/pool/1/historic"
    ```
8. In case you want to know the pool_id for all the pool addresses tracked:
    ```
    curl -X GET "http://127.0.0.1:3000/v1/api/pool/mapping"
    ```

Thank you!
