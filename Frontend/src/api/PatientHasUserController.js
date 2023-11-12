import axios from './axios';

export default {
    url: '/patienthasuser/',
    getByIdUser(id) {
        return axios.get(this.url + 'get-by-id-user/' + id)
    },
    insert(data) {
        return axios.post(this.url + 'insert', data)
    }
}
