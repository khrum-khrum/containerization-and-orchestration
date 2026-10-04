import http from 'k6/http';
import { check } from 'k6';

const TARGET = __ENV.TARGET || 'http://shop-api.lab-4.svc.cluster.local';

export const options = {
  scenarios: {
    orders: {
      executor: 'constant-arrival-rate',
      rate: Number(__ENV.RATE || 30),
      timeUnit: '1s',
      duration: __ENV.DURATION || '10m',
      preAllocatedVUs: 20,
      maxVUs: 200,
    },
  },
  thresholds: {
    http_req_duration: ['p(95)<300'],
    http_req_failed: ['rate<0.01'],
  },
  summaryTrendStats: ['avg', 'p(50)', 'p(95)', 'p(99)', 'max'],
};

export default function () {
  const res = http.post(`${TARGET}/order`, JSON.stringify({ item: `item-${__VU}-${__ITER}` }), {
    headers: { 'Content-Type': 'application/json' },
    timeout: '5s',
  });
  check(res, { 'order created': (r) => r.status === 201 });
}
