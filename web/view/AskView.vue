<script setup>
import useSendQuestion from '@/composable/useSendQuestion';
import { fetchJSON } from '@/helper/fetch';
import { ref } from 'vue';

const { form, loading, submit, status } = useSendQuestion()
const props = defineProps({
    email: String,
})

const user = ref({})

async function getUser() {
    const { json } = await fetchJSON(`/api/v1/users/${props.email}`)
    if (json.data) {
        user.value = json.data.user
    }
}

getUser()

</script>

<template>
    <div class="flex justify-center mt-10 mx-5 text-center">
        <form class="flex flex-col overflow-hidden shadow-2xl font-bold rounded-[2rem] self-start">
            <div class="bg-gradient-to-br from-blue-600 to-blue-400 text-white p-8">Ask {{ user.email }} Anything!</div>
            <textarea
                v-model="form.question" name="" id="" cols="30" rows="5"
                placeholder="Type your question."
                class="border border-blue-300 px-3"
                ></textarea>
            <button @click="submit" class="block py-3 bg-blue-300" :disabled="loading == true || status == 'success'">
                <span v-if="status === 'success'">Terkirim!</span>
                <span v-else-if="status === 'fail'">Tidak Terkirim, kirim ulang!</span>
                <span v-else="loading">Loading...</span>
                <span v-else>Kirim</span>
            </button>
        </form>
    </div>
</template>