<script setup>

defineProps({
    MonkeyName: String,
    MonkeyBreed: String,
    MonkeyImg: String,
    MonkeyFact: String,
})

</script>

<template>
    <div>
        <v-hover v-slot="{ isHovering, props }">
            <v-card class="mx-auto" max-width="344" v-bind="props">
                <img :src="`../images/${MonkeyImg}`">

                <v-card-text>
                    <h2 class="text-h6 text-primary">
                        {{ MonkeyName }}
                    </h2>
                </v-card-text>

                <v-overlay :model-value="isHovering" class="align-center justify-center" scrim="#036358" contained>
                    <v-card>
                        <ul>
                            <li>Species Name: {{ MonkeyBreed }}</li>
                            <li>Fun Fact! {{ MonkeyFact }}</li>
                        </ul>
                    </v-card>
                </v-overlay>
            </v-card>
        </v-hover>
    </div>
</template>