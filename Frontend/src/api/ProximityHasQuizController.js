import axios from './axios';

export default {
    url: '/proximity/',
    get(data) {
        return axios.get(this.url + 'get', data)
    },
    getByIdQuiz(id) {
        return axios.get(this.url + 'get-by-id-quiz/' + id)
    },
    getByIdPatient(id) {
        return axios.get(this.url + 'get-by-id-patient/' + id)
    },
    getByIdPerson(id) {
        return axios.get(this.url + 'get-by-id-person/' + id)
    },
    insert(data) {
        return axios.insert(this.url + 'insert', data)
    },
    update(data) {
        return axios.put(this.url + 'update', data)
    },
}
