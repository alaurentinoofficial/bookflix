import axios from '@/repositories/axios.js';

const resource = '/book';

export default {
    getAll() {
        return axios.get(`${resource}`);
    },
    getById(id) {
        return axios.get(`${resource}/${id}`);
    },
    add(payload) {
        return axios.post(`${resource}`, payload);
    },
    update(payload, id) {
        return axios.put(`${resource}/${id}`, payload);
    },
    delete(id) {
        return axios.delete(`${resource}/${id}`)
    },

    // MANY OTHER ENDPOINT RELATED STUFFS
};