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
    <h1>Le monk</h1>
    <p>{{ monkeys }}</p>
    <!-- <v-row no-gutters>
        <v-col cols="6" v-for="monkey in monkeys" :key="monkey.id">
            <MonkeyCard :titleText="monkey.monkeyName" :imageName="monkey.monkeyImage"></MonkeyCard>
        </v-col>
    </v-row> -->
</template>