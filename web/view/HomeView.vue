<script setup>
import Nav from '@/component/Nav.vue';
import LogoutButton from '@/component/LogoutButton.vue'
import * as user from '@/store/userStore'

const link = `${window.location.origin}/${user.get().value.email}`

</script>

<template>
    <Nav />
    <LogoutButton />
    <div class="max-w-2xl m-auto">
        <div class="flex flex-col items-center gap-10 mt-28">
            <h1>Salin Link Dibawah</h1>
            <div class="text-base p-3 border rounded box-border active:outline-white active:outline active:outline-1"
            >{{ link }}</div>
        </div>
    </div>
</template>