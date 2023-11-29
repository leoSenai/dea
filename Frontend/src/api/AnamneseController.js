import axios from './axios';

const appUrl = import.meta.env.VITE_API_URL_DEV + 'anamnese/'

export default {
  getAll() {
    return axios.get(appUrl + 'get-all')
  },
  getById(id) {
    return axios.get(appUrl + 'get-by-id/' + id)
  },
  getByIdUserPatient(data){
    return axios.get(appUrl + 'get-by-id-user-patient/'+ data['IdUser'] + '/'+data['IdPatient'])
  },
  insert(data) {
    return axios.post(appUrl + 'insert', data)
  },
  update(data) {
    return axios.put(appUrl + 'update', data)
  },
  getLaudo(idUser, idPatient, grau){
    return axios.get(appUrl + 'report/'+idUser+'/'+idPatient+'/'+grau)
  },
}
