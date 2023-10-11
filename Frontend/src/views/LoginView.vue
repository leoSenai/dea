<template>
  <div class="row login-container">
    <div :class="[pageSize === 'xs' ? 'col-12' : 'col-5']">
      <div class="login">
        <div class="login-logo">
          <img
            :src="logo"
            alt=""
          >
        </div>
        <h3 class="login-title">
          Clínica Motivar
        </h3>
        <form
          :class="pageSize === 'xl' || pageSize === 'lg' || pageSize === 'md' ? 'w-70' : ''"
          @submit.prevent="submitForm"
        >
          <div class="login-form">
            <InputTemplate
              v-model="model.user"
              color="primary"
              outlined
              dark
              label="Email ou Telefone"
              class="q-mt-md"
            >
              <template #before-label>
                <PhUser class="icon-color" />
              </template>
            </InputTemplate>
            <InputTemplate
              v-model="model.password"
              color="primary"
              outlined
              dark
              label="Senha"
              type="password"
              class="q-mt-md"
            >
              <template #before-label>
                <PhLock class="icon-color" />
              </template>
            </InputTemplate>
          </div>
          <div class="login-button">
            <ButtonTemplate
              class="q-mt-lg btn-login heading-6"
              type="submit"
            >
              Entrar
            </ButtonTemplate>
          </div>
        </form>
      </div>
    </div>
    <div
      v-if="pageSize !== 'xs'"
      class="col-7"
    >
      <div class="login-background" />
    </div>
  </div>
</template>

<script>
import Cookie from '../cookie'
import logo from '../assets/imgs/logo.png';
import InputTemplate from '../components/InputPrimary.vue';
import ButtonTemplate from '../components/ButtonPrimary.vue';
import { PhUser } from '@phosphor-icons/vue';
import { PhLock } from '@phosphor-icons/vue';

export default {
  components: {
    InputTemplate,
    ButtonTemplate,
    PhUser,
    PhLock,
  },
  data() {
    return {
      model: {
        user: '',
        password: ''
      },
      logo: logo,
    };
  },
  mounted(){
    if(!document.getElementsByClassName('content')[0].classList.contains('login-screen')){
      document.getElementsByClassName('content')[0].classList.add('login-screen')
    }
  },
  computed: {
    pageSize() {
      return this.$q.screen.name
    }
  },
  methods: {
    submitForm() {
      const th = this;
      th.$api.AuthController.login(th.model).then(({ data }) => {
        if (data) {
          Cookie.set({ name: 'authToken', value: data.data })
          this.$router.push('/')
        }
      })
    }
  },

};
</script>

<style scoped>
.w-70 {
  width: 70%;
}

.login-container {
  width: 100%;
  max-height: 100vh;
  overflow: hidden;
}

.login {
  height: 100vh;
  background: linear-gradient(180deg, #386923 35%, #163408 100%);
  display: flex;
  align-items: center;
  flex-direction: column;
  padding: 3rem;
}

.login-form {
  width: 100%;
}

.login-logo>img {
  width: 6.25rem;
}

.login-title {
  font-size: 2rem;
  color: #fff !important;
}

.btn-login {
  width: 100%;
  font-size: 1rem;
  font-weight: 500;
  color: #fff !important
}

.login-button {
  width: 100%;
}

.login-background {
  background-size: contain;
  background-repeat: no-repeat;
  height: 100%;
  width: 100%;
}

.login-background {
  background: url('../assets/imgs/login-background.png') no-repeat;
  background-size: cover;
}

.icon-color {
  color: var(--neutral-white);
}
</style>