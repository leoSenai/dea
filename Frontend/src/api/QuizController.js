import axios from 'axios';

const appUrl = import.meta.env.VITE_API_URL_DEV + '/quiz/'

export default {
  getAll() {
    return axios.get(appUrl + 'get-all')
  },
  getById(id) {
    return axios.get(appUrl + 'get-by-id/' + id)
  },
  insert(data) {
    return axios.post(appUrl + 'insert', data)
  },
  update(data) {
    return axios.post(appUrl + 'update', data)
  }
}
