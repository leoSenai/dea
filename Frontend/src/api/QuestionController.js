import axios from './axios';

export default {
  url: '/question/',
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
    return axios.put(this.url + 'update', data)
  },
  delete(id) {
    return axios.delete(this.url + id)
  }
}
