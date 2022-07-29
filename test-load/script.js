import http from 'k6/http';

export const BASE_URI = 'http://app:8080'
export const VIRTUAL_USERS = 60
export const ITERATIONS = 60

export const options = {
    scenarios: {
        get_tasks: {
            executor: 'per-vu-iterations',
            exec: 'getTasks',
            vus: VIRTUAL_USERS,
            iterations: ITERATIONS,
        },
        add_task: {
            executor: 'per-vu-iterations',
            exec: 'addTask',
            vus: VIRTUAL_USERS,
            iterations: ITERATIONS,
        },
        get_task: {
            executor: 'per-vu-iterations',
            exec: 'getTask',
            vus: VIRTUAL_USERS,
            iterations: ITERATIONS,
        },
        replace_task: {
            executor: 'per-vu-iterations',
            exec: 'replaceTask',
            vus: VIRTUAL_USERS,
            iterations: ITERATIONS,
        },
        delete_task: {
            executor: 'per-vu-iterations',
            exec: 'deleteTask',
            vus: VIRTUAL_USERS,
            iterations: ITERATIONS,
        },
    },
};

export function getTasks() {
    http.get(BASE_URI + '/tasks');
}

export function addTask() {
    const payload = JSON.stringify({
        title: 'Lorem ipsum dolor',
        description: 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed dia',
    });

    const response = http.post(BASE_URI + '/tasks', payload);

    return response.json()
}

export function getTask() {
    let uuid = addTask().uuid

    const response = http.get(BASE_URI + '/tasks/' + uuid);

    return response.json()
}

export function replaceTask() {
    let uuid = getTask().uuid

    const payload = JSON.stringify({
        title: 'Lorem ipsum dolor',
        description: 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed dia',
    });

    http.put(BASE_URI + '/tasks/' + uuid, payload);

    return uuid
}

export function deleteTask() {
    let uuid = replaceTask()

    http.del(BASE_URI + '/tasks/' + uuid);
}
