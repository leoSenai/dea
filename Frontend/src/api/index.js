import DoctorController from './DoctorController'
import QuizController from './QuizController'
import QuestionController from './QuestionController'
import AuthController from './AuthController'
import PatientController from './PatientController'
export default {
    install: (app, options) => {
      app.config.globalProperties.$api = {
        DoctorController,
        AuthController,
        QuizController,
        QuestionController,
        PatientController,
      }
      app.provide('api', options)
    },
}