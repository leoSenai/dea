import axios from './axios';

export default {
    url: '/auth/',
    config: {
        headers: {
            Authorization: `; ${document.cookie}`.split('; authToken=').pop().split(';').shift()
        }
    },
    login(data) {
        return axios.post(this.url + 'login', data, this.config)
    },
    logout() {
        return 'not success'
    }
}
