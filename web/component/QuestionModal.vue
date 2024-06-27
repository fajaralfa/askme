<script setup>
const props = defineProps({
    title: String,
    question: String,
    show: Boolean,
})

</script>

<template>
    <div v-if="show" id="modal"
        class="fixed top-0 w-full h-full !bg-[rgba(0,0,0,0.4)] flex items-center justify-center">
        <button class="absolute right-10 top-20 lg:right-1/3 text-white font-bold"
            @click="$emit('show', false)">Close</button>
        <div class="mx-6 w-full text-2xl font-bold text-center rounded-[2rem] overflow-hidden lg:!w-[24rem]">
            <div class="bg-gradient-to-br from-blue-600 to-blue-400 text-white p-8">{{ title }}</div>
            <div class="bg-white p-8">{{ question }}</div>
        </div>
    </div>
</template>