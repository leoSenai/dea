import axios from './axios';

export default {
    url: '/proximity/',
    getByIdPerson(id) {
        return axios.get(this.url + 'get-by-id-person/' + id)
    },
    getByIdPatient(id) {
        return axios.get(this.url + 'get-by-id-patient/' + id)
    },
    insert(data) {
        return axios.post(this.url + 'insert', data)
    }
}
