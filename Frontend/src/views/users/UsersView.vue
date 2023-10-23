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
          <PhPlus />
        </button>
      </div>
    </div>
    <div
      v-if="model.hasError"
      class="error user"
    >
      {{ model.message }}
    </div>
    <div
      v-for="user in model.data"
      v-else
      :key="user.Id"
      class="row user"
    >
      <p>
        {{ user.Name }}
      </p>
      <div class="user-actions">
        <button
          type="button"
          @click="openAddEditModal(user)"
        >
          <PhPencil />
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
        <PhPlus />
      </button>
    </div>
  </div>
  <UsersAddEditModal
    ref="addEdit"
    @close="load"
  />
</template>
  <script>
  import { PhPlus, PhPencil } from '@phosphor-icons/vue';
  import UsersAddEditModal from './UsersAddEditModal.vue';
  
  export default {
    components: {
      PhPlus,
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
  <style>
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
  
  .title-add-user button {
    background: var(--primary);
    border-radius: 8px;
    border: none;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: .2s;
    gap: .5rem;
    padding: .5rem 1rem;
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
    border: 1px solid;
    padding: 1rem;
    border-radius: 4px;
    display: flex;
    cursor:text;
    justify-content: space-between;
    align-items: center;
  }

  .user:hover{
  background-color: rgba(200, 255, 172, 0.041);
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
    background: var(--neutral-dark-gray);
  }
  </style>