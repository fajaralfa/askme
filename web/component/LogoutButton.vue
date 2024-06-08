<script setup>
import useLogout from '@/composable/useLogout'

const { submit } = useLogout()
</script>

<template>
    <button @click="submit"
        class="px-4 py-2 bg-blue-300 text-white rounded mt-5 absolute bottom-5 left-5">Logout</button>
</template>