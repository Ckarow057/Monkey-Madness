<script setup>
import { ref } from 'vue'
import MonkeyCard from './components/MonkeyCard.vue'

const id = 0
const monkeys = ref("")

fetch('http://localhost:8000/displayallmonkeys', {
    method: 'POST'
})
    .then(response => response.json())
    .then(data => {
        monkeys.value = data
        console.log(monkeys.value)
    })
</script>

<template>
    <h1>Our Monkeys</h1>
    <v-row no-gutters>
        <v-col cols="4" v-for="monkey in monkeys" :key="monkey.MokneyID">
            <MonkeyCard :MonkeyName="monkey.MonkeyName" :MonkeyBreed="monkey.MonkeyBreed" :MonkeyImg="monkey.MonkeyImg"
                :MonkeyFact="monkey.MonkeyFact">
            </MonkeyCard>
        </v-col>
    </v-row>
</template>