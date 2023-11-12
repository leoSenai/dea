import axios from './axios';

export default {
    url: '/proximityhasquiz/',
    getAll(data) {
        return axios.get(this.url + 'getAll', data)
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
    getByIdQuizPerson(idquiz, idperson) {
        return axios.get(this.url + 'get-by-id-quiz-person/' + idquiz + '/' + idperson) 
    },
}
