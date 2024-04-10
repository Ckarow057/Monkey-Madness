<script setup>
import { ref } from 'vue'
import MonkeyCard from './components/MonkeyCard.vue'

let id = 0;
const monkeys = ref([
    { id: id++, monkeyName: 'Monkey1', monkeyImage: 'placeholder.png' }
])
// do fetch from monkey api here

</script>

<template>
    <h1>Le monk</h1>
    <v-row no-gutters>
        <v-col cols="6" v-for="monkey in monkeys" :key="monkey.id">
            <MonkeyCard :titleText="monkey.monkeyName" :iamgeName="monkey.monkeyImage"></MonkeyCard>
        </v-col>
    </v-row>
</template>