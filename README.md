# OpenIP API - IP to Country Resolver

Welcome to the OpenIP API, a free and unlimited IP to Country resolver that allows developers to retrieve country and city information based on an IP address. This API is built using the Flask web framework in Python.

## Table of Contents

- [Usage](#usage)
  - [API Endpoint](#api-endpoint)
  - [Example Request](#example-request)
  - [Example Response](#example-response)
  - [Handling Errors](#handling-errors)
- [License](#license)
- [Issues](#issues)

---

## Usage

### API Endpoint

- **Endpoint:** `https://open-ip-api.vercel.app/api`
- **Method:** `GET`
- **Parameters:**
  - `ip` (required): The IP address for which you want to resolve the country and city.

### Example Request

```bash
curl https://open-ip-api.vercel.app/api?ip=1.1.1.1
