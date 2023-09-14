import axios from './axios';

export default {
    url: '/doctor/',
    config: {
        headers: {
            Authorization: `; ${document.cookie}`.split('; authToken=').pop().split(';').shift()
        }
    },
    getAll() {
        return axios.get(this.url + 'get-all', this.config)
    },
    insert() {
        return axios.post(this.url + 'insert', { Name: 'Felipe', Crm: '1243', IdCbo: 1 }, this.config)
    }
}
