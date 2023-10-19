import DoctorController from './DoctorController'
import QuizController from './QuizController'
import QuestionController from './QuestionController'
import AuthController from './AuthController'
import UsersController from './UsersController'
import CboController from './CboController'
import ServicesController from './ServicesController'
import ProximityController from './ProximityController'
import ProximityHasQuizController from './ProximityHasQuizController'
import PersonController from './PersonController'

export default {
  install: (app, options) => {
    app.config.globalProperties.$api = {
      DoctorController,
      AuthController,
      QuizController,
      QuestionController,
      UsersController,
      CboController,
      ServicesController,
      ProximityController,
      ProximityHasQuizController,
      PersonController
    }
    app.provide('api', options)
  },
}