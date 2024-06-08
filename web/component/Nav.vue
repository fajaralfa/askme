<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
    <nav class="flex justify-center gap-10 p-6">
        <RouterLink :to="{name: 'home'}" class="text-[rgba(0,0,0,0.3)]"
            exactActiveClass="font-bold !text-black">PLAY</RouterLink>
        <RouterLink :to="{name: 'inbox'}" class="text-[rgba(0,0,0,0.3)]"
            exactActiveClass="font-bold !text-black">INBOX</RouterLink>
    </nav>
</template>