import axios from 'axios';

export default {
    url: 'user/',
    getAll() {
        return axios.get(this.url + 'get-all')
    },
    insert(data) {
        return axios.post(this.url + 'insert', data)
    },
    update(data) {
        return axios.put(this.url + 'update', data)
    },
    resetPassword(id) {
        return axios.put(this.url + 'reset-password', { 'IdUser': id })
    }
}