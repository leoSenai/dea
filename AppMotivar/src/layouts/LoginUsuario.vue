<template>
   <q-layout >
   <q-page-container>
    <q-page class="flex flex-center" style="background: radial-gradient(circle, #386923 0%, #519832 100%)">

    <q-card class="rounded" style="width: 150vh">
    <img src="../assets/img/mesh-463.png" class="wave" alt="login-wave">
    <div class="row" style="height: 98vh" >
      <div class="col-0 col-md-6 flex fundo-login justify-center content-center">
        <img src="../assets/img/login-user.svg"  class="responsive" alt="login-image">
      </div>
      <div :class="{'justify-center': $q.screen.md || $q.screen.sm || $q.screen.xs  || $q.screen.lg  }"
            class="col-12 col-md-6 flex content-center">
        <q-card :style="{'width': cardStyle}">
          <q-card-section>
            <q-avatar
            size="100px"
            class="absolute-center shadow-10">
              <img src="../assets/img/logoPerhaps.png" alt="logo">
            </q-avatar>
          </q-card-section>
          <q-card-section>
            <div class="q-pt-lg">
              <div class="col text-h6 ellipsis flex justify-center">
                <h3 class="text-h3 text-uppercase q-my-none text-weight-regular">Login</h3>
              </div>
            </div>
          </q-card-section>
          <q-card-section>
            <q-form class="q-gutter-md" @submit.prevent="submitForm">

              <q-input
              label="Usuário"
              outlined
              v-model="form.usuario"
              :rules="[(val) => val.length > 0 || 'Usuário é obrigatório']"/>

              <q-input
              label="Senha"
              outlined
              type="password"
              v-model="form.senha"
              :rules="[(val) => val.length > 6 || 'Senha é obrigatório']"/>

              <div>
                <q-btn class="full-width" color="primary" label="Login" type="submit" :to="{name: 'IndexPage'}" rounded></q-btn>
                <div class="text-center q-mt-sm q-gutter-lg">
                  <router-link class="text-black" to="/LoginEsqueceSenha">Esqueceu a senha?</router-link>
                </div>
              </div>
            </q-form>
          </q-card-section>
        </q-card>
      </div>
    </div>
    </q-card>
  </q-page>
  </q-page-container>
  </q-layout>
</template>
<script>
import { useQuasar } from 'quasar'
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'LoginUsuario',
  setup () {
    const $q = useQuasar()
    const cardStyle = ref($q.screen.lt.sm ? '90%' : '70%')
    const getUsers = async () => {
      // criar validação para login e retirar o 'to' do botão submit
    }
    const submitForm = () => {
      getUsers()
    }
    const form = ref({
      usuario: '',
      senha: ''
    })
    return {
      form,
      submitForm,
      cardStyle
    }
  }
})
</script>

<style scoped>
.wave{
  position: fixed;
  height: 100%;
  left: 0;
  bottom: 0;
  z-index: -1;
}

.responsive{
  width: 95%;
  height: auto;
  z-index: 1;
}

.fundo-login{
  background-color: #3c69349d;
}

</style>
