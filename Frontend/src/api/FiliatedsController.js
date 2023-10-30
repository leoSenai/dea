import axios from './axios';

export default {
    url: '/filiateds/',
    update(data) {
        return axios.put(this.url + 'update', data)
    },
}
