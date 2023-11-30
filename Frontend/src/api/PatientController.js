import axios from './axios';

const appUrl = import.meta.env.VITE_API_URL_DEV + 'patient/'

export default {
  getAll() {
    return axios.get(appUrl + 'get-all')
  },
  getById(id) {
    return axios.get(appUrl + 'get-by-id/' + id)
  },
  getByDoc(doc) {
    return axios.get(appUrl + 'get-by-doc/' + doc)
  },
  insert(data) {
    return axios.post(appUrl + 'insert', data)
  },
  update(data) {
    return axios.put(appUrl + 'update', data)
  },
  resetPassword(id) {
    return axios.put(appUrl + 'reset-password', { 'patient_id': id })
  }
}
