const axios = require('axios');

const ip = "1.1.1.1";
const url = `https://openipapi.cloudshield.club/api?ip=${ip}`;

axios.get(url)
    .then(response => {
        const data = response.data;
        console.log(`IP: ${data.ip}\nCountry: ${data.country}\nCity: ${data.city}`);
    })
    .catch(error => {
        console.error(`Error: ${error.message}`);
    });
