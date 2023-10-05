import axios from './axios';

export default {
    url: '/services/',
    getAll() {
      return axios.get(this.url + 'get-all')
    },
}