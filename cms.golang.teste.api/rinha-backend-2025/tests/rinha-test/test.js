import http from "k6/http";
import { check } from "k6";

export let options = {
  vus: 50,
  duration: "30s",
  thresholds: { http_req_duration: ["p(95)<200"] },
};

export default function () {
  const payload = JSON.stringify({
    correlationId: `${__VU}-${__ITER}`,
    amount: Math.random() * 100,
  });
  const url = `http://localhost:8080/payments`;
  const headers = { headers: { "Content-Type": "application/json" } };
  const res = http.post(url, payload, headers);
  check(res, { accepted: (r) => r.status === 202 });
}

// k6 run --out prometheus-remote-write=http://localhost:9090/api/v1/write loadtest/test.js
