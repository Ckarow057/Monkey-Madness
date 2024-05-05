<script setup>
import { ref } from 'vue'
import MonkeyCard from './components/MonkeyCard.vue'
import MyFooter from './components/Footer.vue'

const id = 0
const monkeys = ref("")

fetch('http://localhost:8000/displayallmonkeys', {
    method: 'POST'
})
    .then(response => response.json())
    .then(data => {
        monkeys.value = data
    })
</script>

<template>
    <h1>Our Monkeys</h1>
    <v-row>
        <v-col cols="6" v-for="monkey in monkeys" :key="monkey.MonkeyID">
            <MonkeyCard :MonkeyName="monkey.MonkeyName" :MonkeyBreed="monkey.MonkeyBreed" :MonkeyImg="monkey.MonkeyImg"
                :MonkeyFact="monkey.MonkeyFact">
            </MonkeyCard>
        </v-col>
    </v-row>
    <MyFooter></MyFooter>
</template>