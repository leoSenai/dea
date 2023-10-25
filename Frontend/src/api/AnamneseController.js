import axios from 'axios';

export default {
    url: 'user/',
    getByUserIdAndPatientId() {
        return axios.get(this.url + 'get-all')
    }
}