const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        name: 'home',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: '',
        name: 'CadastroPaciente',
        component: () => import('pages/CadastroPaciente.vue')
      },
      {
        path: '',
        name: 'CadastroPessoasProximas',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: '',
        name: 'CadastroQuestionario',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: '',
        name: 'Pacientes',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: '',
        name: 'Questionarios',
        component: () => import('pages/IndexPage.vue')
      }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
