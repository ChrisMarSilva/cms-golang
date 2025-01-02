import http from "k6/http";
import { check, sleep } from 'k6';

// export const options = {
//   vus: 1000,
//   duration: "1m", // "60s",
// };

// export const options = {
    // stages: [
        // { target: 50, duration: '5m' },
        // { target: 100, duration: '0m' },
        // { target: 100, duration: '5m' },
        // { target: 200, duration: '5m' },
    // ]
// };

export const options = {
  scenarios: {
    constant_load: {
      executor: "ramping-arrival-rate",
      startRate: 0,
      timeUnit: "1s",
      preAllocatedVUs: 10,
      maxVUs: 50,
      stages: [
        { duration: "30s", target: 50 },
        { duration: "1m30s", target: 100 },
        { duration: "30s", target: 0 },
      ],
    },
  },
};

export default () => { // export default function () {
  const host = "http://localhost:8080";
  const response = http.get(host);
  
  check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
  });
    
  sleep(1);
};