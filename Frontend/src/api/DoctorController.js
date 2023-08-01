import axios from 'axios';

const appUrl = import.meta.env.VITE_API_URL_DEV + '/doctor/'

export default {
    getAll () {
        return axios.get(appUrl + 'get-all')
    },
    insert () {
        return axios.post(appUrl + 'insert', { Name: 'Felipe', Crm: '1243', IdCbo: 1 })
    }
}
