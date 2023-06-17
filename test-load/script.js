import {randomString} from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import http from 'k6/http';
import {check} from "k6";

export const BASE_URI = 'http://localhost:8080'

export const options = {
    thresholds: {
        http_req_duration: ['p(90) < 400', 'p(95) < 600'],
    },
    scenarios: {
        // full_scenario_shared: {
        //     executor: 'shared-iterations',
        //     vus: 10,
        //     exec: 'fullScenario',
        //     iterations: 2000,
        // },
        // full_scenario_constant: {
        //     executor: 'constant-vus',
        //     exec: 'fullScenario',
        //     vus: 400,
        //     duration: '30s',
        // },
        full_scenario: {
            executor: 'ramping-vus',
            exec: 'fullScenario',
            startVUs: 0,
            stages: [
                {duration: '1m', target: 100},
                {duration: '1m', target: 0},
            ],
            gracefulRampDown: '20s',
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
    const response = http.get(`${BASE_URI}/tasks`);

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
            name: `${BASE_URI}/tasks/{taskId}`
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
        tags: {
            name: `${BASE_URI}/tasks/{taskId}`
        },
    };

    const response = http.del(`${BASE_URI}/tasks/${task_id}`, null, params);

    check(response, {
        'deleteTask is status 204': (r) => r.status === 204,
    });
}
