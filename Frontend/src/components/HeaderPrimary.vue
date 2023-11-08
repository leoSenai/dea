<template>
  <div>
    <header class="bg-primary-700">
      <button
        class="hamburguer"
        @click="toggleSidebar"
      >
        <PhList size="1.4em" />
      </button>
      <div
        class="logo white"
        @click="goHome"
      >
        <q-img :src="LogoSrc" />
        <h6 class="header-text">
          Clínica Motivar
        </h6>
      </div>
      <div
        class="user-header"
        @click="openProfileMenu"
      >
        <div :class="['user', 'white', showExitDropDown ? 'user-dropdown-active' : '']">
          <ph-user-circle
            class="user-icon"
            regular
          />
          <span class="text-body"> {{ username }} </span>
          <Transition name="top-to-bottom">
            <div
              v-show="showExitDropDown"
              class="dropdown-exit"
            >
              <button
                type="button"
                @click="logout"
              >
                <PhSignOut
                  weight="bold"
                  size="16"
                />
                Sair
              </button>
            </div>
          </Transition>
        </div>
      </div>
    </header>
    <div
      :class="['sidebar', 'bg-success', isSidebarActive ? 'active' : '']"
      @blur="hideSidebar"
    >
      <template
        v-for="link in links"
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
import { PhUserCircle, PhList, PhDoor, PhSignOut } from '@phosphor-icons/vue';
import ButtonPrimary from '../components/ButtonPrimary.vue';
import cookie from '../utils/cookie';

export default {
  components: {
    PhUserCircle,
    PhList,
    ButtonPrimary,
    PhDoor,
    PhSignOut
  },
  props: {
    links: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      LogoSrc: '/logo.png',
      isSidebarActive: false,
      showExitDropDown: false
    };
  },
  computed: {
    username() {
      return cookie.getAuthUser(cookie.get('authToken'));
    },
  },
  methods: {
    toggleSidebar() {
      this.isSidebarActive = !this.isSidebarActive;
    },
    hideSidebar() {
      this.isSidebarActive = false;
    },
    goHome() {
      this.$router.push('/');
    },
    logout() {
      const th = this;
      th.$api.AuthController.logout();
      this.$router.push('/login');
    },
    openProfileMenu() {
      this.showExitDropDown = !this.showExitDropDown
    },
  },
};
</script>

<style scoped>
.header-text {
  text-align: center;
}

header {
  position: sticky;
  min-width: 100%;
  z-index: 999;
  display: flex;
  align-items: center;
  justify-content: space-between;

  padding: 0.5rem 1rem;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 2rem;
  cursor: pointer;
}

.user {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-icon {
  height: 100%;
  width: 1.5rem;
}

.q-img {
  height: 100%;
  width: 4rem;
}

.sidebar {
  box-shadow: none;
  position: fixed;
  z-index: 999;
  min-height: 100%;
  padding: 2rem 0rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  transition: 0.3s ease-in-out;
  transform: translateX(0%);
}

.link {
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-width: 100%;
  padding: 1rem 2rem;
  gap: 1rem;
  transition: 200ms ease;
  cursor: pointer;
  color: white;
}

.link-icon {
  font-size: 1.25rem;
}

.hamburguer {
  display: flex;
  font-size: 1.5rem;
  background: none;
  border: none;
  cursor: pointer;
  margin-right: 1em;
  color: #fff;
}

.sidebar {
  transform: translateX(-100%);
}

.sidebar.active {
  transform: translateX(0%);
  box-shadow: 4px 3px 12px black;
}

.user-header {
  cursor: pointer;
  position: relative;
}

.user.user-dropdown-active {
  border-radius: 4px 4px 0 0;
}

.dropdown-exit {
  position: absolute;
  bottom: -4.25rem;
  left: 0;
  width: 100%;
  display: flex;
}

.dropdown-exit button {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  background: var(--primary-700);
  border: none;
  padding: .5rem;
  border-radius: 4px;
  cursor: pointer;
  transition: .3s;
}

.dropdown-exit button:hover {
  filter: brightness(0.8);
}
</style>
