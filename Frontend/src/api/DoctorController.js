import axios from './axios';

export default {
    url: '/doctor/',
    getAll() {
        return axios.get(this.url + 'get-all')
    },
    insert() {
        return axios.post(this.url + 'insert', { Name: 'Felipe', Crm: '1243', IdCbo: 1 })
    }
}
