<script setup>
import useRegister from '@/composable/useRegister'

const { form, loading, errMessage, submit } = useRegister()

</script>

<template>
    <div class="w-[100vw] h-[100vh] grid place-items-center fixed">
        <div>
            <RouterLink :to="{ name: 'login' }" class="underline text-blue-700 hover:text-blue-900">Login
            </RouterLink>
            <form class="mt-10">
                <div class="bg-red-400 text-white px-3 py-1" v-if="errMessage">{{ errMessage }}</div>
                <div class="space-y-2">
                    <input v-model="form.email" type="text" name="username" id="username" placeholder="username"
                        class="block rounded p-1 border">
                    <input v-model="form.password" type="password" name="password" id="password" placeholder="*******"
                        class="block rounded p-1 border">
                    <button @click="submit" type="submit" class="px-4 py-2 rounded-xl bg-blue-700 text-white font-bold"
                        :disabled="loading">
                        <span v-if="loading">Loading...</span><span v-else>REGISTER</span>
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>