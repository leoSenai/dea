<template>
  <div class="q-pa-md">
    <q-layout
      view="hHh Lpr lff"
      class="shadow-2 rounded-borders"
    >
      <q-header
        elevated
        class="bg-primary-700"
      >
        <q-toolbar>
          <q-btn
            flat
            round
            dense
            :icon="menu"
            @click="drawer = !drawer"
          >
            <PhList />
          </q-btn>
          <q-toolbar-title>
            Clínica Motivar
          </q-toolbar-title>
        </q-toolbar>
      </q-header>

      <q-drawer
        v-model="drawer"
        show-if-above

        :mini="miniState"
        mini-to-overlay
        :width="200"
        :breakpoint="500"

        bordered
        class="bg-success"
        @mouseover="miniState = false"
        @mouseout="miniState = true"
      >
        <q-scroll-area
          class="fit"
          :horizontal-thumb-style="{ opacity: 0 }"
        >
          <q-list padding>
            <q-item
              v-for="link in links"
              :key="link.path"
              v-ripple
              clickable
            >
              <q-item-section avatar>
                <component
                  :is="link.icon"
                  class="link-icon"
                />
              </q-item-section>
              <q-item-section>
                {{ link.name }}
              </q-item-section>
            </q-item>
          </q-list>
        </q-scroll-area>
      </q-drawer>

      <q-page-container>
        <q-page padding>
          <p
            v-for="n in 15"
            :key="n"
          >
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Fugit nihil praesentium molestias a adipisci, dolore vitae odit, quidem consequatur optio voluptates asperiores pariatur eos numquam rerum delectus commodi perferendis voluptate?
          </p>
        </q-page>
      </q-page-container>
    </q-layout>
  </div>
</template>

<script>
import { PhList } from '@phosphor-icons/vue';

export default {
  components: {
    PhList,
  },
  props: {
    links: {
      type: Array,
      required: true,
    },
  },
  data () {
    return {
      logoSrc: 'src/assets/imgs/Logo.png',
      drawer: false,
      miniState: true
    }
  }
}
</script>
