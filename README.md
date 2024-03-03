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
``` 

### Example Response
```
{
  "ip": "1.1.1.1",
  "country": "United States",
  "city": "Buffalo, NY"
}

```

### Handling Errors
If the ip parameter is missing, the API will respond with a 400 status code and an error message:
``` 
{
  "error": "IP parameter is missing"
}

```

---

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Issues
If you encounter any issues or have questions, feel free to open an [open an issue](https://github.com/ajdev05/OpenIP-API/TrackAPP/issues) on this repository. We appreciate your feedback and contributions!
