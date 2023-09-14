import axios from './axios';

export default {
  url: '/question/',
  config: {
    headers: {
      Authorization: `; ${document.cookie}`.split('; authToken=').pop().split(';').shift()
    }
  },
  getAll() {
    return axios.get(this.url + 'get-all', this.config)
  },
  getById(id) {
    return axios.get(this.url + 'get-by-id/' + id, this.config)
  },
  insert(data) {
    return axios.post(this.url + 'insert', data, this.config)
  },
  update(data) {
    return axios.put(this.url + 'update', data, this.config)
  },
  delete(id) {
    return axios.delete(this.url + id, this.config)
  }
}
