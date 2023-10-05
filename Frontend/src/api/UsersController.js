import axios from 'axios';

export default {
    url: 'user/',
    getAll() {
        return axios.get(this.url + 'get-all')
    }
}