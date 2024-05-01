<script setup>
import { ref } from 'vue'
import RandomMonkey from './components/RandomMonkey.vue';
import MonkeyCard from './components/MonkeyCard.vue';

const monkey = ref("")

fetch('http://localhost:8000/displaymonkey', {
    method: 'POST'
})
    .then(response => response.json())
    .then(data => {
        monkey.value = data
        console.log(monkey)
    })
</script>

<template>
    <h1>This is Home Page</h1>
    <!-- make welcome message -->
    <div class="bg-green-lighten-5 text-brown-lighten-1 d-flex justify-center text-h2 pa-8">Welcome to Monkey Madness
    </div>
    <!-- add random monkey -->
    <MonkeyCard :MonkeyName="monkey.MonkeyName" :MonkeyBreed="monkey.MonkeyImg" :MonkeyImg="monkey.MonkeyImg"
        :MonkeyFact="monkey.MonkeyFact">
    </MonkeyCard>
    <!-- footer with random things-->
</template>