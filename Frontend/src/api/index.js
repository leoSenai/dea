import DoctorController from './DoctorController'
import QuizController from './QuizController'
import QuestionController from './QuestionController'
import AuthController from './AuthController'
import Cookie from '../cookie'
export default {
  install: (app, options) => {
    app.config.globalProperties.$api = {
      DoctorController,
      AuthController,
      QuizController,
      QuestionController,
    }

    // Refresh token every 30min
    setInterval(() => {
      if (Cookie.get('authToken')) {
        AuthController.refreshToken()
      }
    }, 30 * 60 * 1000);

    app.provide('api', options)
  },
}