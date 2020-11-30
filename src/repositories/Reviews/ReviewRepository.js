import axios from '@/repositories/axios.js';
import authHeader from '@/services/authHeader';

const resource = '/book/{bookid}/review';

export default {
    getById(bookId, id) {
        console.log(id)
        return axios.get(`${resource.replace('{bookid}', bookId)}/${id}`);
    },
    getAllByBookId(bookId) {
        return axios.get(`${resource.replace('{bookid}', bookId)}`);
    },
    add(payload, bookId) {
        return axios.post(`${resource.replace('{bookid}', bookId)}`, payload, { headers: authHeader() });
    },
    update(payload, id) {
        return axios.put(`${resource}/${id}`, payload);
    },
    delete(bookId, id) {
        return axios.delete(`${resource.replace('{bookid}', bookId)}/${id}`)
    }
};