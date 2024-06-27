<script setup>
import useLogin from '@/composable/useLogin';

const { form, loading, errMessage, submit } = useLogin()

</script>

<template>
    <div class="w-[100vw] h-[100vh] grid place-items-center fixed bg-gradient-to-br from-blue-300 to-blue-900">
        <div>
            <form>
                <div class="bg-red-400 text-white px-3 py-1" v-if="errMessage">{{ errMessage }}</div>
                <div class="space-y-3">
                    <input v-model="form.email" type="text" name="username" id="username" placeholder="username"
                        class="block rounded p-1 border-2 bg-blue-100">
                    <input v-model="form.password" type="password" name="password" id="password" placeholder="*******"
                        class="block rounded p-1 border-2 bg-blue-100">
                    <div>
                        <RouterLink :to="{ name: 'register' }"
                        class="text-white hover:text-blue-200 underline">
                            Register</RouterLink>
                    </div>
                    <button @click="submit" type="submit"
                        class="px-4 py-1 rounded-xl bg-blue-700 hover:bg-blue-200 text-blue-200 hover:text-blue-700 transition font-bold" :disabled="loading">
                        <span v-if="loading">Loading...</span><span v-else>LOGIN</span>
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>