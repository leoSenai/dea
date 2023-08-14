<template>
  <div
    v-if="modelValue"
    class="modal"
  >
    <q-card class="card">
      <div
        class="close-button"
        @click="$emit('update:modelValue', false)"
      >
        <PhX
          size="1rem"
          color="black"
        />
      </div>
      <h5 class="modal-title">
        <slot name="modal-title" />
      </h5>
      <q-card-section class="modal-content">
        <slot name="modal-content" />
      </q-card-section>
      <q-card-actions class="modal-actions">
        <slot name="modal-actions" />
      </q-card-actions>
    </q-card>
  </div>
</template>
<script>
import { PhX } from '@phosphor-icons/vue';

export default {
  components: {
    PhX,
  },
  props: {
    modelValue: {
      type: Boolean,
      default: false
    }
  },
  emits: ['update:modelValue'],
}
</script>
<style scoped>
.modal {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;

  color: black;
  background: rgba(0, 0, 0, .6);
  display: flex;
  padding: 2rem 1rem;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2rem;

  z-index: 1;

}

.card {
  position: relative;
  display: flex;
  flex-direction: column;
  background: #fff;
  height: calc(100vh - 13rem);
  width: 100%;
  border-radius: 4px;
  padding: 1rem;
}

.close-button {
  position: absolute;
  right: 0.5rem;
  top: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-radius: 99999px;
  padding: .25rem;
}

.modal-content {
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  width: 100%;
  flex: 1;
}

.close-button:hover {
  background: rgba(0, 0, 0, .15);
}

.modal-actions {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: .5rem;
}

@media (min-width: 768px) {
  .modal {
    padding: 4rem 1rem;
  }

  .card {
    width: 70%;
  }
}
</style>