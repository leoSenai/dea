<template>
  <div>
    <div class="row">
      <div class="col-12 col-lg-5">
        <div class="login flex items-center column justify-center q-px-lg">
          <div class="login-logo">
            <img :src="logo" alt="" />
          </div>
          <div class="login-title">
            <h2>Clínica Motivar</h2>
          </div>
          <div class="login-container">
            <div class="login-form">
              <Input
                color="primary"
                outlined
                dark
                v-model="text"
                label="Email ou Telefone"
                class="login-input q-mt-md"
                ><template v-slot:before-label
                  ><PhUser class="icon-color"></PhUser></template
              ></Input>
              <Input
                color="primary"
                outlined
                dark
                v-model="text"
                label="Senha"
                class="login-input q-mt-md q-field__inner"
                ><template v-slot:before-label
                  ><PhLock class="icon-color"></PhLock></template
              ></Input>
            </div>
            <div class="login-button">
              <Button class="q-mt-lg btn-login">Entrar</Button>
            </div>
          </div>
        </div>
      </div>
      <div class="col-lg-7">
        <div class="login-img">
          <div class="bg-img"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import logo from "../assets/imgs/Logo.png";
import imagem_principal from "../assets/imgs/imagem-principal.png";
import Input from "../components/Input.vue";
import Button from "../components/Button.vue";
import { PhUser } from "@phosphor-icons/vue";
import { PhLock } from "@phosphor-icons/vue";
export default {
  components: {
    Input,
    Button,
    PhUser,
    PhLock,
  },
  data() {
    return {
      logo: logo,
      imagem_principal: imagem_principal,
    };
  },
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

.login-logo > img {
  width: 110px;
  height: 98px;
}

.login-title > h2 {
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
}
</style>