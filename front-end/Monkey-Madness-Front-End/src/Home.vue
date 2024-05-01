<script setup>
import { ref } from 'vue'
import RandomMonkey from './components/RandomMonkey.vue';
import MonkeyCard from './components/MonkeyCard.vue';

const id = 0
const monkeys = ref("")

fetch('http://localhost:8000/displaymonkey', {
    method: 'POST'
})
    .then(response => response.json())
    .then(data => {
        monkeys.value = data
        console.log(monkey)
    })
</script>

<template>
    <h1>This is Home Page</h1>
    <!-- make welcome message -->
    <div class="bg-green-lighten-5 text-brown-lighten-1 d-flex justify-center text-h2 pa-8">Welcome to Monkey Madness
    </div>
    <!-- add random monkey -->
    <v-row no-gutters>
        <v-col cols="4" v-for="monkey in monkeys" :key="monkey.MonkeyID">
            <RandomMonkey :MonkeyName="monkey.MonkeyName" :MonkeyBreed="monkey.MonkeyBreed"
                :MonkeyImg="monkey.MonkeyImg" :MonkeyFact="monkey.MonkeyFact">
            </RandomMonkey>
        </v-col>
    </v-row>
    <!-- footer with random things-->
</template>