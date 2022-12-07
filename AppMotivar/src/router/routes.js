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
        path: 'CadastroPacientes/:id?',
        name: 'CadastroPacientes',
        component: () => import('pages/Cadastros/CadastroPacientes.vue')
      },
      {
        path: 'CadastroQuestionario/:id?',
        name: 'CadastroQuestionario',
        component: () => import('src/pages/Cadastros/CadastroQuestionario.vue')
      },
      {
        path: 'CadastroPessoasProximas/:id?',
        name: 'CadastroPessoasProximas',
        component: () => import('src/pages/Cadastros/CadastroPessoasProximas.vue')
      },
      {
        path: 'CadastroConselho/:id?',
        name: 'CadastroConselho',
        component: () => import('src/pages/Cadastros/CadastroConselho.vue')
      },
      {
        path: 'ListaQuestionario',
        name: 'ListaQuestionario',
        component: () => import('src/pages/Cadastros/ListaQuestionario.vue')
      },
      {
        path: 'ListaPessoasProximas',
        name: 'ListaPessoasProximas',
        component: () => import('src/pages/Cadastros/ListaPessoasProximas.vue')
      },
      {
        path: 'ListaConselho',
        name: 'ListaConselho',
        component: () => import('src/pages/Cadastros/ListaConselho.vue')
      },
      {
        path: 'ListaPacientes',
        name: 'ListaPacientes',
        component: () => import('src/pages/Cadastros/ListaPacientes.vue')
      }
    ]
  },
  {
    path: '/LoginUsuario',
    name: 'LoginUsuario',
    component: () => import('layouts/LoginUsuario.vue')
  },
  {
    path: '/LoginEsqueceSenha',
    name: 'LoginEsqueceSenha',
    component: () => import('layouts/LoginEsqueceSenha.vue')
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
