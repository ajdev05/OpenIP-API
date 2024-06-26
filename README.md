  
![GitHub language](https://img.shields.io/badge/language-Python-blue)   
![GitHub language](https://img.shields.io/badge/language-JavaScript-yellow)

# OpenIP API - IP to Country Resolver
  
Welcome to the OpenIP API, a free and unlimited IP to Country resolver that empowers developers to seamlessly retrieve e country and city information based on a given IP address. This API is built using the Flask web framework in Python, offering a robust and efficient solution for geolocation needs.

## Features  
 
- **Free and Unlimited Access:** OpenIP API provides free and unrestricted access to its IP to Country resolution service, allowing developers to integrate geolocation functionality without worrying about usage limits.

- **Fast and Reliable:** Built on the Flask web framework, the OpenIP API ensures speedy and reliable responses, making it suitable for real-time applications and services.

- **Simple Integration:** With a RESTful design, the API is easy to integrate into a variety of applications, including web, mobile, and desktop platforms. The use of standard HTTP methods and JSON responses simplifies the integration process.



## Table of Contents

- [Usage](#usage)
  - [API Endpoint](#api-endpoint)
  - [Example Request](#example-request)
  - [Example Response](#example-response)
  - [Handling Errors](#handling-errors)
- [Code Examples](#Code-Examples)
- [License](#license)
- [Issues](#issues)
- [Discussion](#Discussion)

--- 

## Usage

Simply make a GET request to the API endpoint with the desired IP address as a parameter, and the API will respond with detailed information about the corresponding country and city.

### API Endpoint

- **Endpoint:** `https://openipapi.cloudshield.club/api`
- **Method:** `GET`
- **Parameters:**
  - `ip` (required): The IP address for which you want to resolve the country and city.

### Example Request

```bash
curl https://openipapi.cloudshield.club/api?ip=1.1.1.1
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

## Code Examples

### Python

- [View Python3 example here](https://github.com/ajdev05/OpenIP-API/blob/main/ExampleGET.py) 

### Node.JS

- [View Node.JS example here](https://github.com/ajdev05/OpenIP-API/blob/main/ExampleGET.js)

### Go

- [View Python3 example here](https://github.com/ajdev05/OpenIP-API/blob/main/ExampleGET.go)
---

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Issues
If you encounter any issues or have questions, feel free to open an [open an issue](https://github.com/ajdev05/OpenIP-API/issues) on this repository. We appreciate your feedback and contributions!


### Discussion
Join the [Discussion](https://github.com/ajdev05/OpenIP-API/discussions/) and collaborate with the community. Whether you have questions, ideas, or just want to connect, this is the space to engage. We encourage an open and inclusive environment for sharing experiences and helping each other make the most out of the OpenIP API.

