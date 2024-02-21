import http from 'k6/http';
import { randomItem, randomString, randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export default function () {

    const idCliente = randomIntBetween(1, 5);

    // const url = `http://localhost:3000/clientes/${idCliente.toString()}/transacoes`; // Degub 
    const url = `http://localhost:9999/clientes/${idCliente.toString()}/transacoes`; // Docker 

    const headers = { headers: { 'Accept-Encoding': 'gzip, deflate', 'Content-Type': 'application/json' } };

    //const payload = { valor: 1, tipo: 'c', descricao: 'FIXO' };
    const payload = { valor: randomIntBetween(100, 10000), tipo: randomItem(['c', 'd']), descricao: randomString(10, `aeioubcdfghijpqrstuv`) };
    //console.log(payload);

    const res = http.post(url, JSON.stringify(payload), headers);
    //console.log(res.body);
}
