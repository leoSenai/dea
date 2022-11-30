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
      },
      {
        path: 'ListaConselho',
        name: 'ListaConselho',
        component: () => import('src/pages/Cadastros/ListaConselho.vue')
      },
      {
        path: 'CadastroConselho/:id?',
        name: 'CadastroConselho',
        component: () => import('src/pages/Cadastros/CadastroConselho.vue')
      }
    ]
  },
  {
    path: '/LoginUsuario',
    name: 'LoginUsuario',
    component: () => import('layouts/LoginUsuario.vue')
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
