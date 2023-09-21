import axios from './axios';

export default {
    url: '/auth/',
    login(data) {
        return axios.post(this.url + 'login', data)
    },
    logout() {
        return 'not success'
    },
    refreshToken() {
        return axios.get(this.url + 'refresh-token')
    }
}
