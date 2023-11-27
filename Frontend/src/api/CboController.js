import axios from './axios';

export default {
    url: '/cbo/',
    getAll() {
      return axios.get(this.url + 'get-all')
    },
    update(data) {
      return axios.put(this.url + 'update', data)
    },
}