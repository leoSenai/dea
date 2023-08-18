import DoctorController from './DoctorController'
import QuizController from './QuizController'
import QuestionController from './QuestionController'

export default {
  install: (app, options) => {
    app.config.globalProperties.$api = {
      DoctorController,
      QuizController,
      QuestionController,
    }
    app.provide('api', options)
  },
}