<script setup>
import useLogin from '@/composable/useLogin';

const {form, loading, errMessage, submit} = useLogin()

</script>

<template>
    <div class="w-[100vw] h-[100vh] grid place-items-center fixed">
        <form>
            <div class="bg-red-400 text-white px-3 py-1" v-if="errMessage">{{ errMessage }}</div>
            <div class="space-y-2">
                <input v-model="form.email" type="text" name="username" id="username" placeholder="username"
                    class="block rounded p-1 border">
                <input v-model="form.password" type="password" name="password" id="password" placeholder="*******"
                    class="block rounded p-1 border">
                <button @click="submit" type="submit" class="px-4 py-2 rounded-xl bg-blue-700 text-white font-bold" :disabled="loading">
                    <span v-if="loading">Loading...</span><span v-else>LOGIN</span>
                </button>
            </div>
        </form>
    </div>
</template>