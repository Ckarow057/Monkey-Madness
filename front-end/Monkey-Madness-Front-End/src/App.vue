<script setup>
import { ref, computed } from 'vue'
import Home from './Home.vue'
import Encyclopedia from './Encyclopedia.vue'
import Donate from './Donate.vue'

const routes = {
  '/': Home,
  '/encyclopedia': Encyclopedia,
  '/donate': Donate,
}

const currentPath = ref(window.location.hash)
const drawer = ref(false)

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash
})

const currentView = computed(() => {
  return routes[currentPath.value.slice(1) || '/'] || NotFound
})
</script>

<template>
  <v-app>
    <v-navigation-drawer color="#5C4033" v-model="drawer">
      <v-list-item prependIcon="mdi-home" href="#/" title="Home" @click="drawer = !drawer"></v-list-item>
      <v-list-item prependIcon="mdi-book-open-page-variant-outline" href="#/encyclopedia" title="Monkey Encyclopedia"
        @click="drawer = !drawer"></v-list-item>
      <v-list-item prependIcon="mdi-cash-multiple" href="#/donate" title="Donate"
        @click="drawer = !drawer"></v-list-item>
    </v-navigation-drawer>
    <v-app-bar color="#098205">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title>Monkey Madness</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <component :is="currentView"></component>
    </v-main>
  </v-app>
</template>