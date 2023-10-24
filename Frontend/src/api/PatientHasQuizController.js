import axios from './axios';

export default {
    url: '/patienthasquiz/',
    getByIdQuiz(id) {
        return axios.get(this.url + 'get-by-id-quiz/' + id)
    },
    getByIdPatient(id) {
        return axios.get(this.url + 'get-by-id-patient/' + id)
    },
    insert(data) {
        return axios.insert(this.url + 'insert', data)
    },
    update(data) {
        return axios.put(this.url + 'update', data)
    },
}
