import http from "k6/http";
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";
import { check, sleep } from "k6";

export const options = {
  // vus: 10,  // número de usuários virtuais
  // duration: "5s",  // duração total do teste

  stages: [
    { duration: "2m", target: 1000 }, // ramp up
    { duration: "8m", target: 1000 },
    { duration: "2m", target: 0 },
  ],

  // scenarios: {
  //   ramping: {
  //     executor: "ramping-vus",
  //     startVUs: 0,
  //     stages: [
  //       { duration: "2m", target: 1000 }, // ramp-up para 100 VUs em 30s
  //       { duration: "8m", target: 1000 }, // sustentar 100 VUs por 2m
  //       { duration: "2m", target: 0 }, // ramp-down em 30s
  //     ],
  //     gracefulRampDown: "10s",
  //   },

  //   // load_test: {
  //   //   executor: "ramping-arrival-rate",
  //   //   startRate: 0,
  //   //   timeUnit: "1s",
  //   //   preAllocatedVUs: 10,
  //   //   maxVUs: 2000,
  //   //   stages: [
  //   //     { target: 1000, duration: "2m" }, // chega a 1000 req/s em 1m
  //   //     { target: 1000, duration: "8m" }, // sustenta 1000 req/s por 3m
  //   //     { target: 0, duration: "2m" }, // resolvendo carga ao final
  //   //   ],
  //   // },
  // },

  // discardResponseBodies: true,

  thresholds: {
    http_req_failed: ["rate<0.01"], // menos de 1% de erros
    http_req_duration: ["p(95)<999"], // 95% das requisições < 999ms
  },
};

export default function () {
  const url = `http://localhost:9999/v1/person`;
  const payload = JSON.stringify({ name: `Generated UUID: ${uuidv4()}` });
  const params = { headers: { "Content-Type": "application/json" } };
  const res = http.post(url, payload, params);

  //check(res, { "status 202": (r) => r.status === 202 });
  check(res, { accepted: (r) => r.status === 202 });
  sleep(1);
}
