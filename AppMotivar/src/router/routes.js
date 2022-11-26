const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        name: 'IndexPage',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: 'IndexPage',
        name: 'IndexPage',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: 'CadastroPacientes',
        name: 'CadastroPacientes',
        component: () => import('pages/Cadastros/CadastroPacientes.vue')
      },
      {
        path: 'CadastroPessoasProximas',
        name: 'CadastroPessoasProximas',
        component: () => import('src/pages/Cadastros/CadastroPessoasProximas.vue')
      },
      {
        path: 'CadastroQuestionario',
        name: 'CadastroQuestionario',
        component: () => import('src/pages/Cadastros/CadastroQuestionario.vue')
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
