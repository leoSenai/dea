import axios from './axios';
import Cookie from '../cookie';

export default {
    url: '/auth/',
    login(data) {
        return axios.post(this.url + 'login', data)
    },
    logout() {
        Cookie.delete('authToken')
    },
    parseJwt(tokenStr) {
        const base64Url = tokenStr.split('.')[1];
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function (c) {
          return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        }).join(''));
        return JSON.parse(jsonPayload);
    }
}
