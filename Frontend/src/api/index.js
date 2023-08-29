import DoctorController from './DoctorController'
import AuthController from './AuthController'

export default {
    install: (app, options) => {
      app.config.globalProperties.$api = {
        DoctorController,
        AuthController
      }
      app.provide('api', options)
    },
}