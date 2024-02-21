import http from 'k6/http';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export default function () {

    const idCliente = randomIntBetween(1, 5);

    //const url = `http://localhost:3000/clientes/${idCliente.toString()}/extrato`; // Degub 
    const url = `http://localhost:9999/clientes/${idCliente.toString()}/extrato`; // Docker 

    const params = { headers: { 'Accept-Encoding': 'gzip, deflate', 'Content-Type': 'application/json' } };

    const res = http.get(url, params);
    // console.log(res.body);
}
