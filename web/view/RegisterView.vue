<script setup>
import useRegister from '@/composable/useRegister'

const {form, loading, errMessage, submit} = useRegister()

</script>

<template>
    <div class="p-3">
        <form>
            <div class="bg-red-400 text-white px-3 py-1" v-if="errMessage">{{ errMessage }}</div>
            <div class="space-y-2">
                <input v-model="form.email" type="email" name="email" id="email" placeholder="example@email.com"
                    class="block rounded">
                <input v-model="form.password" type="password" name="password" id="password" placeholder="*******"
                    class="block rounded">
                <button @click="submit" type="submit" class="bg-white px-4 py-2 rounded" :disabled="loading">
                    <span v-if="loading">Loading...</span><span v-else>Register</span>
                </button>
            </div>
        </form>
    </div>
</template>