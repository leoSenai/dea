import axios from 'axios';

const appUrl = import.meta.env.VITE_API_URL_DEV + '/auth/'

export default {
    login (formInfo) {
        console.log(formInfo)
        return axios.get(appUrl + 'login')
    },
    logout () {
        return axios.get(appUrl + 'logout')
    }
}
