import axios from './axios';

const appUrl = import.meta.env.VITE_API_URL_DEV + 'auth/'

export default {
    login(formInfo) {
        return axios.post(appUrl + 'login', formInfo)
    },
    logout() {
        return 'not success'
    }
}
