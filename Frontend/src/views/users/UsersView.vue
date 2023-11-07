<template>
  <div class="user-content">
    <div class="user-title">
      <div class="title">
        <h3>Usuários</h3>
      </div>
      <div
        v-if="!isMobile"
        class="title-add-user"
      >
        <button
          type="button"
          @click="openAddEditModal()"
        >
          Adicionar
          <PhPlus color="white" />
        </button>
      </div>
    </div>
    <div
      v-if="model.hasError"
      class="error user"
    >
      {{ model.message }}
    </div>
    <div v-if="model.data.length==0">
      <i>Não há usuários cadastrados até o momento.</i>
    </div>
    <div
      v-for="user in model.data"
      v-else
      :key="user.Id"
      class="row user"
    >
      <p @click="openViewModal(user)">
        {{ user.Name }}
      </p>
      <div class="user-actions">
        <button
          type="button"
          @click="openViewModal(user)"
        >
          <PhEye color="black" />
        </button>
        <button
          type="button"
          @click="openAddEditModal(user)"
        >
          <PhPencil color="black" />
        </button>
      </div>
    </div>
    <div
      v-if="isMobile"
      class="add-user"
    >
      <button
        type="button"
        @click="openAddEditModal()"
      >
        <PhPlus color="white" />
      </button>
    </div>
    <UsersAddEditModal
      ref="addEdit"
      @close="load"
    />
    <ViewUserModal
      ref="viewUser"
      @close="load"
    />
  </div>
</template>
<script>
import { PhPlus, PhPencil, PhEye } from '@phosphor-icons/vue';
import UsersAddEditModal from './UsersAddEditModal.vue';
import ViewUserModal from './UserViewModal.vue'

export default {
  components: {
    PhPlus,
    ViewUserModal,
    PhEye,
    PhPencil,
    UsersAddEditModal
  },
  data() {
    return {
      model: {
        data: [],
        hasError: false,
        message: ''
      },
    }
  },
  computed: {
    isMobile() {
      return this.$q.screen.xs || this.$q.screen.sm
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    openViewModal(currentUser){
      this.$refs.viewUser.openModal(currentUser)
    },
    load() {
      const th = this;
      th.$api.UsersController.getAll().then(({ data }) => {
        th.model = data
      }).catch(({ response }) => {
        th.model = {
          ...response.data,
          hasError: true
        }
      })
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current)
    }
  },
}
</script>
<style scoped>
.user-content {
  padding: 3rem 1.5rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
}

.user-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.row.user{
  padding: 0;
  width: 100%;
  display: flex;
  cursor: pointer;
}

.row.user p{
  padding: 1em;
  width: -webkit-fill-available;
}

.row.user .user-actions{
  position: auto;
  display: flex;
  padding-right: 1em;
  width: auto;
  height: auto;
  margin-left: -60%;
}

.title-add-user button {
  background: var(--primary);
  border-radius: 8px;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: .2s;
  gap: .5rem;
  padding: 1rem 2rem;
  color: white;
  cursor: pointer;
}

.title-add-user button:hover {
  filter: brightness(0.8);
}

.error {
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
}

.add-user {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.add-user button {
  background: var(--primary);
  border-radius: 99999px;
  font-size: 2rem;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 3rem;
  height: 3rem;
  transition: .2s;
}

.add-user button:hover {
  filter: brightness(0.8);
  cursor: pointer;
}

.user {
  border: 1px solid var(--neutral-dark-gray);
  color: var(--neutral-dark-gray);
  padding: 0.8rem;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.user p {
  margin: 0;
}

.user button {
  border: none;
  background: none;
  cursor: pointer;
  padding: .5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: .1s;
  border-radius: 9999px;
}

.user button:hover {
  background: var(--neutral-gray);
}

.user-actions{
  display: flex;
  gap: 1em;
}
</style>