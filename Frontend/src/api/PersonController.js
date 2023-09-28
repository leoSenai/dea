import axios from './axios';

export default {
  url: '/person/',
  getAll() {
    return axios.get(this.url + 'get-all')
  },
  getById(id) {
    return axios.get(this.url + 'get-by-id/' + id)
  },
  insert(data) {
    return axios.post(this.url + 'insert', data)
  },
  update(data) {
    return axios.post(this.url + 'update', data)
  }
}
