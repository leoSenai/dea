<template>
  <div class="quiz-content">
    <div class="quiz-title">
      <div
        v-if="isMobile"
        class="title-go-back"
      >
        <button
          type="button"
          @click="goBack()"
        >
          <PhCaretLeft />
        </button>
      </div>
      <div class="title">
        <h4>Pessoas próximas de jorginho</h4>
      </div>
      <div
        v-if="!isMobile"
        class="title-add-quiz"
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
      class="error quiz"
    >
      {{ model.message }}
    </div>
    <div
      v-for="quiz in model.data"
      v-else
      :key="quiz.IdQuiz"
      class="row quiz"
    >
      <p>
        {{ quiz.Name }}
      </p>
      <div class="quiz-actions">
        <button
          type="button"
          @click="openAddEditModal(quiz)"
        >
          <PhPencil />
        </button>
      </div>
    </div>
    <div
      v-if="isMobile"
      class="add-quiz"
    >
      <button
        type="button"
        @click="openAddEditModal()"
      >
        <PhPlus />
      </button>
    </div>
  </div>
  <ProximityAddEditModal
    ref="addEdit"
    @close="load"
  />
</template>
<script>
import { PhPlus, PhPencil, PhCaretLeft } from '@phosphor-icons/vue';
import ProximityAddEditModal from './ProximityAddEditModal.vue';

export default {
  components: {
    PhPlus,
    ProximityAddEditModal,
    PhPencil,
    PhCaretLeft
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
    },
    patientId() {
      return this.$router.currentRoute.value.params.id
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      const th = this;
      th.$api.ProximityController.getPersonsByIdPatient(th.patientId).then(({ data }) => {
        th.model = data
      })
    },
    openAddEditModal(current) {
      this.$refs.addEdit.openModal(current)
    },
    goBack() {
      this.$router.push('/paciente/id')
    }
  },
}
</script>
<style>
.quiz-content {
  padding: 3rem 1.5rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: .75rem;
}

.quiz-title {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.title-add-quiz button {
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

.title-add-quiz button:hover {
  filter: brightness(0.8);
}

.title-go-back button {
  background: transparent;
  border: none;
  height: 1rem;
  width: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.quiz-title:has(.title-go-back) {
  justify-content: flex-start;
  align-items: center;
  gap: 1rem;
}

.error {
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
  padding: 1rem;
}

.add-quiz {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.add-quiz button {
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

.add-quiz button:hover {
  filter: brightness(0.8);
  cursor: pointer;
}

.quiz {
  border: 1px solid;
  padding: 1rem;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.quiz p {
  margin: 0;
}

.quiz button {
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

.quiz button:hover {
  background: var(--neutral-dark-gray);
}
</style>