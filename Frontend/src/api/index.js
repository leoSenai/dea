import DoctorController from './DoctorController'
import QuizController from './QuizController'

export default {
  install: (app, options) => {
    app.config.globalProperties.$api = {
      DoctorController,
      QuizController
    }
    app.provide('api', options)
  },
}