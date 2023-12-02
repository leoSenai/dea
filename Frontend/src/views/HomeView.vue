<template>
  <div class="home">
    <div class="home-options">
      <template
        v-for="link in linksMenu"
        :key="link.path"
      >
        <router-link
          :to="link.path"
          class="link"
          tabindex="0"
          @blur="hideSidebar"
        >
          {{ link.name }}
          <component
            :is="link.icon"
            class="link-icon"
          />
        </router-link>
      </template>
    </div>
  </div>
</template>
<script>
import {
  PhBookOpen,
  PhPerson,
  PhScooter,
  PhArticle,
  PhUsers,
  PhUserList,
} from '@phosphor-icons/vue';
import Cookie from '../utils/cookie';
import { RoleEnum } from '../utils/Enum';

export default {
  components: {
    PhScooter,
    PhBookOpen,
    PhPerson,
    PhArticle,
    PhUsers,
    PhUserList,
  },
  data() {
    return {
      userType: '',
      linksData: [],
    };
  },
  computed: {
    linksMenu() {
      return this.getLinks();
    },
  },
  mounted() {
    this.getData();
    this.fixScreenSize();
  },
  methods: {
    getLinks() {
      const authToken = Cookie.get('authToken');
      const typeUser = Cookie.getUserType(authToken);
      const userId = Cookie.getUserId(authToken);

      switch (typeUser) {
        case RoleEnum.Patient: {
          return [
            {
              path: '/paciente/' + userId + '/questionarios',
              name: 'Questionários',
              icon: PhArticle,
            },
          ];
        }
        case RoleEnum.Administrator: {
          return [{ path: '/usuarios', name: 'Médicos', icon: PhUsers }];
        }
        case RoleEnum.User: {
          return [
            { path: '/questionarios', name: 'Questionários', icon: PhArticle },
            { path: '/pacientes', name: 'Pacientes', icon: PhUserList },
          ];
        }
        case RoleEnum.Person: {
          return [
            {
              path: '/pessoas-proximas/' + userId + '/questionarios',
              name: 'Questionários',
              icon: PhArticle,
            },
          ];
        }
      }
    },
    goLinkMenu(linktag) {
      this.$router.push(linktag);
    },
    getData() {
      this.userType = Cookie.getUserType(Cookie.get('authToken'));
      this.linksData = this.links;
    },
    fixScreenSize() {
      try {
        const contentLoginScreen =
          document.body.getElementsByClassName('login-screen');
        if (contentLoginScreen.length != 0) {
          contentLoginScreen[0].classList.remove('login-screen');
        }
      } catch (err) {
        console.error(err);
      }
    },
  },
};
</script>
<style scoped>
.home {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url(../assets/imgs/home-background.svg) no-repeat;
  background-position: 90% 80%;
  background-size: 30rem;
}

.home-options {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 30rem;
  border-radius: 15px;
  background: var(--primary-700);
  padding: 3rem 1rem;
}

.link {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  padding: 1rem;
  border-radius: 4px;
  border: 1px solid white;
  background: var(--primary-700);
  cursor: pointer;
  transition: 0.2s;
  color: white;
}

.link:hover {
  filter: brightness(0.8);
}
</style>
