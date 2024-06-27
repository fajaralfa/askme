<script setup>
import useGetUser from '@/composable/useGetUser';
import useSendQuestion from '@/composable/useSendQuestion';

const props = defineProps({
    email: String,
})

const { form, loading, submit, status } = useSendQuestion()
const { user, getUser } = useGetUser(props.email)

await getUser()

</script>

<template>
    <div class="flex justify-center mt-10 mx-5 text-center">
        <form class="flex flex-col overflow-hidden shadow-2xl font-bold rounded-[2rem] self-start">
            <div class="bg-gradient-to-br from-blue-600 to-blue-400 text-white p-8">Ask {{ user.email }} Anything!</div>
            <textarea v-model="form.question" name="" id="" cols="30" rows="5" placeholder="Type your question."
                class="border border-blue-300 px-3"></textarea>
            <button @click="submit" class="block py-3 bg-blue-300" :disabled="loading == true || status == 'success'">
                <span v-if="status === 'success'">Terkirim!</span>
                <span v-else-if="status === 'fail'">Tidak Terkirim, kirim ulang!</span>
                <span v-else-if="loading">Loading...</span>
                <span v-else>Kirim</span>
            </button>
        </form>
    </div>
</template>