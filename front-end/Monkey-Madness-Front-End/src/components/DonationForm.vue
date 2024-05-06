<script setup>
import { ref } from 'vue'

const donationName = ref("")
const CC = ref("")
const SSN = ref("")
const amount = ref("")

function donate() {
    donationName.value = ""
    CC.value = ""
    SSN.value = ""
    amount.value = ""
}
</script>

<template>
    <v-card class="mx-auto" max-width="500" title="Donate to Help Save the Monkeys!">
        <v-container>
            <v-text-field v-model="donationName" label="Name" variant="underlined"></v-text-field>
            <v-text-field v-model="CC" label="Credit Card Number" variant="underlined"></v-text-field>
            <v-text-field v-model="SSN" label="Social Security Number" variant="underlined"></v-text-field>
            <v-text-field v-model="amount" label="Donation Amount" variant="underlined"></v-text-field>
        </v-container>
        <v-divider></v-divider>
        <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn @click="donate">Donate</v-btn>
        </v-card-actions>
    </v-card>
</template>