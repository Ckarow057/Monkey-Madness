<script setup>

defineProps({
    titleText: String,
    imageName: String,
})

</script>

<template>
    <v-card>
        <v-card-title>{{ titleText }}</v-card-title>
        <img :src="`../../public/images/${imageName}`">
        <br>
        <v-card-text>
            <slot></slot>
        </v-card-text>
    </v-card>
</template>