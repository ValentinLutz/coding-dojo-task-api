import http from 'k6/http';

export const BASE_URI = 'http://localhost:8080'
export const VIRTUAL_USERS = 1000
export const ITERATIONS = 10

export const options = {
    scenarios: {
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
    let taskId = addTask().task_id

    const response = http.get(BASE_URI + '/tasks/' + taskId);

    return response.json()
}

export function replaceTask() {
    let taskId = getTask().task_id

    const payload = JSON.stringify({
        title: 'Lorem ipsum dolor',
        description: 'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed dia',
    });

    http.put(BASE_URI + '/tasks/' + taskId, payload);

    return taskId
}

export function deleteTask() {
    let taskId = replaceTask()

    http.del(BASE_URI + '/tasks/' + taskId);
}
