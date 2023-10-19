<template>
  <div class="home">
    <div class="home-options">
      <ul>
        <li class="menu-item" @click="goLinkMenu('pacientes')">Pacientes<PhScooter /></li>
        <li class="menu-item" @click="goLinkMenu('questionarios')">Questionários<PhBookOpen /></li>
        <li class="menu-item" @click="goLinkMenu('usuarios')">Usuários<PhPerson /></li>
      </ul>
    </div>
  </div>
</template>
<script>
import { PhBookOpen, PhPerson, PhScooter } from '@phosphor-icons/vue';

export default {
  components: { PhScooter, PhBookOpen, PhPerson },
    data() {
        return {};
    },
    mounted(){
      try{
        document.getElementsByClassName('content')[0].classList.remove('login-screen')
      }catch{
        //do nothing
      }
    },
    methods: {
      goLinkMenu(linktag){
        this.$router.push('/'+linktag)
      }
    },
};
</script>
<style scoped>

  .menu-item:hover{
      cursor: pointer;
      opacity: 70%;
  }

  .home{
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .home-options{
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    background-color: var(--primary-700);
    border-radius: 50px;
    width: 30em;
    height: 30em;
    display: flex;
    opacity:  100%;
    justify-content: center;
    align-items: self-end;
  }

  .home-options ul{
    width: 85%;
    height: 85%;
  }

  
  .home-options ul li {
    width: auto;
    color: green;
    background-color: #f5f5f5;
    height:20%;
    margin: 20px;
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 20px;
  }

  @media (max-width: 500px) {

    .home-options{
      background-color: transparent;
      box-shadow: none;
    }
    .home-options ul li {
      font-size: 75%;
      box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    } 
  } 
</style>
