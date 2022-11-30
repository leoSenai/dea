<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-avatar>
          <q-img
            src="../assets/logo/logoPerhaps.png"
            alt="Logo Clínica Motivar"
            flat
            round
            aria-label="Menu"
            @click.prevent="toggleLeftDrawer"
          />
        </q-avatar>
        <q-toolbar-title> Clínica Motivar </q-toolbar-title>
        <DropDownLogout />
      </q-toolbar>
    </q-header>
    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list >
        <q-item-label header > Profissional </q-item-label>
        <EssentialLink
          v-for="route in essentialLinks"
          :key="route.title"
          v-bind="route"
        />
      </q-list>
    </q-drawer>
    <q-page-container>
      <FloatingActionButton v-if="$q.platform.is.mobile" />
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from 'vue'
import EssentialLink from 'components/EssentialLink.vue'
import FloatingActionButton from 'components/FloatingActionButton.vue'
import DropDownLogout from 'components/DropDownLogout.vue'

const linksList = [
  {
    title: 'Home',
    icon: 'home',
    route: { name: 'IndexPage' }
  },
  {
    title: 'Cadastro de paciente',
    icon: 'mode_edit',
    route: { name: 'CadastroPacientes' }
  },
  {
    title: 'Cadastro pessoas próximas',
    icon: 'mode_edit',
    route: { name: 'ListaPessoasProximas' }
  },
  {
    title: 'Cadastro de questionário',
    icon: 'mode_edit',
    route: { name: 'CadastroQuestionario' }
  },
  {
    title: 'Cadastro de Conselho',
    icon: 'mode_edit',
    route: { name: 'ListaConselho' }
  }

]

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink,
    FloatingActionButton,
    DropDownLogout
  },

  setup () {
    const leftDrawerOpen = ref(false)
    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  }
})
</script>
