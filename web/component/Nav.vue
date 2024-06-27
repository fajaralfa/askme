<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
    <nav class="flex justify-center gap-10 p-6">
        <RouterLink :to="{ name: 'home' }" class="text-white" exactActiveClass="font-bold">PLAY</RouterLink>
        <RouterLink :to="{ name: 'inbox' }" class="text-white" exactActiveClass="font-bold">INBOX</RouterLink>
    </nav>
</template>