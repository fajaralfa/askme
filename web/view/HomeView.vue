<script setup>
import Nav from '@/component/Nav.vue';
import LogoutButton from '@/component/LogoutButton.vue'
import * as user from '@/store/userStore'

const link = `${window.location.origin}/${user.get().value.email}`

</script>

<template>
    <div class="w-[100vw] min-h-[100vh] place-items-center fixed bg-gradient-to-br from-blue-400 to-blue-600">
        <Nav />
        <LogoutButton />
        <div class="max-w-2xl m-auto">
            <div class="flex flex-col items-center gap-10 mt-28">
                <h1 class="text-white animate-bounce">Salin Link Dibawah</h1>
                <div
                    class="text-base p-3 border rounded box-border bg-blue-200 active:outline-white active:outline active:outline-1">
                    {{ link }}</div>
            </div>
        </div>
    </div>
</template>