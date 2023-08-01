import DoctorController from './DoctorController'

export default {
    install: (app, options) => {
      app.config.globalProperties.$api = {
        DoctorController
      }
      app.provide('api', options)
    },
}