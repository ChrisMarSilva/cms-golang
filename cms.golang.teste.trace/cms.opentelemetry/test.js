import http from "k6/http";
import { check } from "k6";

export let options = {
  vus: 500,
  duration: "30s",
  thresholds: { http_req_duration: ["p(95)<200"] },
};

export default function () {
  const url = `http://localhost:8080/`;
  const res = http.get(url);
  check(res, { accepted: (r) => r.status === 202 });
}
