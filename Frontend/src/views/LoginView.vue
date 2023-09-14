<template>
  <div>
    <div class="row">
      <div class="col-12 col-md-5">
        <div class="login flex items-center column justify-center q-px-lg">
          <div class="login-logo">
            <img
              :src="logo"
              alt=""
            />
          </div>
          <div class="login-title">
            <h2>Clínica Motivar</h2>
          </div>
          <div class="login-container">
            <div class="login-form">
              <InputTemplate
                color="primary"
                outlined
                dark
                v-model="text"
                label="Email ou Telefone"
                class="login-input q-mt-md"
                ><template v-slot:before-label
                  ><PhUser class="icon-color"></PhUser></template>
              ></InputTemplate>
              <InputTemplate
                color="primary"
                outlined
                dark
                v-model="text"
                label="Senha"
                class="login-input q-mt-md q-field__inner"
                ><template v-slot:before-label
                  ><PhLock class="icon-color"></PhLock></template
              ></InputTemplate>
            </div>
            <div class="login-button">
              <ButtonTemplate class="q-mt-lg btn-login">Entrar</ButtonTemplate>
            </div>
          </div>
        </div>
      </div>
      <div class="col-md-7">
        <div class="login-img">
          <div class="bg-img"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import logo from '../assets/imgs/Logo.png';
import imagem_principal from '../assets/imgs/imagem-principal.png';
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
      formInfo: {
        user: '',
        password: ''
      },
      logo: logo,
      imagem_principal: imagem_principal,
    };
  },
  methods: {
    submitForm() {
      const th = this;
      th.$api.AuthController.login(th.formInfo)
    }
  }
};
</script>

<style scoped>
.login {
  height: 100vh;
  background: #386923;
  background: linear-gradient(180deg, #386923 35%, #163408 100%);
}

.login-form {
  width: 100%;
}

.login-logo>img {
  width: 110px;
  height: 98px;
}

.login-title>h2 {
  font-size: 38px;
}

.login-container {
  width: 80%;
}

@media screen and (max-width: 992px) {
  .login-container {
    width: 75% !important;
  }
}

@media screen and (max-width: 768px) {
  .login-container {
    width: 100% !important;
  }
}

.btn-login {
  width: 100% !important;
  font-size: 18px !important;
  font-weight: bold;
}

.login-button {
  width: 100% !important;
}

.login-img {
  background-image: url("../assets/imgs/imagem-principal.png");
  background-size: cover;
  background-position-x: -50px;
  background-repeat: no-repeat;
  height: 100%;
  width: 100%;
}

.bg-img {
  background-color: rgba(0, 0, 0, 0.324);
  width: 100%;
  height: 100%;
}

.icon-color {
  color: var(--neutral-white);
}</style>