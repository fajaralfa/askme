<script setup>
import useLogout from '@/composable/useLogout'

const { submit } = useLogout()
</script>

<template>
    <button @click="submit"
        class="px-4 py-2 bg-blue-200 hover:bg-blue-900 text-blue-900 hover:text-blue-200 transition rounded mt-5 absolute bottom-5 left-5">Logout</button>
</template>