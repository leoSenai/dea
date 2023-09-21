import AuthController from '../api/AuthController'

export default {
  get(name) {
    return `; ${document.cookie}`.split(`; ${name}=`).pop().split(';').shift()
  },
  set({ name, value, isDeleting = false }) {
    let cookie = ''
    let expireAt = '';
    if (isDeleting) {
      expireAt = '-1'
    } else {
      const date = new Date();
      date.setTime(date.getTime() + (24 * 60 * 60 * 1000));
      expireAt = `expires=${date.toUTCString()}`;
    }
    cookie = `${name}=${value};expires=${expireAt}`
    document.cookie = cookie;
  },
  delete(name) {
    this.set({ name, value: '', isDeleting: true })
  },
  getAuthUser(tokenStr){
    return AuthController.parseJwt(tokenStr)['Username']
  }
}