import {randomString} from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import http from 'k6/http';
import {check} from "k6";

export const BASE_URI = 'http://localhost:8080'

export const options = {
    thresholds: {
        http_req_duration: ['p(90) < 400', 'p(95) < 800', 'p(99) < 1600'],
    },
    scenarios: {
        full_scenario_ramping_arrival_rate: {
            exec: 'fullScenario',
            executor: 'ramping-arrival-rate',
            startRate: 0,
            timeUnit: '1s',
            preAllocatedVUs: 1000,
            maxVUs: 1000,
            stages: [
                {target: 2000, duration: '4m'},
            ],
        },
    },
};

export function fullScenario() {
    const task_id = postTask().task_id;
    getTask(task_id);
    putTask(task_id);
    getTasks();
    deleteTask(task_id);
}

export function getTasks() {
    const params = {
        responseType: 'none',
        tags: {
            name: `${BASE_URI}/tasks`
        },
    };

    const response = http.get(`${BASE_URI}/tasks`, params);

    check(response, {
        'getOrders is status 200': (r) => r.status === 200,
    });
}

export function postTask() {
    const payload = JSON.stringify({
        title: randomString(20),
        description: randomString(400),
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {
            name: `${BASE_URI}/tasks`
        },
    };

    const response = http.post(`${BASE_URI}/tasks`, payload, params);

    check(response, {
        'postTask is status 201': (r) => r.status === 201,
    });

    return response.json();
}

export function getTask(task_id) {
    const params = {
        responseType: 'none',
        tags: {
            name: `${BASE_URI}/tasks/{taskId}`
        },
    };

    const response = http.get(`${BASE_URI}/tasks/${task_id}`, params);

    check(response, {
        'getTask is status 200': (r) => r.status === 200,
    });
}

export function putTask(task_id) {
    const payload = JSON.stringify({
        title: randomString(40),
        description: randomString(800),
    });

    const params = {
        responseType: 'none',
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {
            name: `${BASE_URI}/tasks/{taskId}`
        },
    };

    const response = http.put(`${BASE_URI}/tasks/${task_id}`, payload, params);

    check(response, {
        'putTask is status 204': (r) => r.status === 204,
    });
}

export function deleteTask(task_id) {
    const params = {
        responseType: 'none',
        tags: {
            name: `${BASE_URI}/tasks/{taskId}`
        },
    };

    const response = http.del(`${BASE_URI}/tasks/${task_id}`, null, params);

    check(response, {
        'deleteTask is status 204': (r) => r.status === 204,
    });
}
