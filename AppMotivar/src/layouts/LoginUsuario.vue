<template>
    <img src="../assets/img/wave.svg" class="wave" alt="login-wave">
    <div class="row" style="height: 98vh">
      <div class="col-0 col-md-6 flex justify-center content-center">
        <img src="../assets/img/login.svg" class="responsive" alt="login-image">
      </div>
      <div v-bind:class="{'justify-center': $q.screen.md || $q.screen.sm || $q-screen.xs}"
            class="col-12 col-md-6 flex content-center">
        <q-card v-bind:style="$q.screen.lt.sm ? {'width':'80%'} : {'width':'50%'}">
          <q-card-section>
            <q-avatar size="100px" class="absolute-center shadow-10">
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
              <q-input label="Usuário" v-model="login.usuario">
              </q-input>
              <q-input label="Senha" type="password" v-model="login.senha">
              </q-input>
              <div>
                <q-btn class="full-width" color="primary" label="Login" type="submit" rounded></q-btn>
                <div class="text-center q-mt-sm q-gutter-lg">
                  <router-link class="text-black" to="/">Esqueceu a senha?</router-link>
                </div>
              </div>
            </q-form>
          </q-card-section>
        </q-card>
      </div>
    </div>
</template>
<script>
import { useQuasar } from 'quasar'
let $q
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'LoginUsuario',
  data () {
    return {
      login: {
        usuario: '',
        senha: ''
      }
    }
  },
  methods: {
    submitForm () {
      if (!this.login.usuario || !this.login.senha) {
        $q.notify({
          type: 'negative',
          message: 'Os dados informados são inválidos!'
        })
      } else if (this.login.senha.length < 6) {
        $q.notify({
          type: 'negative',
          message: 'A senha deve ter 6 ou mais caracteres!'
        })
      } else {
        console.log('login')
      }
    }
  },
  mounted () {
    $q = useQuasar()
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

</style>
