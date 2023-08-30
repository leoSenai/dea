import DoctorController from './DoctorController'
import QuizController from './QuizController'
import QuestionController from './QuestionController'
import AuthController from './AuthController'
export default {
    install: (app, options) => {
      app.config.globalProperties.$api = {
        DoctorController,
        AuthController,
        QuizController,
        QuestionController,
      }
      app.provide('api', options)
    },
}