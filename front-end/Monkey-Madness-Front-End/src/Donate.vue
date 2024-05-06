<script setup>
import { ref } from 'vue';
import Footer from './components/Footer.vue';
import DonationForm from './components/DonationForm.vue'

const fools = ref("")

fetch('http://localhost:8000/displaypersonalinfo', {
    method: 'POST'
})
    .then(response => response.json())
    .then(data => {
        fools.value = data
    })

</script>

<template>
    <div class="bg-green-lighten-5 text-brown-lighten-1 d-flex justify-center text-h2 pa-8">Donate to Help us Support
        and Save the Monkeys
    </div>
    <DonationForm></DonationForm>
    <br>
    <v-hover v-slot="{ isHovering, props }">
        <v-card class="mx-auto" max-width="100%" v-bind="props">
            <v-img src="../images/banner.jpg" Width="100%"></v-img>

            <v-overlay :model-value="isHovering" class="align-center justify-center" scrim="#036358" contained>
                <v-card>
                    <v-table theme="dark">
                        <thead>
                            <tr>
                                <!-- <th class="text-left">
                                    ID
                                </th> -->
                                <th class="text-left">
                                    Credit Card Number
                                </th>
                                <th class="text-left">
                                    Social Security Number
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="fool in fools" :key="fool.UserID">
                                <!-- <td>{{ fool.UserID }}</td> -->
                                <td>{{ fool.UserSSN }}</td>
                                <td>{{ fool.UserCardNum }}</td>
                            </tr>
                        </tbody>
                    </v-table>
                </v-card>
            </v-overlay>
        </v-card>
    </v-hover>
    <Footer></Footer>
</template>